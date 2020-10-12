package services

import (
	"context"

	"github.com/pojntfx/liwasc/pkg/concurrency"
	"github.com/pojntfx/liwasc/pkg/databases"
	proto "github.com/pojntfx/liwasc/pkg/proto/generated"
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
	return nil, status.Errorf(codes.Unimplemented, "method StartNetworkScan not implemented")
}
