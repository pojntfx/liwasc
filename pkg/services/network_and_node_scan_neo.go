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
	models "github.com/pojntfx/liwasc/pkg/sql/generated/network_and_node_scan_neo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NetworkAndNodeScanNeoService struct {
	proto.UnimplementedNetworkAndNodeScanNeoServiceServer

	device                        string
	ports2packetsDatabase         *databases.Ports2PacketDatabase
	networkAndNodeScanNeoDatabase *databases.NetworkAndNodeScanNeoDatabase
	portScannerConcurrencyLimiter *concurrency.GoRoutineLimiter
}

func NewNetworkAndNodeScanNeoService(
	device string,
	ports2packetsDatabase *databases.Ports2PacketDatabase,
	networkAndNodeScanNeoDatabase *databases.NetworkAndNodeScanNeoDatabase,
	portScannerConcurrencyLimiter *concurrency.GoRoutineLimiter,
) *NetworkAndNodeScanNeoService {
	return &NetworkAndNodeScanNeoService{
		device:                        device,
		ports2packetsDatabase:         ports2packetsDatabase,
		networkAndNodeScanNeoDatabase: networkAndNodeScanNeoDatabase,
		portScannerConcurrencyLimiter: portScannerConcurrencyLimiter,
	}
}

func (s *NetworkAndNodeScanNeoService) StartNetworkScan(ctx context.Context, networkScanNeoStartMessage *proto.NetworkScanNeoStartMessage) (*proto.NetworkScanNeoReferenceMessage, error) {
	// Create network scan in DB
	dbNetworkScan := &models.NetworkScan{}
	if err := s.networkAndNodeScanNeoDatabase.CreateNetworkScan(dbNetworkScan); err != nil {
		log.Printf("could not create network scan %v in DB: %v\n", dbNetworkScan.ID, err)

		return nil, status.Errorf(codes.Unknown, "could not create network scan in DB")
	}

	// Create and open network scanner
	networkScanner := scanners.NewNetworkScanner(s.device)
	networks, err := networkScanner.Open()
	if err != nil {
		log.Printf("could not open network scanner for network scan %v: %v\n", dbNetworkScan.ID, err)

		return nil, status.Errorf(codes.Unknown, "could not open network scanner")
	}

	// Start network scan
	log.Printf("starting network scan %v for networks: %v\n", dbNetworkScan.ID, networks)

	// Transmit network scan
	go func() {
		if err := networkScanner.Transmit(); err != nil {
			log.Printf("could not transmit for network scan %v: %v\n", dbNetworkScan.ID, err)
		}
	}()

	// Receive network scan
	go func() {
		receiveCtx, cancel := context.WithTimeout(
			context.Background(),
			time.Millisecond*time.Duration(networkScanNeoStartMessage.GetNetworkScanTimeout()),
		)
		defer cancel()

		if err := networkScanner.Receive(receiveCtx); err != nil {
			log.Printf("could not receive for network scan %v: %v\n", dbNetworkScan.ID, err)
		}
	}()

	// Read network scan
	go func() {
		for {
			node := networkScanner.Read()
			// Network scan is done
			if node == nil {
				log.Printf("network scan %v is done\n", dbNetworkScan.ID)

				break
			}

			// Handle node
			go func() {
				// Create node
				dbNode := &models.Node{
					NetworkScanID: dbNetworkScan.ID,
					MacAddress:    node.MACAddress.String(),
				}

				if err := s.networkAndNodeScanNeoDatabase.CreateNode(dbNode); err != nil {
					log.Printf("could not create node %v for network scan %v in DB: %v\n", dbNode.ID, dbNetworkScan.ID, err)

					return
				}

				// Create node scan in DB
				dbNodeScan := &models.NodeScan{
					NodeID: dbNode.ID,
				}
				if err := s.networkAndNodeScanNeoDatabase.CreateNodeScan(dbNodeScan); err != nil {
					log.Printf("could not create node scan %v for node %v for network scan %v in DB: %v\n", dbNodeScan.ID, dbNode.ID, dbNetworkScan.ID, err)

					return
				}

				// Create node scanner
				nodeScanner := scanners.NewPortScanner(
					node.IPAddress.String(),
					0,
					math.MaxInt16,
					time.Millisecond*time.Duration(networkScanNeoStartMessage.GetNodeScanTimeout()),
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

				// Start node scan
				log.Printf("starting node scan %v for node %v for network scan: %v\n", dbNodeScan.ID, dbNode.ID, dbNetworkScan.ID)

				// Transmit node scan
				go func() {
					if err := nodeScanner.Transmit(); err != nil {
						log.Printf("could not transmit for node scan %v for node %v for network scan %v: %v\n", dbNodeScan.ID, dbNode.ID, dbNetworkScan.ID, err)
					}
				}()

				// Read node scan
				go func() {
					for {
						port := nodeScanner.Read()
						// Node scan is done
						if port == nil {
							log.Printf("node scan %v for network scan %v is done\n", dbNodeScan.ID, dbNetworkScan.ID)

							break
						}

						// Handle port
						if port.Open {
							go func() {
								log.Printf("found open port %v/%v for node %v for network scan %v\n", port.Port, port.Protocol, dbNode.ID, dbNetworkScan.ID)

								// Create service
								dbService := &models.Service{
									NodeScanID:        dbNode.ID,
									PortNumber:        int64(port.Port),
									TransportProtocol: port.Protocol,
								}

								if err := s.networkAndNodeScanNeoDatabase.CreateService(dbService); err != nil {
									log.Printf("could not create service %v for node scan %v for node %v for network scan %v in DB: %v\n", dbService.ID, dbNodeScan.ID, dbNode.ID, dbNetworkScan.ID, err)

									return
								}
							}()
						}
					}

					dbNodeScan.Done = 1
					if err := s.networkAndNodeScanNeoDatabase.UpdateNodeScan(dbNodeScan); err != nil {
						log.Printf("could not update node scan %v in DB: %v\n", dbNodeScan.ID, err)
					}
				}()
			}()
		}

		dbNetworkScan.Done = 1
		if err := s.networkAndNodeScanNeoDatabase.UpdateNetworkScan(dbNetworkScan); err != nil {
			log.Printf("could not update network scan %v in DB: %v\n", dbNetworkScan.ID, err)
		}
	}()

	// Return reference to network scan
	protoNetworkScanReferenceMessage := &proto.NetworkScanNeoReferenceMessage{
		ID:        dbNetworkScan.ID,
		CreatedAt: dbNetworkScan.CreatedAt.String(),
		Done: func() bool {
			if dbNetworkScan.Done == 1 {
				return true
			}

			return false
		}(),
	}

	return protoNetworkScanReferenceMessage, nil
}
