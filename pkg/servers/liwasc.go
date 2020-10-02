package servers

import (
	"net"

	proto "github.com/pojntfx/liwasc/pkg/proto/generated"
	"github.com/pojntfx/liwasc/pkg/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type LiwascServer struct {
	listenAddress             string
	networkAndNodeScanService *services.NetworkAndNodeScanService
	nodeWakeService           *services.NodeWakeService
}

func NewLiwascServer(listenAddress string, networkAndNodeScanService *services.NetworkAndNodeScanService, nodeWakeService *services.NodeWakeService) *LiwascServer {
	return &LiwascServer{listenAddress, networkAndNodeScanService, nodeWakeService}
}

func (s *LiwascServer) Open() error {
	listenAddress, err := net.ResolveTCPAddr("tcp", s.listenAddress)
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", listenAddress)
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	reflection.Register(server)
	proto.RegisterNetworkAndNodeScanServiceServer(server, s.networkAndNodeScanService)
	proto.RegisterNodeWakeServiceServer(server, s.nodeWakeService)

	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}
