package services

import (
	"context"
	"log"
	"math"
	"time"

	"github.com/pojntfx/liwasc/pkg/concurrency"
	"github.com/pojntfx/liwasc/pkg/databases"
	proto "github.com/pojntfx/liwasc/pkg/proto/generated"
	"github.com/pojntfx/liwasc/pkg/scanners"
	models "github.com/pojntfx/liwasc/pkg/sql/generated/node_and_port_scan"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NodeAndPortScanPortService struct {
	proto.UnimplementedNodeAndPortScanNeoServiceServer

	device                        string
	ports2packetsDatabase         *databases.Ports2PacketDatabase
	nodeAndPortScanDatabase       *databases.NodeAndPortScanDatabase
	portScannerConcurrencyLimiter *concurrency.GoRoutineLimiter
}

func NewNodeAndPortScanPortService(
	device string,
	ports2packetsDatabase *databases.Ports2PacketDatabase,
	nodeAndPortScanDatabase *databases.NodeAndPortScanDatabase,
	portScannerConcurrencyLimiter *concurrency.GoRoutineLimiter,
) *NodeAndPortScanPortService {
	return &NodeAndPortScanPortService{
		device:                        device,
		ports2packetsDatabase:         ports2packetsDatabase,
		nodeAndPortScanDatabase:       nodeAndPortScanDatabase,
		portScannerConcurrencyLimiter: portScannerConcurrencyLimiter,
	}
}

func (s *NodeAndPortScanPortService) StartNodeScan(ctx context.Context, nodeScanStartMessage *proto.NodeScanStartNeoMessage) (*proto.NodeScanReferenceNeoMessage, error) {
	// Create node scan in DB
	dbNodeScan := &models.NodeScan{}
	if err := s.nodeAndPortScanDatabase.CreateNodeScan(dbNodeScan); err != nil {
		log.Printf("could not create node scan %v in DB: %v\n", dbNodeScan.ID, err)

		return nil, status.Errorf(codes.Unknown, "could not create node scan in DB")
	}

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

				break
			}

			// Handle node
			go func() {
				// Create node
				dbNode := &models.Node{
					NodeScanID: dbNodeScan.ID,
					MacAddress: node.MACAddress.String(),
				}

				if err := s.nodeAndPortScanDatabase.CreateNode(dbNode); err != nil {
					log.Printf("could not create node %v for node scan %v in DB: %v\n", dbNode.ID, dbNodeScan.ID, err)

					return
				}

				// Create port scan in DB
				dbPortScan := &models.PortScan{
					NodeID: dbNode.ID,
				}
				if err := s.nodeAndPortScanDatabase.CreatePortScan(dbPortScan); err != nil {
					log.Printf("could not create node scan %v for node %v for node scan %v in DB: %v\n", dbPortScan.ID, dbNode.ID, dbNodeScan.ID, err)

					return
				}

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

								// Create port
								dbPort := &models.Port{
									PortScanID:        dbNode.ID,
									PortNumber:        int64(port.Port),
									TransportProtocol: port.Protocol,
								}

								if err := s.nodeAndPortScanDatabase.CreatePort(dbPort); err != nil {
									log.Printf("could not create port %v for port scan %v for node %v for node scan %v in DB: %v\n", dbPort.ID, dbPortScan.ID, dbNode.ID, dbNodeScan.ID, err)

									return
								}
							}()
						}
					}

					dbPortScan.Done = 1
					if err := s.nodeAndPortScanDatabase.UpdatePortScan(dbPortScan); err != nil {
						log.Printf("could not update port scan %v for node %v for node scan %v in DB: %v\n", dbPortScan.ID, dbNode.ID, dbNodeScan.ID, err)
					}
				}()
			}()
		}

		dbNodeScan.Done = 1
		if err := s.nodeAndPortScanDatabase.UpdateNodeScan(dbNodeScan); err != nil {
			log.Printf("could not update node scan %v in DB: %v\n", dbNodeScan.ID, err)
		}
	}()

	// Return reference to node scan
	protoNodeScanReferenceMessage := &proto.NodeScanReferenceNeoMessage{
		ID:        dbNodeScan.ID,
		CreatedAt: dbNodeScan.CreatedAt.String(),
		Done: func() bool {
			if dbNodeScan.Done == 1 {
				return true
			}

			return false
		}(),
	}

	return protoNodeScanReferenceMessage, nil
}
