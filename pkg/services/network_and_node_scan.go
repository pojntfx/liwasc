package services

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"

import (
	"context"
	"log"
	"math"
	"strings"
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
	networkScanMessengers           cmap.ConcurrentMap
	nodeScanMessengers              cmap.ConcurrentMap
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
		networkScanMessengers:           cmap.New(),
		nodeScanMessengers:              cmap.New(),
	}
}

func (s *NetworkAndNodeScanService) TriggerNetworkScan(ctx context.Context, scanTriggerMessage *proto.NetworkScanTriggerMessage) (*proto.NetworkScanReferenceMessage, error) {
	// Create a scan
	networkScan := &liwascModels.NetworkScan{
		Done: 0,
	}
	networkScanID, err := s.liwascDatabase.CreateNetworkScan(networkScan)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "could not create network scan in DB: %v", err.Error())
	}

	networkScanner := scanners.NewNetworkScanner(s.device)
	err, _ = networkScanner.Open()
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "could not open network scanner: %v", err.Error())
	}

	networkScanMessenger := messenger.New(0, true)
	s.networkScanMessengers.Set(string(networkScanID), networkScanMessenger)

	// Receive packets
	go func() {
		receiveCtx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(scanTriggerMessage.GetNetworkScanTimeout()))
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

			if _, err := s.liwascDatabase.UpsertNode(dbNode, networkScanID); err != nil {
				log.Println("could not create node in DB", err)

				break
			}

			// Scan for open ports for node
			// TODO: This is very expensive. The port scanners should be coordinated to run sequentially so that CPU usage isn't that high.
			portScanner := scanners.NewPortScanner(node.IPAddress.String(), 0, math.MaxUint16, time.Millisecond*time.Duration(scanTriggerMessage.GetNodeScanTimeout()), []string{"tcp", "udp"}, func(port int) ([]byte, error) {
				packet, err := s.ports2PacketsDatabase.GetPacket(port)

				if err != nil {
					return nil, err
				}

				return packet.Packet, err
			})

			// Dial and/or transmit packets ("start a scan")
			go func() {
				if err := portScanner.Transmit(); err != nil {
					log.Println("could not transmit from node scanner", err)
				}
			}()

			go func() {
				// Create a scan
				nodeScan := &liwascModels.NodeScan{
					Done: 0,
				}
				nodeScanID, err := s.liwascDatabase.CreateNodeScan(nodeScan, dbNode.MacAddress, networkScanID)
				if err != nil {
					log.Println("could not create node scan in DB", err)

					return
				}

				nodeScanMessenger := messenger.New(0, true)
				s.nodeScanMessengers.Set(string(nodeScanID), nodeScanMessenger)

				for {
					port := portScanner.Read()

					// Port scan is done
					if port == nil {
						break
					}

					// Note above does not apply here, there is no point in transmitting this info if the ports are closed
					if port.Open {
						services, err := s.serviceNamesPortNumbersDatabase.GetService(port.Port, port.Protocol)
						if err != nil {
							if !strings.Contains(err.Error(), "could find service") {
								log.Println("could not get services for port", err)
							}
						}

						var dbService *liwascModels.Service
						if len(services) > 0 {
							dbService = &liwascModels.Service{
								ServiceName:             services[0].ServiceName,
								PortNumber:              int64(port.Port),
								TransportProtocol:       services[0].TransportProtocol,
								Description:             services[0].Description,
								Assignee:                services[0].Assignee,
								Contact:                 services[0].Contact,
								RegistrationDate:        services[0].RegistrationDate,
								ModificationDate:        services[0].ModificationDate,
								Reference:               services[0].Reference,
								ServiceCode:             services[0].ServiceCode,
								UnauthorizedUseReported: services[0].UnauthorizedUseReported,
								AssignmentNotes:         services[0].AssignmentNotes,
							}
						} else {
							dbService = &liwascModels.Service{
								ServiceName:             "Non-Registered Service",
								PortNumber:              int64(port.Port),
								TransportProtocol:       port.Protocol,
								Description:             "",
								Assignee:                "",
								Contact:                 "",
								RegistrationDate:        "",
								ModificationDate:        "",
								Reference:               "",
								ServiceCode:             "",
								UnauthorizedUseReported: "",
								AssignmentNotes:         "",
							}
						}

						if _, err := s.liwascDatabase.UpsertService(dbService, dbNode.MacAddress, nodeScanID, networkScanID); err != nil {
							log.Println("could not create node in DB", err)

							break
						}

						nodeScanMessenger.Broadcast(dbService)
					}
				}

				nodeScanMessenger.Reset()

				nodeScan.Done = 1
				if _, err := s.liwascDatabase.UpdateNodeScan(nodeScan); err != nil {
					log.Println("could not update node scan in DB", err)

					return
				}

				s.nodeScanMessengers.Remove(string(nodeScanID))
			}()

			networkScanMessenger.Broadcast(dbNode)
		}

		networkScanMessenger.Reset()

		networkScan.Done = 1
		if _, err := s.liwascDatabase.UpdateNetworkScan(networkScan); err != nil {
			log.Println("could not update network scan in DB", err)

			return
		}

		s.networkScanMessengers.Remove(string(networkScanID))
	}()

	return &proto.NetworkScanReferenceMessage{NetworkScanID: networkScanID}, nil
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

	var networkScan *liwascModels.NetworkScan
	if scanReferenceMessage.GetNetworkScanID() == -1 {
		networkScan, err = s.liwascDatabase.GetNewestNetworkScan()
		if err != nil {
			return status.Errorf(codes.Unknown, "could not get latest scan from DB: %v", err.Error())
		}
	} else {
		networkScan, err = s.liwascDatabase.GetNetworkScan(scanReferenceMessage.GetNetworkScanID())
		if err != nil {
			return status.Errorf(codes.Unknown, "could not get scan from DB: %v", err.Error())
		}
	}

	nodeScansForNetworkScanAndNode := make(map[string]int64)
	for _, dbNode := range allNodes {
		nodeScanID, err := s.liwascDatabase.GetNodeScanIDByNetworkScanIDAndNodeID(dbNode.MacAddress, networkScan.ID)
		if err != nil {
			if strings.Contains(err.Error(), "sql: no rows in result set") {
				nodeScansForNetworkScanAndNode[dbNode.MacAddress] = -1
			} else {
				return status.Errorf(codes.Unknown, "could not get node scan from DB: %v", err.Error())
			}
		} else {
			nodeScansForNetworkScanAndNode[dbNode.MacAddress] = nodeScanID
		}

		protoNode := &proto.DiscoveredNodeMessage{
			NodeScanID: nodeScansForNetworkScanAndNode[dbNode.MacAddress],
			LucidNode: &proto.LucidNodeMessage{
				PoweredOn: func() bool {
					for nodeID := range matchingNewestScans {
						if nodeID == dbNode.MacAddress {
							if scanReferenceMessage.GetNetworkScanID() == -1 {
								if networkScan.ID == matchingNewestScans[dbNode.MacAddress][0] { // If the node is in the newest scan, it is powered on
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

	if networkScan.Done == 1 {
		return nil
	}

	msgr, exists := s.networkScanMessengers.Get(string(networkScan.ID))
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
			NodeScanID: nodeScansForNetworkScanAndNode[dbNode.MacAddress],
			LucidNode: &proto.LucidNodeMessage{
				PoweredOn:    true,
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

func (s *NetworkAndNodeScanService) SubscribeToNewOpenServices(nodeScanReferenceMessage *proto.NodeScanReferenceMessage, stream proto.NetworkAndNodeScanService_SubscribeToNewOpenServicesServer) error {
	return nil
}
