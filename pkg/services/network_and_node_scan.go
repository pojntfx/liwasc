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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NetworkAndNodeScanService struct {
	proto.UnimplementedNetworkAndNodeScanServiceServer

	device             string
	networkScanners    cmap.ConcurrentMap
	mac2VendorDatabase *databases.MAC2VendorDatabase
	liwascDatabase     *databases.LiwascDatabase
}

func NewNetworkAndNodeScanService(device string, mac2VendorDatabase *databases.MAC2VendorDatabase, liwascDatabase *databases.LiwascDatabase) *NetworkAndNodeScanService {
	return &NetworkAndNodeScanService{device: device, networkScanners: cmap.New(), mac2VendorDatabase: mac2VendorDatabase, liwascDatabase: liwascDatabase}
}

func (s *NetworkAndNodeScanService) TriggerNetworkScan(ctx context.Context, scanTriggerMessage *proto.NetworkScanTriggerMessage) (*proto.NetworkScanReferenceMessage, error) {
	// Create a scan
	scan := &liwascModels.Scan{
		Done: 0,
	}

	scanID, err := s.liwascDatabase.CreateScan(scan)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "could not create scan in DB: %v", err.Error())
	}

	networkScanner := scanners.NewNetworkScanner(s.device)
	err, _ = networkScanner.Open()
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "could not open network scanner: %v", err.Error())
	}
	s.networkScanners.Set(string(scanID), networkScanner)

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
		}

		scan.Done = 1
		if _, err := s.liwascDatabase.UpdateScan(scan); err != nil {
			log.Println("could not update scan in DB", err)

			return
		}
	}()

	return &proto.NetworkScanReferenceMessage{NetworkScanID: scanID}, nil
}

func (s *NetworkAndNodeScanService) SubscribeToNewNodes(scanReferenceMessage *proto.NetworkScanReferenceMessage, stream proto.NetworkAndNodeScanService_SubscribeToNewNodesServer) error {
	allNodes, err := s.liwascDatabase.GetAllNodes()
	if err != nil {
		return status.Errorf(codes.Unknown, "could not get nodes from DB: %v", err.Error())
	}

	// TODO: Subscribe to messenger for discovered nodes if messenger is set in cmap until recv node is nil (scan finished)

	for _, dbNode := range allNodes {
		protoNode := &proto.DiscoveredNodeMessage{
			NodeScanID: -1, // TODO: Get from join table; select newest in join table for this node
			LucidNode: &proto.LucidNodeMessage{
				PoweredOn:    false, // TODO: Get from join table; if scanID is in the array, it is powered on
				MACAddress:   dbNode.MacAddress,
				IPAddress:    dbNode.IPAddress,
				Vendor:       dbNode.Vendor,
				Registry:     dbNode.Registry,
				Organization: dbNode.Organization,
				Address:      dbNode.Address,
				Visible: func() bool {
					visible := false
					if dbNode.Visible == 1 {
						visible = true
					}

					return visible
				}(),
			},
		}

		if err := stream.Send(protoNode); err != nil {
			return status.Errorf(codes.Unknown, "could not send node to frontend: %v", err.Error())
		}
	}

	return nil
}
