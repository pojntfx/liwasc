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
	mac2vendorModels "github.com/pojntfx/liwasc/pkg/sql/generated/mac2vendor"
	networkAndNodeScanModels "github.com/pojntfx/liwasc/pkg/sql/generated/network_and_node_scan"
	"github.com/robfig/cron"
	"github.com/ugjka/messenger"
	"golang.org/x/sync/semaphore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NetworkAndNodeScanService struct {
	proto.UnimplementedNetworkAndNodeScanServiceServer

	device                          string
	mac2VendorDatabase              *databases.MAC2VendorDatabase
	serviceNamesPortNumbersDatabase *databases.ServiceNamesPortNumbersDatabase
	ports2PacketsDatabase           *databases.Ports2PacketDatabase
	networkAndNodeScanDatabase      *databases.NetworkAndNodeScanDatabase
	networkScanMessengers           cmap.ConcurrentMap
	nodeScanMessengers              cmap.ConcurrentMap
	portScannerLock                 *semaphore.Weighted
	periodicScanCronExpression      string
	periodicNetworkScanTimeout      int
	periodicNodeScanTimeout         int
	cron                            *cron.Cron
	periodicScanMessenger           *messenger.Messenger
}

func NewNetworkAndNodeScanService(
	device string,
	mac2VendorDatabase *databases.MAC2VendorDatabase,
	serviceNamesPortNumbersDatabase *databases.ServiceNamesPortNumbersDatabase,
	ports2PacketsDatabase *databases.Ports2PacketDatabase,
	networkAndNodeScanDatabase *databases.NetworkAndNodeScanDatabase,
	portScannerLock *semaphore.Weighted,
	periodicScanCronExpression string,
	periodicNetworkScanTimeout int,
	periodicNodeScanTimeout int,
) *NetworkAndNodeScanService {
	return &NetworkAndNodeScanService{
		device:                          device,
		mac2VendorDatabase:              mac2VendorDatabase,
		serviceNamesPortNumbersDatabase: serviceNamesPortNumbersDatabase,
		ports2PacketsDatabase:           ports2PacketsDatabase,
		networkAndNodeScanDatabase:      networkAndNodeScanDatabase,
		networkScanMessengers:           cmap.New(),
		nodeScanMessengers:              cmap.New(),
		portScannerLock:                 portScannerLock,
		periodicScanCronExpression:      periodicScanCronExpression,
		periodicNetworkScanTimeout:      periodicNetworkScanTimeout,
		periodicNodeScanTimeout:         periodicNodeScanTimeout,
		cron:                            cron.New(),
		periodicScanMessenger:           messenger.New(0, true),
	}
}

func (s *NetworkAndNodeScanService) Open() error {
	if err := s.cron.AddFunc(s.periodicScanCronExpression, func() {
		protoNetworkScanTriggerMessage := &proto.NetworkScanTriggerMessage{
			NetworkScanTimeout: int64(s.periodicNetworkScanTimeout),
			NodeScanTimeout:    int64(s.periodicNodeScanTimeout),
		}

		protoNetworkScanReferenceMessage, err := s.TriggerNetworkScan(context.Background(), protoNetworkScanTriggerMessage)
		if err != nil {
			log.Println("could not trigger network scan", err)

			return
		}

		dbPeriodicNetworkScan := &networkAndNodeScanModels.PeriodicNetworkScansNetworkScan{
			NodeScanID: protoNetworkScanReferenceMessage.GetNetworkScanID(),
		}

		if _, err := s.networkAndNodeScanDatabase.CreatePeriodicNetworkScan(dbPeriodicNetworkScan); err != nil {
			log.Println("could not create network scan in DB", err)

			return
		}

		s.periodicScanMessenger.Broadcast(dbPeriodicNetworkScan.NodeScanID)
		// TODO: Add subscribe rpc for latest periodic scans; this RPC should fetch the latest scan ID from the append-only table, send the scan ID to the frontend, subscribe to the messenger and send all further ones to the latter as well
	}); err != nil {
		return err
	}

	s.cron.Run()

	return nil
}

func (s *NetworkAndNodeScanService) TriggerNetworkScan(ctx context.Context, scanTriggerMessage *proto.NetworkScanTriggerMessage) (*proto.NetworkScanReferenceMessage, error) {
	// Create a scan
	networkScan := &networkAndNodeScanModels.NetworkScan{
		Done: 0,
	}
	networkScanID, err := s.networkAndNodeScanDatabase.CreateNetworkScan(networkScan)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "could not create network scan in DB: %v", err.Error())
	}

	networkScanner := scanners.NewNetworkScanner(s.device)
	if err, _ = networkScanner.Open(); err != nil {
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

			dbNode := &networkAndNodeScanModels.Node{
				MacAddress:   node.MACAddress.String(),
				IPAddress:    node.IPAddress.String(),
				Vendor:       vendor.Vendor.String,
				Registry:     vendor.Registry,
				Organization: vendor.Organization.String,
				Address:      vendor.Address.String,
				Visible:      vendor.Visibility,
			}

			if _, err := s.networkAndNodeScanDatabase.UpsertNode(dbNode, networkScanID); err != nil {
				log.Println("could not create node in DB", err)

				break
			}

			_, err = s.startPortScan(dbNode.MacAddress, dbNode.IPAddress, networkScanID, scanTriggerMessage.GetNodeScanTimeout())
			if err != nil {
				log.Println("could not start node scan", err)

				break
			}

			networkScanMessenger.Broadcast(dbNode)
		}

		networkScanMessenger.Reset()

		networkScan.Done = 1
		if _, err := s.networkAndNodeScanDatabase.UpdateNetworkScan(networkScan); err != nil {
			log.Println("could not update network scan in DB", err)

			return
		}

		s.networkScanMessengers.Remove(string(networkScanID))
	}()

	return &proto.NetworkScanReferenceMessage{NetworkScanID: networkScanID}, nil
}

func (s *NetworkAndNodeScanService) TriggerNodeScan(ctx context.Context, nodeScanTriggerMessage *proto.NodeScanTriggerMessage) (*proto.NodeScanReferenceMessage, error) {
	node, err := s.networkAndNodeScanDatabase.GetNode(nodeScanTriggerMessage.GetMACAddress())
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "could not get nodes from DB: %v", err.Error())
	}

	nodeScanID, err := s.startPortScan(node.MacAddress, node.IPAddress, -1, nodeScanTriggerMessage.GetTimeout())
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "could not get node from DB: %v", err.Error())
	}

	protoNodeScanReferenceMessage := &proto.NodeScanReferenceMessage{
		MACAddress: node.MacAddress,
		NodeScanID: nodeScanID,
	}

	return protoNodeScanReferenceMessage, nil
}

func (s *NetworkAndNodeScanService) SubscribeToNewNodes(scanReferenceMessage *proto.NetworkScanReferenceMessage, stream proto.NetworkAndNodeScanService_SubscribeToNewNodesServer) error {
	allNodes, err := s.networkAndNodeScanDatabase.GetAllNodes()
	if err != nil {
		return status.Errorf(codes.Unknown, "could not get nodes from DB: %v", err.Error())
	}

	matchingNewestScans, err := s.networkAndNodeScanDatabase.GetNewestNetworkScansForNodes(allNodes)
	if err != nil {
		return status.Errorf(codes.Unknown, "could not get scans from DB: %v", err.Error())
	}

	var networkScan *networkAndNodeScanModels.NetworkScan
	if scanReferenceMessage.GetNetworkScanID() == -1 {
		networkScan, err = s.networkAndNodeScanDatabase.GetNewestNetworkScan()
		if err != nil {
			return status.Errorf(codes.Unknown, "could not get latest scan from DB: %v", err.Error())
		}
	} else {
		networkScan, err = s.networkAndNodeScanDatabase.GetNetworkScan(scanReferenceMessage.GetNetworkScanID())
		if err != nil {
			return status.Errorf(codes.Unknown, "could not get scan from DB: %v", err.Error())
		}
	}

	nodeScansForNetworkScanAndNode := make(map[string]int64)
	for _, dbNode := range allNodes {
		nodeScanID, err := s.networkAndNodeScanDatabase.GetNodeScanIDByNetworkScanIDAndNodeID(dbNode.MacAddress, networkScan.ID)
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
			LucidNode: &proto.NodeMetadataMessage{
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
		dbNode := receivedNode.(*networkAndNodeScanModels.Node)

		if _, ok := nodeScansForNetworkScanAndNode[dbNode.MacAddress]; !ok {
			nodeScanID, err := s.networkAndNodeScanDatabase.GetNodeScanIDByNetworkScanIDAndNodeID(dbNode.MacAddress, networkScan.ID)
			if err != nil {
				if strings.Contains(err.Error(), "sql: no rows in result set") {
					nodeScansForNetworkScanAndNode[dbNode.MacAddress] = -1
				} else {
					return status.Errorf(codes.Unknown, "could not get node scan from DB: %v", err.Error())
				}
			} else {
				nodeScansForNetworkScanAndNode[dbNode.MacAddress] = nodeScanID
			}
		}

		protoNode := &proto.DiscoveredNodeMessage{
			NodeScanID: nodeScansForNetworkScanAndNode[dbNode.MacAddress],
			LucidNode: &proto.NodeMetadataMessage{
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
	nodeScanID := nodeScanReferenceMessage.GetNodeScanID()
	if nodeScanID == -1 {
		newestNodeScanID, err := s.networkAndNodeScanDatabase.GetNewestNodeScanIDForNodeID(nodeScanReferenceMessage.GetMACAddress())
		if err != nil {
			return status.Errorf(codes.Unknown, "could not get scan ID from DB: %v", err.Error())
		}

		nodeScanID = newestNodeScanID
	}

	nodeScan, err := s.networkAndNodeScanDatabase.GetNodeScan(nodeScanID)
	if err != nil {
		return status.Errorf(codes.Unknown, "could not get scan from DB: %v", err.Error())
	}

	services, err := s.networkAndNodeScanDatabase.GetServicesForNodeScanID(nodeScanID)
	if err != nil {
		return status.Errorf(codes.Unknown, "could not get service from DB: %v", err.Error())
	}

	for _, dbService := range services {
		protoService := &proto.DiscoveredServiceMessage{
			MACAddress:              nodeScanReferenceMessage.GetMACAddress(),
			ServiceName:             dbService.ServiceName,
			PortNumber:              dbService.PortNumber,
			TransportProtocol:       dbService.TransportProtocol,
			Description:             dbService.Description,
			Assignee:                dbService.Assignee,
			Contact:                 dbService.Contact,
			RegistrationDate:        dbService.RegistrationDate,
			ModificationDate:        dbService.ModificationDate,
			Reference:               dbService.Reference,
			ServiceCode:             dbService.ServiceCode,
			UnauthorizedUseReported: dbService.UnauthorizedUseReported,
			AssignmentNotes:         dbService.AssignmentNotes,
		}

		if err := stream.Send(protoService); err != nil {
			return status.Errorf(codes.Unknown, "could not send service to frontend: %v", err.Error())
		}
	}

	if nodeScan.Done == 1 {
		return nil
	}

	msgr, exists := s.nodeScanMessengers.Get(string(nodeScan.ID))
	if !exists || msgr == nil {
		return nil
	}

	client, err := msgr.(*messenger.Messenger).Sub()
	if err != nil {
		return status.Errorf(codes.Unknown, "could not subscribe to services")
	}

	for receivedNode := range client {
		dbService := receivedNode.(*networkAndNodeScanModels.Service)

		protoService := &proto.DiscoveredServiceMessage{
			MACAddress:              nodeScanReferenceMessage.GetMACAddress(),
			ServiceName:             dbService.ServiceName,
			PortNumber:              dbService.PortNumber,
			TransportProtocol:       dbService.TransportProtocol,
			Description:             dbService.Description,
			Assignee:                dbService.Assignee,
			Contact:                 dbService.Contact,
			RegistrationDate:        dbService.RegistrationDate,
			ModificationDate:        dbService.ModificationDate,
			Reference:               dbService.Reference,
			ServiceCode:             dbService.ServiceCode,
			UnauthorizedUseReported: dbService.UnauthorizedUseReported,
			AssignmentNotes:         dbService.AssignmentNotes,
		}

		if err := stream.Send(protoService); err != nil {
			return status.Errorf(codes.Unknown, "could not send service to frontend: %v", err.Error())
		}
	}

	return nil
}

func (s *NetworkAndNodeScanService) startPortScan(nodeID string, ipAddress string, networkScanID int64, timeout int64) (int64, error) {
	// Scan for open ports for node
	// TODO: This is very expensive. The port scanners should be coordinated to run sequentially so that CPU usage isn't that high.
	portScanner := scanners.NewPortScanner(ipAddress, 0, math.MaxUint16, time.Millisecond*time.Duration(timeout), []string{"tcp", "udp"}, s.portScannerLock, func(port int) ([]byte, error) {
		packet, err := s.ports2PacketsDatabase.GetPacket(port)

		if err != nil {
			return nil, err
		}

		return packet.Packet, err
	})

	nodeScanIDChan := make(chan int64)

	// Dial and/or transmit packets ("start a scan")
	go func() {
		if err := portScanner.Transmit(); err != nil {
			log.Println("could not transmit from node scanner", err)
		}
	}()

	go func() {
		// Create a scan
		nodeScan := &networkAndNodeScanModels.NodeScan{
			Done: 0,
		}
		nodeScanID, err := s.networkAndNodeScanDatabase.CreateNodeScan(nodeScan, nodeID, networkScanID)
		if err != nil {
			log.Println("could not create node scan in DB", err)

			return
		}

		nodeScanIDChan <- nodeScanID

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

				var dbService *networkAndNodeScanModels.Service
				if len(services) > 0 {
					dbService = &networkAndNodeScanModels.Service{
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
					dbService = &networkAndNodeScanModels.Service{
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

				if _, err := s.networkAndNodeScanDatabase.UpsertService(dbService, nodeID, nodeScanID, networkScanID); err != nil {
					log.Println("could not create node in DB", err)

					break
				}

				nodeScanMessenger.Broadcast(dbService)
			}
		}

		nodeScanMessenger.Reset()

		nodeScan.Done = 1
		if _, err := s.networkAndNodeScanDatabase.UpdateNodeScan(nodeScan); err != nil {
			log.Println("could not update node scan in DB", err)

			return
		}

		s.nodeScanMessengers.Remove(string(nodeScanID))
	}()

	nodeScanID := <-nodeScanIDChan

	return nodeScanID, nil
}

func (s *NetworkAndNodeScanService) DeleteNode(ctx context.Context, nodeDeleteMessage *proto.NodeDeleteMessage) (*proto.NodeMetadataMessage, error) {
	dbNode, err := s.networkAndNodeScanDatabase.DeleteNode(nodeDeleteMessage.GetMACAddress())
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "could not get delete node from DB: %v", err.Error())
	}

	protoNode := &proto.NodeMetadataMessage{
		PoweredOn:    false, // Should not be relevant here, the node is being deleted
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
	}

	return protoNode, nil
}
