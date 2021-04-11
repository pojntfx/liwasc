package services

import (
	"context"
	"log"
	"math"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	api "github.com/pojntfx/liwasc/pkg/api/proto/v1"
	models "github.com/pojntfx/liwasc/pkg/db/sqlite/node_and_port_scan"
	"github.com/pojntfx/liwasc/pkg/persisters"
	"github.com/pojntfx/liwasc/pkg/scanners"
	"github.com/pojntfx/liwasc/pkg/validators"
	cron "github.com/robfig/cron/v3"
	"github.com/ugjka/messenger"
	"golang.org/x/sync/semaphore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NodeAndPortScanPortService struct {
	api.UnimplementedNodeAndPortScanServiceServer

	device string

	ports2packetsPersister   *persisters.Ports2PacketPersister
	nodeAndPortScanPersister *persisters.NodeAndPortScanPersister

	portScannerSemaphore *semaphore.Weighted

	nodeScanMessenger *messenger.Messenger
	nodeMessenger     *messenger.Messenger
	portScanMessenger *messenger.Messenger
	portMessenger     *messenger.Messenger

	nodeScannerLock sync.Mutex
	portScannerLock sync.Mutex

	periodicScanCronExpression string
	periodicNodeScanTimeout    int
	periodicPortScanTimeout    int

	cron *cron.Cron

	contextValidator *validators.ContextValidator
}

func NewNodeAndPortScanPortService(
	device string,

	ports2packetsPersister *persisters.Ports2PacketPersister,
	nodeAndPortScanPersister *persisters.NodeAndPortScanPersister,

	portScannerSemaphore *semaphore.Weighted,

	periodicScanCronExpression string,
	periodicNodeScanTimeout int,
	periodicPortScanTimeout int,

	contextValidator *validators.ContextValidator,
) *NodeAndPortScanPortService {
	return &NodeAndPortScanPortService{
		device: device,

		ports2packetsPersister:   ports2packetsPersister,
		nodeAndPortScanPersister: nodeAndPortScanPersister,

		portScannerSemaphore: portScannerSemaphore,

		nodeScanMessenger: messenger.New(0, true),
		nodeMessenger:     messenger.New(0, true),
		portScanMessenger: messenger.New(0, true),
		portMessenger:     messenger.New(0, true),

		periodicScanCronExpression: periodicScanCronExpression,
		periodicNodeScanTimeout:    periodicNodeScanTimeout,
		periodicPortScanTimeout:    periodicPortScanTimeout,

		cron: cron.New(),

		contextValidator: contextValidator,
	}
}

func (s *NodeAndPortScanPortService) Open() error {
	if _, err := s.cron.AddFunc(s.periodicScanCronExpression, func() {
		go func() {
			protoNodeScanStartMessage := &api.NodeScanStartMessage{
				NodeScanTimeout: int64(s.periodicNodeScanTimeout),
				PortScanTimeout: int64(s.periodicPortScanTimeout),
			}

			log.Printf("starting periodic node scan\n")

			protoNodeScanMessage, err := s.startInternalNodeScan(context.Background(), protoNodeScanStartMessage)
			if err != nil {
				log.Printf("could not start periodic node scan: %v\n", err)

				return
			}

			log.Printf("started periodic node scan %v\n", protoNodeScanMessage.GetID())
		}()
	}); err != nil {
		return err
	}

	s.cron.Run()

	return nil
}

func (s *NodeAndPortScanPortService) StartNodeScan(ctx context.Context, nodeScanStartMessage *api.NodeScanStartMessage) (*api.NodeScanMessage, error) {
	// Authorize
	valid, err := s.contextValidator.Validate(ctx)
	if err != nil || !valid {
		return nil, status.Errorf(codes.Unauthenticated, "could not authorize: %v", err)
	}

	// Validate
	if nodeScanStartMessage.GetNodeScanTimeout() < 1 {
		return nil, status.Error(codes.InvalidArgument, "node scan timeout can't be lower than 1")
	}

	if nodeScanStartMessage.GetPortScanTimeout() < 1 {
		return nil, status.Error(codes.InvalidArgument, "port scan timeout can't be lower than 1")
	}

	// Start the node scan
	return s.startInternalNodeScan(ctx, nodeScanStartMessage)
}

func (s *NodeAndPortScanPortService) startInternalNodeScan(_ context.Context, nodeScanStartMessage *api.NodeScanStartMessage) (*api.NodeScanMessage, error) {
	// Create and broadcast node scan in DB
	dbNodeScan := &models.NodeScan{}
	if err := s.nodeAndPortScanPersister.CreateNodeScan(dbNodeScan); err != nil {
		log.Printf("could not create node scan in DB: %v\n", err)

		return nil, status.Errorf(codes.Unknown, "could not create node scan in DB")
	}
	s.nodeScanMessenger.Broadcast(dbNodeScan)

	// Create and open node scanner
	nodeScanner := scanners.NewNodeScanner(s.device)
	networks, err := nodeScanner.Open()
	if err != nil {
		log.Printf("could not open node scanner for node scan %v: %v\n", dbNodeScan.ID, err)

		return nil, status.Errorf(codes.Unknown, "could not open node scanner")
	}

	// Queueing node scan
	log.Printf("queueing node scan %v for networks: %v\n", dbNodeScan.ID, networks)

	// Lock the node scanner
	nodeScannerReady := make(chan bool)
	go func() {
		// Lock the node scanner
		s.nodeScannerLock.Lock()

		// Start node scan
		log.Printf("starting node scan %v for networks: %v\n", dbNodeScan.ID, networks)

		for i := 0; i < 3; i++ {
			nodeScannerReady <- true
		}
	}()

	// Transmit node scan
	go func() {
		<-nodeScannerReady

		if err := nodeScanner.Transmit(); err != nil {
			log.Printf("could not transmit for node scan %v: %v\n", dbNodeScan.ID, err)
		}
	}()

	// Receive node scan
	go func() {
		<-nodeScannerReady

		receiveCtx, cancel := context.WithTimeout(
			context.Background(),
			time.Millisecond*time.Duration(nodeScanStartMessage.GetNodeScanTimeout()),
		)
		defer cancel()

		if err := nodeScanner.Receive(receiveCtx); err != nil {
			log.Printf("could not receive for node scan %v: %v\n", dbNodeScan.ID, err)
		}
	}()

	// Read node scan
	go func() {
		<-nodeScannerReady

		for {
			var node *scanners.DiscoveredNode
			node = nodeScanner.Read()

			// Node scan is done
			if node == nil {
				log.Printf("node scan %v is done\n", dbNodeScan.ID)

				// Broadcast node scan completion
				dbNode := &models.Node{
					NodeScanID: dbNodeScan.ID,
					MacAddress: "",
				}
				s.nodeMessenger.Broadcast(dbNode)

				break
			}

			// Handle node
			go func() {
				// Fetch/Create and broadcast node
				var dbNode *models.Node
				// Create node in DB
				dbNode = &models.Node{
					NodeScanID: dbNodeScan.ID,
					MacAddress: node.MACAddress.String(),
					IPAddress:  node.IPAddress.String(),
				}
				if err := s.nodeAndPortScanPersister.CreateNode(dbNode); err != nil {
					log.Printf("could not create node %v for node scan %v in DB: %v\n", dbNode.ID, dbNodeScan.ID, err)

					return
				}
				s.nodeMessenger.Broadcast(dbNode)

				// Create and broadcast port scan in DB
				dbPortScan := &models.PortScan{
					NodeID: dbNode.ID,
				}
				if err := s.nodeAndPortScanPersister.CreatePortScan(dbPortScan); err != nil {
					log.Printf("could not create node scan %v for node %v for node scan %v in DB: %v\n", dbPortScan.ID, dbNode.ID, dbNodeScan.ID, err)

					return
				}
				s.portScanMessenger.Broadcast(dbPortScan)

				// If non-scoped scan or node in scoped scan, do a port scan, else copy last scan's port information
				if nodeScanStartMessage.GetMACAddress() == "" || nodeScanStartMessage.GetMACAddress() == dbNode.MacAddress {
					// Create port scanner
					portscanner := scanners.NewPortScanner(
						dbNode.IPAddress,
						0,
						math.MaxUint16,
						time.Millisecond*time.Duration(nodeScanStartMessage.GetPortScanTimeout()),
						[]string{"tcp", "udp"},
						s.portScannerSemaphore,
						func(port int) ([]byte, error) {
							packet, err := s.ports2packetsPersister.GetPacket(port)
							if err != nil {
								return nil, err
							}

							return packet.Packet, nil
						},
					)

					// Queueing port scan
					log.Printf("queueing port scan %v for node %v for node scan %v\n", dbPortScan.ID, dbNode.ID, dbNodeScan.ID)

					// Lock the port scanner
					portScannerReady := make(chan bool)
					go func() {
						// Lock the port scanner
						s.portScannerLock.Lock()

						// Start port scan
						log.Printf("starting port scan %v for node %v for node scan %v\n", dbPortScan.ID, dbNode.ID, dbNodeScan.ID)

						for i := 0; i < 2; i++ {
							portScannerReady <- true
						}
					}()

					// Transmit port scan
					go func() {
						<-portScannerReady

						if err := portscanner.Transmit(); err != nil {
							log.Printf("could not transmit for port scan %v for node %v for node scan %v: %v\n", dbPortScan.ID, dbNode.ID, dbNodeScan.ID, err)
						}
					}()

					// Read port scan
					go func() {
						<-portScannerReady

						for {
							port := portscanner.Read()
							// Port scan is done
							if port == nil {
								log.Printf("port scan %v for node %v for node scan %v is done\n", dbPortScan.ID, dbNode.ID, dbNodeScan.ID)

								// Broadcast port scan completion
								dbPort := &models.Port{
									PortScanID: dbPortScan.ID,
									PortNumber: -1,
								}
								s.portMessenger.Broadcast(dbPort)

								break
							}

							// Handle port
							if port.Open {
								go func() {
									log.Printf("found open port %v/%v for port scan %v for node %v for node scan %v\n", port.Port, port.Protocol, dbPortScan.ID, dbNode.ID, dbNodeScan.ID)

									// Create and broadcast port in DB
									dbPort := &models.Port{
										PortScanID:        dbPortScan.ID,
										PortNumber:        int64(port.Port),
										TransportProtocol: port.Protocol,
									}
									if err := s.nodeAndPortScanPersister.CreatePort(dbPort); err != nil {
										log.Printf("could not create port %v for port scan %v for node %v for node scan %v in DB: %v\n", dbPort.ID, dbPortScan.ID, dbNode.ID, dbNodeScan.ID, err)

										return
									}
									s.portMessenger.Broadcast(dbPort)
								}()
							}
						}

						// Set port scan to done
						dbPortScan.Done = 1
						if err := s.nodeAndPortScanPersister.UpdatePortScan(dbPortScan); err != nil {
							log.Printf("could not update port scan %v for node %v for node scan %v in DB: %v\n", dbPortScan.ID, dbNode.ID, dbNodeScan.ID, err)
						}
						s.portScanMessenger.Broadcast(dbPortScan)

						// Unlock the scanner
						s.portScannerLock.Unlock()
					}()
				} else {
					// Get the latest port scan for this node
					latestPortScan, err := s.nodeAndPortScanPersister.GetLatestPortScanForNodeId(dbNode.MacAddress)
					if err != nil {
						log.Printf("could not get last finished port scan for node %v from DB: %v\n", dbNode.ID, err)

						return
					}

					// Get the ports for this last port scan
					dbPorts, err := s.nodeAndPortScanPersister.GetPorts(latestPortScan.ID)
					if err != nil {
						log.Printf("could not get ports for port scan %v from DB: %v\n", latestPortScan.ID, err)

						return
					}

					// Copy the ports to the new port scan
					for _, port := range dbPorts {
						// Create and broadcast port in DB
						dbPort := &models.Port{
							PortScanID:        dbPortScan.ID,
							PortNumber:        int64(port.PortNumber),
							TransportProtocol: port.TransportProtocol,
						}
						if err := s.nodeAndPortScanPersister.CreatePort(dbPort); err != nil {
							log.Printf("could not create port %v for port scan %v for node %v for node scan %v in DB: %v\n", dbPort.ID, dbPortScan.ID, dbNode.ID, dbNodeScan.ID, err)

							return
						}
						s.portMessenger.Broadcast(dbPort)
					}

					// Broadcast port scan completion
					dbPort := &models.Port{
						PortScanID: dbPortScan.ID,
						PortNumber: -1,
					}
					s.portMessenger.Broadcast(dbPort)

					// Set port scan to done
					dbPortScan.Done = 1
					if err := s.nodeAndPortScanPersister.UpdatePortScan(dbPortScan); err != nil {
						log.Printf("could not update port scan %v for node %v for node scan %v in DB: %v\n", dbPortScan.ID, dbNode.ID, dbNodeScan.ID, err)
					}
					s.portScanMessenger.Broadcast(dbPortScan)
				}
			}()
		}

		// Set node scan to done
		dbNodeScan.Done = 1
		if err := s.nodeAndPortScanPersister.UpdateNodeScan(dbNodeScan); err != nil {
			log.Printf("could not update node scan %v in DB: %v\n", dbNodeScan.ID, err)
		}
		s.nodeScanMessenger.Broadcast(dbNodeScan)

		// Unlock the node scanner
		s.nodeScannerLock.Unlock()
	}()

	// Return reference to node scan
	protoNodeScanMessage := &api.NodeScanMessage{
		ID:        dbNodeScan.ID,
		CreatedAt: dbNodeScan.CreatedAt.Format(time.RFC3339),
		Done: func() bool {
			if dbNodeScan.Done == 1 {
				return true
			}

			return false
		}(),
	}

	return protoNodeScanMessage, nil
}

func (s *NodeAndPortScanPortService) SubscribeToNodeScans(_ *empty.Empty, stream api.NodeAndPortScanService_SubscribeToNodeScansServer) error {
	// Authorize
	valid, err := s.contextValidator.Validate(stream.Context())
	if err != nil || !valid {
		return status.Errorf(codes.Unauthenticated, "could not authorize: %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(2)

	messengerReady := make(chan bool)

	// Get node scans from messenger
	go func() {
		dbNodeScans, err := s.nodeScanMessenger.Sub()
		if err != nil {
			log.Printf("could not get node scans from messenger: %v\n", err)

			return
		}
		defer s.nodeScanMessenger.Unsub(dbNodeScans)

		messengerReady <- true

		for dbNodeScan := range dbNodeScans {
			protoNodeScan := &api.NodeScanMessage{
				CreatedAt: dbNodeScan.(*models.NodeScan).CreatedAt.Format(time.RFC3339),
				Done: func() bool {
					if dbNodeScan.(*models.NodeScan).Done == 1 {
						return true
					}

					return false
				}(),
				ID: dbNodeScan.(*models.NodeScan).ID,
			}

			if err := stream.Send(protoNodeScan); err != nil {
				log.Printf("could send node scan %v to client: %v\n", protoNodeScan.ID, err)

				return
			}
		}

		wg.Done()
	}()

	// Get node scans from persister
	go func() {
		<-messengerReady

		dbNodeScans, err := s.nodeAndPortScanPersister.GetNodeScans()
		if err != nil {
			log.Printf("could not get node scans from DB: %v\n", err)

			return
		}

		for _, dbNodeScan := range dbNodeScans {
			s.nodeScanMessenger.Broadcast(dbNodeScan)
		}

		wg.Done()
	}()

	wg.Wait()

	return nil
}

func (s *NodeAndPortScanPortService) SubscribeToNodes(nodeScanMessage *api.NodeScanMessage, stream api.NodeAndPortScanService_SubscribeToNodesServer) error {
	// Authorize
	valid, err := s.contextValidator.Validate(stream.Context())
	if err != nil || !valid {
		return status.Errorf(codes.Unauthenticated, "could not authorize: %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(3)

	// Get nodes from messenger (priority 3)
	go func() {
		// Get node scan from DB and check if it is done
		dbNodeScan, err := s.nodeAndPortScanPersister.GetNodeScan(nodeScanMessage.GetID())
		if err != nil {
			log.Printf("could not get node scan %v from DB, continouing to messenger subscription: %v\n", nodeScanMessage.GetID(), err)

			dbNodeScan = &models.NodeScan{
				Done: 0,
			}
		}

		// If node scan is not done, then sub and send nodes until it is done
		if dbNodeScan.Done == 0 {
			dbNodes, err := s.nodeMessenger.Sub()
			if err != nil {
				log.Printf("could not get nodes for node scan %v from messenger: %v\n", nodeScanMessage.GetID(), err)

				return
			}
			defer s.nodeMessenger.Unsub(dbNodes)

			for dbNode := range dbNodes {
				if dbNode.(*models.Node).NodeScanID == nodeScanMessage.GetID() {
					// Node scan is done, so return
					if dbNode.(*models.Node).MacAddress == "" {
						break
					}

					protoNode := &api.NodeMessage{
						CreatedAt:  dbNode.(*models.Node).CreatedAt.Format(time.RFC3339),
						ID:         dbNode.(*models.Node).ID,
						IPAddress:  dbNode.(*models.Node).IPAddress,
						MACAddress: dbNode.(*models.Node).MacAddress,
						NodeScanID: dbNode.(*models.Node).NodeScanID,
						Priority:   3,
						PoweredOn:  true, // Found in this node scan so it has to be powered on
					}

					if err := stream.Send(protoNode); err != nil {
						log.Printf("could send node %v for node scan %v to client: %v\n", protoNode.GetID(), nodeScanMessage.GetID(), err)

						return
					}
				}
			}
		}

		wg.Done()
	}()

	// Get nodes from persister (priority 2)
	go func() {
		dbNodes, err := s.nodeAndPortScanPersister.GetNodes(nodeScanMessage.GetID())
		if err != nil {
			log.Printf("could not get nodes for node scan %v from DB: %v\n", nodeScanMessage.GetID(), err)

			return
		}

		for _, dbNode := range dbNodes {
			protoNode := &api.NodeMessage{
				CreatedAt:  dbNode.CreatedAt.Format(time.RFC3339),
				ID:         dbNode.ID,
				IPAddress:  dbNode.IPAddress,
				MACAddress: dbNode.MacAddress,
				NodeScanID: dbNode.NodeScanID,
				Priority:   2,
				PoweredOn:  true, // Found in this node scan so it has to be powered on
			}

			if err := stream.Send(protoNode); err != nil {
				log.Printf("could send node %v for node scan %v to client: %v\n", protoNode.GetID(), nodeScanMessage.GetID(), err)

				return
			}
		}

		wg.Done()
	}()

	// Get lookback nodes from persister (priority 1)
	go func() {
		dbNodes, err := s.nodeAndPortScanPersister.GetLookbackNodes()
		if err != nil {
			log.Printf("could not get lookback nodes from DB: %v\n", err)

			return
		}

		for _, dbNode := range dbNodes {
			protoNode := &api.NodeMessage{
				CreatedAt:  dbNode.CreatedAt.Format(time.RFC3339),
				ID:         dbNode.ID,
				IPAddress:  dbNode.IPAddress,
				MACAddress: dbNode.MacAddress,
				NodeScanID: dbNode.NodeScanID,
				Priority:   1,
				PoweredOn:  false, // Not found in this node scan, so set to false. Higher priorities will overwrite this.
			}

			if err := stream.Send(protoNode); err != nil {
				log.Printf("could send lookback node %v to client: %v\n", protoNode.GetID(), err)

				return
			}
		}

		wg.Done()
	}()

	wg.Wait()

	return nil
}

func (s *NodeAndPortScanPortService) SubscribeToPortScans(nodeMessage *api.NodeMessage, stream api.NodeAndPortScanService_SubscribeToPortScansServer) error {
	// Authorize
	valid, err := s.contextValidator.Validate(stream.Context())
	if err != nil || !valid {
		return status.Errorf(codes.Unauthenticated, "could not authorize: %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(2)

	messengerReady := make(chan bool)

	// Get port scans from messenger
	go func() {
		dbPortScans, err := s.portScanMessenger.Sub()
		if err != nil {
			log.Printf("could not get port scans from messenger: %v\n", err)

			return
		}
		defer s.portScanMessenger.Unsub(dbPortScans)

		messengerReady <- true

		newestPortScanDate := time.Unix(0, 0)
		for dbPortScan := range dbPortScans {
			if dbPortScan.(*models.PortScan).NodeID == nodeMessage.GetID() {
				protoPortScan := &api.PortScanMessage{
					CreatedAt: dbPortScan.(*models.PortScan).CreatedAt.Format(time.RFC3339),
					Done: func() bool {
						if dbPortScan.(*models.PortScan).Done == 1 {
							return true
						}

						return false
					}(),
					ID:     dbPortScan.(*models.PortScan).ID,
					NodeID: dbPortScan.(*models.PortScan).NodeID,
				}

				if err := stream.Send(protoPortScan); err != nil {
					log.Printf("could send port scan %v to client: %v\n", protoPortScan.ID, err)

					return
				}

				// If this is the first iteration, init with the initial date
				// Works because a) messenger is non-blocking and b) messenger delivers in order, so newest
				// messages come first
				if newestPortScanDate == time.Unix(0, 0) {
					newestPortScanDate = dbPortScan.(*models.PortScan).CreatedAt
				}

				// There can only be one port scan per node, so if at least one port scan is done, break.
				if protoPortScan.Done && (dbPortScan.(*models.PortScan).CreatedAt.After(newestPortScanDate) || dbPortScan.(*models.PortScan).CreatedAt.Equal(newestPortScanDate)) {
					break
				}
			}
		}

		wg.Done()
	}()

	// Get port scans from persister
	go func() {
		<-messengerReady

		dbPortScans, err := s.nodeAndPortScanPersister.GetPortScans(nodeMessage.GetID())
		if err != nil {
			log.Printf("could not get port scans from DB: %v\n", err)

			return
		}

		for _, dbPortScan := range dbPortScans {
			s.portScanMessenger.Broadcast(dbPortScan)
		}

		wg.Done()
	}()

	wg.Wait()

	return nil
}

func (s *NodeAndPortScanPortService) SubscribeToPorts(portScanMessage *api.PortScanMessage, stream api.NodeAndPortScanService_SubscribeToPortsServer) error {
	// Authorize
	valid, err := s.contextValidator.Validate(stream.Context())
	if err != nil || !valid {
		return status.Errorf(codes.Unauthenticated, "could not authorize: %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(2)

	// Get ports from messenger (priority 2)
	go func() {
		// Get port scan from DB and check if it is done
		dbPortScan, err := s.nodeAndPortScanPersister.GetPortScan(portScanMessage.GetID())
		if err != nil {
			log.Printf("could not get port scan %v from DB, continouing to messenger subscription: %v\n", portScanMessage.GetID(), err)

			dbPortScan = &models.PortScan{
				Done: 0,
			}
		}

		// If port scan is not done, then sub and send ports until it is done
		if dbPortScan.Done == 0 {
			dbPorts, err := s.portMessenger.Sub()
			if err != nil {
				log.Printf("could not get ports for port scan %v from messenger: %v\n", portScanMessage.GetID(), err)

				return
			}
			defer s.portMessenger.Unsub(dbPorts)

			for dbPort := range dbPorts {
				if dbPort.(*models.Port).PortScanID == portScanMessage.GetID() {
					// Port scan is done, so return
					if dbPort.(*models.Port).PortNumber == -1 {
						break
					}

					protoPort := &api.PortMessage{
						CreatedAt:         dbPort.(*models.Port).CreatedAt.Format(time.RFC3339),
						ID:                dbPort.(*models.Port).ID,
						Priority:          2,
						PortNumber:        dbPort.(*models.Port).PortNumber,
						PortScanID:        dbPort.(*models.Port).PortScanID,
						TransportProtocol: dbPort.(*models.Port).TransportProtocol,
					}

					if err := stream.Send(protoPort); err != nil {
						log.Printf("could send port %v for port scan %v to client: %v\n", protoPort.GetID(), portScanMessage.GetID(), err)

						return
					}
				}
			}
		}

		wg.Done()
	}()

	// Get ports from persister (priority 1)
	go func() {
		dbPorts, err := s.nodeAndPortScanPersister.GetPorts(portScanMessage.GetID())
		if err != nil {
			log.Printf("could not get ports for port scan %v from DB: %v\n", portScanMessage.GetID(), err)

			return
		}

		for _, dbPort := range dbPorts {
			protoPort := &api.PortMessage{
				CreatedAt:         dbPort.CreatedAt.Format(time.RFC3339),
				ID:                dbPort.ID,
				Priority:          1,
				PortNumber:        dbPort.PortNumber,
				PortScanID:        dbPort.PortScanID,
				TransportProtocol: dbPort.TransportProtocol,
			}

			if err := stream.Send(protoPort); err != nil {
				log.Printf("could send port %v for port scan %v to client: %v\n", protoPort.GetID(), portScanMessage.GetID(), err)

				return
			}
		}

		wg.Done()
	}()

	wg.Wait()

	return nil
}
