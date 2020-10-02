package services

import (
	"context"
	"log"

	cmap "github.com/orcaman/concurrent-map"
	"github.com/pojntfx/liwasc/pkg/databases"
	proto "github.com/pojntfx/liwasc/pkg/proto/generated"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NodeWakeService struct {
	proto.UnimplementedNodeWakeServiceServer

	device             string
	nodeWakeDatabase   *databases.NodeWakeDatabase
	nodeWakeScanners   cmap.ConcurrentMap
	nodeWakeMessengers cmap.ConcurrentMap
	getIPAddress       func(string) (string, error)
}

func NewNodeWakeService(
	device string,
	nodeWakeDatabase *databases.NodeWakeDatabase,
	getIPAddress func(string) (string, error),
) *NodeWakeService {
	return &NodeWakeService{
		device:             device,
		nodeWakeDatabase:   nodeWakeDatabase,
		nodeWakeScanners:   cmap.New(),
		nodeWakeMessengers: cmap.New(),
		getIPAddress:       getIPAddress,
	}
}

func (s *NodeWakeService) TriggerNodeWake(ctx context.Context, nodeWakeTriggerMessage *proto.NodeWakeTriggerMessage) (*proto.NodeWakeReferenceMessage, error) {
	log.Println(nodeWakeTriggerMessage)

	return nil, status.Errorf(codes.Unimplemented, "method TriggerNodeWake not implemented")
}
