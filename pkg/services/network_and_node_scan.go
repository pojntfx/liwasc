package services

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"

import (
	"context"
	"log"
	"time"

	cmap "github.com/orcaman/concurrent-map"
	"github.com/pojntfx/liwasc/pkg/databases"
	proto "github.com/pojntfx/liwasc/pkg/proto/generated"
	"github.com/pojntfx/liwasc/pkg/scanners"
	liwascModels "github.com/pojntfx/liwasc/pkg/sql/generated/liwasc"
	mac2vendorModels "github.com/pojntfx/liwasc/pkg/sql/generated/mac2vendor"
	"github.com/ugjka/messenger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NetworkAndNodeScanService struct {
	proto.UnimplementedNetworkAndNodeScanServiceServer

	device                          string
	mac2VendorDatabase              *databases.MAC2VendorDatabase
	serviceNamesPortNumbersDatabase *databases.ServiceNamesPortNumbersDatabase
	ports2PacketsDatabase           *databases.Ports2PacketDatabase
	liwascDatabase                  *databases.LiwascDatabase
	messengers                      cmap.ConcurrentMap
}

func NewNetworkAndNodeScanService(
	device string,
	mac2VendorDatabase *databases.MAC2VendorDatabase,
	serviceNamesPortNumbersDatabase *databases.ServiceNamesPortNumbersDatabase,
	ports2PacketsDatabase *databases.Ports2PacketDatabase,
	liwascDatabase *databases.LiwascDatabase,
) *NetworkAndNodeScanService {
	return &NetworkAndNodeScanService{
		device:                          device,
		mac2VendorDatabase:              mac2VendorDatabase,
		serviceNamesPortNumbersDatabase: serviceNamesPortNumbersDatabase,
		ports2PacketsDatabase:           ports2PacketsDatabase,
		liwascDatabase:                  liwascDatabase,
		messengers:                      cmap.New(),
	}
}

func (s *NetworkAndNodeScanService) TriggerNetworkScan(ctx context.Context, scanTriggerMessage *proto.NetworkScanTriggerMessage) (*proto.NetworkScanReferenceMessage, error) {
	// Create a scan
	scan := &liwascModels.NetworkScan{
		Done: 0,
	}

	scanID, err := s.liwascDatabase.CreateNetworkScan(scan)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "could not create scan in DB: %v", err.Error())
	}

	networkScanner := scanners.NewNetworkScanner(s.device)
	err, _ = networkScanner.Open()
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "could not open network scanner: %v", err.Error())
	}

	msgr := messenger.New(0, true)
	s.messengers.Set(string(scanID), msgr)

	// Receive packets
	go func() {
		receiveCtx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(scanTriggerMessage.GetTimeout()))
		defer cancel()

		if err := networkScanner.Receive(receiveCtx); err != nil {
			log.Println("could not receive from network scanner", err)

			return
		}
	}()

	// Transmit packets ("start a scan")
	go func() {
		if err := networkScanner.Transmit(); err != nil {
			log.Println("could not transmit from network scanner", err)

			return
		}
	}()

	// Read packets
	go func() {
		for {
			node := networkScanner.Read()

			// Network scan is done
			if node == nil {
				break
			}

			// Lookup vendor information for node
			vendor, err := s.mac2VendorDatabase.GetVendor(node.MACAddress.String())
			if err != nil {
				vendor = &mac2vendorModels.Vendordb{}
			}

			dbNode := &liwascModels.Node{
				MacAddress:   node.MACAddress.String(),
				IPAddress:    node.IPAddress.String(),
				Vendor:       vendor.Vendor.String,
				Registry:     vendor.Registry,
				Organization: vendor.Organization.String,
				Address:      vendor.Address.String,
				Visible:      vendor.Visibility,
			}

			if _, err := s.liwascDatabase.UpsertNode(dbNode, scanID); err != nil {
				log.Println("could not create node in DB", err)

				break
			}

			msgr.Broadcast(dbNode)
		}

		msgr.Reset()

		scan.Done = 1
		if _, err := s.liwascDatabase.UpdateNetworkScan(scan); err != nil {
			log.Println("could not update scan in DB", err)

			return
		}

		s.messengers.Remove(string(scanID))
	}()

	return &proto.NetworkScanReferenceMessage{NetworkScanID: scanID}, nil
}

func (s *NetworkAndNodeScanService) SubscribeToNewNodes(scanReferenceMessage *proto.NetworkScanReferenceMessage, stream proto.NetworkAndNodeScanService_SubscribeToNewNodesServer) error {
	allNodes, err := s.liwascDatabase.GetAllNodes()
	if err != nil {
		return status.Errorf(codes.Unknown, "could not get nodes from DB: %v", err.Error())
	}

	matchingNewestScans, err := s.liwascDatabase.GetNewestNetworkScansForNodes(allNodes)
	if err != nil {
		return status.Errorf(codes.Unknown, "could not get scans from DB: %v", err.Error())
	}

	scan, err := s.liwascDatabase.GetNewestNetworkScan()
	if err != nil {
		return status.Errorf(codes.Unknown, "could not get latest scan from DB: %v", err.Error())
	}

	for _, dbNode := range allNodes {
		protoNode := &proto.DiscoveredNodeMessage{
			NodeScanID: -1, // TODO: Add node scan once service/port scanning is implemented
			LucidNode: &proto.LucidNodeMessage{
				PoweredOn: func() bool {
					for nodeID := range matchingNewestScans {
						if nodeID == dbNode.MacAddress {
							if scanReferenceMessage.GetNetworkScanID() == -1 {
								if scan.ID == matchingNewestScans[dbNode.MacAddress][0] { // If the node is in the newest scan, it is powered on
									return true
								}

								return false
							}

							for _, scanID := range matchingNewestScans[dbNode.MacAddress] {
								if scanID == scanReferenceMessage.GetNetworkScanID() { // The node was scanned in this scan; therefore the node is powered on (otherwise it would not have been found)
									return true
								}
							}
						}
					}

					return false
				}(),
				MACAddress:   dbNode.MacAddress,
				IPAddress:    dbNode.IPAddress,
				Vendor:       dbNode.Vendor,
				Registry:     dbNode.Registry,
				Organization: dbNode.Organization,
				Address:      dbNode.Address,
				Visible: func() bool {
					if dbNode.Visible == 1 {
						return true
					}

					return false
				}(),
			},
		}

		if err := stream.Send(protoNode); err != nil {
			return status.Errorf(codes.Unknown, "could not send node to frontend: %v", err.Error())
		}
	}

	if scanReferenceMessage.GetNetworkScanID() != -1 {
		scan, err = s.liwascDatabase.GetNetworkScan(scanReferenceMessage.GetNetworkScanID())
		if err != nil {
			return status.Errorf(codes.Unknown, "could not get scan from DB: %v", err.Error())
		}
	}

	if scan.Done == 1 {
		return nil
	}

	msgr, exists := s.messengers.Get(string(scan.ID))
	if !exists || msgr == nil {
		return nil
	}

	client, err := msgr.(*messenger.Messenger).Sub()
	if err != nil {
		return status.Errorf(codes.Unknown, "could not subscribe to nodes")
	}

	for receivedNode := range client {
		dbNode := receivedNode.(*liwascModels.Node)

		protoNode := &proto.DiscoveredNodeMessage{
			NodeScanID: -1, // TODO: Add node scan once service/port scanning is implemented
			LucidNode: &proto.LucidNodeMessage{
				PoweredOn:    true, // Must be true; otherwise it would not have been found
				MACAddress:   dbNode.MacAddress,
				IPAddress:    dbNode.IPAddress,
				Vendor:       dbNode.Vendor,
				Registry:     dbNode.Registry,
				Organization: dbNode.Organization,
				Address:      dbNode.Address,
				Visible: func() bool {
					if dbNode.Visible == 1 {
						return true
					}

					return false
				}(),
			},
		}

		if err := stream.Send(protoNode); err != nil {
			return status.Errorf(codes.Unknown, "could not send node to frontend: %v", err.Error())
		}
	}

	return nil
}
