package services

import (
	"context"
	"log"
	"math"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pojntfx/liwasc/pkg/concurrency"
	"github.com/pojntfx/liwasc/pkg/databases"
	proto "github.com/pojntfx/liwasc/pkg/proto/generated"
	"github.com/pojntfx/liwasc/pkg/scanners"
	models "github.com/pojntfx/liwasc/pkg/sql/generated/node_and_port_scan"
	"github.com/ugjka/messenger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NodeAndPortScanPortService struct {
	proto.UnimplementedNodeAndPortScanNeoServiceServer

	device string

	ports2packetsDatabase   *databases.Ports2PacketDatabase
	nodeAndPortScanDatabase *databases.NodeAndPortScanDatabase

	portScannerConcurrencyLimiter *concurrency.GoRoutineLimiter

	nodeScanMessenger *messenger.Messenger
	nodeMessenger     *messenger.Messenger
	portScanMessenger *messenger.Messenger
	portMessenger     *messenger.Messenger
}

func NewNodeAndPortScanPortService(
	device string,
	ports2packetsDatabase *databases.Ports2PacketDatabase,
	nodeAndPortScanDatabase *databases.NodeAndPortScanDatabase,
	portScannerConcurrencyLimiter *concurrency.GoRoutineLimiter,
) *NodeAndPortScanPortService {
	return &NodeAndPortScanPortService{
		device: device,

		ports2packetsDatabase:   ports2packetsDatabase,
		nodeAndPortScanDatabase: nodeAndPortScanDatabase,

		portScannerConcurrencyLimiter: portScannerConcurrencyLimiter,

		nodeScanMessenger: messenger.New(0, true),
		nodeMessenger:     messenger.New(0, true),
		portScanMessenger: messenger.New(0, true),
		portMessenger:     messenger.New(0, true),
	}
}

func (s *NodeAndPortScanPortService) StartNodeScan(ctx context.Context, nodeScanStartMessage *proto.NodeScanStartNeoMessage) (*proto.NodeScanNeoMessage, error) {
	// Create and broadcast node scan in DB
	dbNodeScan := &models.NodeScan{}
	if err := s.nodeAndPortScanDatabase.CreateNodeScan(dbNodeScan); err != nil {
		log.Printf("could not create node scan %v in DB: %v\n", dbNodeScan.ID, err)

		return nil, status.Errorf(codes.Unknown, "could not create node scan in DB")
	}
	s.nodeScanMessenger.Broadcast(dbNodeScan)

	// Create and open node scanner
	nodeScanner := scanners.NewNodeScanner(s.device)
	nodes, err := nodeScanner.Open()
	if err != nil {
		log.Printf("could not open node scanner for node scan %v: %v\n", dbNodeScan.ID, err)

		return nil, status.Errorf(codes.Unknown, "could not open node scanner")
	}

	// Start node scan
	log.Printf("starting node scan %v for nodes: %v\n", dbNodeScan.ID, nodes)

	// Transmit node scan
	go func() {
		if err := nodeScanner.Transmit(); err != nil {
			log.Printf("could not transmit for node scan %v: %v\n", dbNodeScan.ID, err)
		}
	}()

	// Receive node scan
	go func() {
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
		for {
			node := nodeScanner.Read()
			// Node scan is done
			if node == nil {
				log.Printf("node scan %v is done\n", dbNodeScan.ID)

				// Broadcast node scan completion
				dbNode := &models.Node{
					NodeScanID: dbNodeScan.ID,
					MacAddress: "-1",
				}
				s.nodeMessenger.Broadcast(dbNode)

				break
			}

			// Handle node
			go func() {
				// Create and broadcast node
				dbNode := &models.Node{
					NodeScanID: dbNodeScan.ID,
					MacAddress: node.MACAddress.String(),
					IPAddress:  node.IPAddress.String(),
				}
				if err := s.nodeAndPortScanDatabase.CreateNode(dbNode); err != nil {
					log.Printf("could not create node %v for node scan %v in DB: %v\n", dbNode.ID, dbNodeScan.ID, err)

					return
				}
				s.nodeMessenger.Broadcast(dbNode)

				// Create and broadcast port scan in DB
				dbPortScan := &models.PortScan{
					NodeID: dbNode.ID,
				}
				if err := s.nodeAndPortScanDatabase.CreatePortScan(dbPortScan); err != nil {
					log.Printf("could not create node scan %v for node %v for node scan %v in DB: %v\n", dbPortScan.ID, dbNode.ID, dbNodeScan.ID, err)

					return
				}
				s.portScanMessenger.Broadcast(dbPortScan)

				// Create port scanner
				portscanner := scanners.NewPortScanner(
					node.IPAddress.String(),
					0,
					math.MaxInt16,
					time.Millisecond*time.Duration(nodeScanStartMessage.GetPortScanTimeout()),
					[]string{"tcp", "udp"},
					s.portScannerConcurrencyLimiter,
					func(port int) ([]byte, error) {
						packet, err := s.ports2packetsDatabase.GetPacket(port)
						if err != nil {
							return nil, err
						}

						return packet.Packet, nil
					},
				)

				// Start port scan
				log.Printf("starting port scan %v for node %v for node scan %v\n", dbPortScan.ID, dbNode.ID, dbNodeScan.ID)

				// Transmit port scan
				go func() {
					if err := portscanner.Transmit(); err != nil {
						log.Printf("could not transmit for port scan %v for node %v for node scan %v: %v\n", dbPortScan.ID, dbNode.ID, dbNodeScan.ID, err)
					}
				}()

				// Read port scan
				go func() {
					for {
						port := portscanner.Read()
						// Port scan is done
						if port == nil {
							log.Printf("port scan %v for node %v for node scan %v is done\n", dbPortScan.ID, dbNode.ID, dbNodeScan.ID)

							break
						}

						// Handle port
						if port.Open {
							go func() {
								log.Printf("found open port %v/%v for port scan %v for node %v for node scan %v\n", port.Port, port.Protocol, dbPortScan.ID, dbNode.ID, dbNodeScan.ID)

								// Create and broadcast port in DB
								dbPort := &models.Port{
									PortScanID:        dbNode.ID,
									PortNumber:        int64(port.Port),
									TransportProtocol: port.Protocol,
								}
								if err := s.nodeAndPortScanDatabase.CreatePort(dbPort); err != nil {
									log.Printf("could not create port %v for port scan %v for node %v for node scan %v in DB: %v\n", dbPort.ID, dbPortScan.ID, dbNode.ID, dbNodeScan.ID, err)

									return
								}
								s.portMessenger.Broadcast(dbPort)
							}()
						}
					}

					// Set port scan to done
					dbPortScan.Done = 1
					if err := s.nodeAndPortScanDatabase.UpdatePortScan(dbPortScan); err != nil {
						log.Printf("could not update port scan %v for node %v for node scan %v in DB: %v\n", dbPortScan.ID, dbNode.ID, dbNodeScan.ID, err)
					}
					s.portScanMessenger.Broadcast(dbPortScan)
				}()
			}()
		}

		// Set node scan to done
		dbNodeScan.Done = 1
		if err := s.nodeAndPortScanDatabase.UpdateNodeScan(dbNodeScan); err != nil {
			log.Printf("could not update node scan %v in DB: %v\n", dbNodeScan.ID, err)
		}
		s.nodeScanMessenger.Broadcast(dbNodeScan)
	}()

	// Return reference to node scan
	protoNodeScanMessage := &proto.NodeScanNeoMessage{
		ID:        dbNodeScan.ID,
		CreatedAt: dbNodeScan.CreatedAt.String(),
		Done: func() bool {
			if dbNodeScan.Done == 1 {
				return true
			}

			return false
		}(),
	}

	return protoNodeScanMessage, nil
}

func (s *NodeAndPortScanPortService) SubscribeToNodeScans(_ *empty.Empty, stream proto.NodeAndPortScanNeoService_SubscribeToNodeScansServer) error {
	var wg sync.WaitGroup

	wg.Add(2)

	// Get node scans from messenger
	go func() {
		dbNodeScans, err := s.nodeScanMessenger.Sub()
		if err != nil {
			log.Printf("could not get node scans from messenger: %v\n", err)

			return
		}

		for dbNodeScan := range dbNodeScans {
			protoNodeScan := &proto.NodeScanNeoMessage{
				CreatedAt: dbNodeScan.(*models.NodeScan).CreatedAt.String(),
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

	// Get node scans from database
	go func() {
		dbNodeScans, err := s.nodeAndPortScanDatabase.GetNodeScans()
		if err != nil {
			log.Printf("could not get node scans from DB: %v\n", err)

			return
		}

		for _, dbNodeScan := range dbNodeScans {
			protoNodeScan := &proto.NodeScanNeoMessage{
				CreatedAt: dbNodeScan.CreatedAt.String(),
				Done: func() bool {
					if dbNodeScan.Done == 1 {
						return true
					}

					return false
				}(),
				ID: dbNodeScan.ID,
			}

			if err := stream.Send(protoNodeScan); err != nil {
				log.Printf("could send node scan %v to client: %v\n", protoNodeScan.ID, err)

				return
			}
		}

		wg.Done()
	}()

	wg.Wait()

	return nil
}

func (s *NodeAndPortScanPortService) SubscribeToNodes(nodeScanMessage *proto.NodeScanNeoMessage, stream proto.NodeAndPortScanNeoService_SubscribeToNodesServer) error {
	var wg sync.WaitGroup

	wg.Add(3)

	// Get nodes from messenger (priority 1)
	go func() {
		// Get node scan from DB and check if it is done
		dbNodeScan, err := s.nodeAndPortScanDatabase.GetNodeScan(nodeScanMessage.GetID())
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

			for dbNode := range dbNodes {
				if dbNode.(*models.Node).NodeScanID == nodeScanMessage.GetID() {
					// Node scan is done, so return
					if dbNode.(*models.Node).MacAddress == "-1" {
						break
					}

					protoNode := &proto.NodeNeoMessage{
						CreatedAt:  dbNode.(*models.Node).CreatedAt.String(),
						ID:         dbNode.(*models.Node).ID,
						IPAddress:  dbNode.(*models.Node).IPAddress,
						MACAddress: dbNode.(*models.Node).MacAddress,
						NodeScanID: dbNode.(*models.Node).NodeScanID,
						Priority:   1,
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

	// Get nodes from database (priority 2)
	go func() {
		dbNodes, err := s.nodeAndPortScanDatabase.GetNodes(nodeScanMessage.GetID())
		if err != nil {
			log.Printf("could not get nodes for node scan %v from DB: %v\n", nodeScanMessage.GetID(), err)

			return
		}

		for _, dbNode := range dbNodes {
			protoNode := &proto.NodeNeoMessage{
				CreatedAt:  dbNode.CreatedAt.String(),
				ID:         dbNode.ID,
				IPAddress:  dbNode.IPAddress,
				MACAddress: dbNode.MacAddress,
				NodeScanID: dbNode.NodeScanID,
				Priority:   2,
			}

			if err := stream.Send(protoNode); err != nil {
				log.Printf("could send node %v for node scan %v to client: %v\n", protoNode.GetID(), nodeScanMessage.GetID(), err)

				return
			}
		}

		wg.Done()
	}()

	// Get lookback nodes from database (priority 3)
	go func() {
		dbNodes, err := s.nodeAndPortScanDatabase.GetLookbackNodes()
		if err != nil {
			log.Printf("could not get lookback nodes from DB: %v\n", err)

			return
		}

		for _, dbNode := range dbNodes {
			protoNode := &proto.NodeNeoMessage{
				CreatedAt:  dbNode.CreatedAt.String(),
				ID:         dbNode.ID,
				IPAddress:  dbNode.IPAddress,
				MACAddress: dbNode.MacAddress,
				NodeScanID: dbNode.NodeScanID,
				Priority:   3,
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
