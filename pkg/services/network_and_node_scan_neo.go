package services

import (
	"context"
	"log"

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

func (s *NetworkAndNodeScanNeoService) StartNetworkScan(ctx context.Context, NetworkScanNeoStartMessage *proto.NetworkScanNeoStartMessage) (*proto.NetworkScanNeoReferenceMessage, error) {
	// Create network scan in DB
	dbNetworkScan := &models.NetworkScan{}
	if err := s.networkAndNodeScanNeoDatabase.CreateNetworkScan(dbNetworkScan); err != nil {
		if err != nil {
			log.Println("could not create network scan in DB", err)

			return nil, status.Errorf(codes.Unknown, "could not create network scan in DB")
		}
	}

	// Create network scanner
	networkScanner := scanners.NewNetworkScanner(s.device)

	networks, err := networkScanner.Open()
	if err != nil {
		log.Println("could not open network scanner", err)

		return nil, status.Errorf(codes.Unknown, "could not open network scanner")
	}

	// Start network scan
	log.Printf("starting network scan %v for networks: %v\n", dbNetworkScan.ID, networks)

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
