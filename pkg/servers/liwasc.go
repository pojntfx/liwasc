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
}

func NewLiwascServer(listenAddress string, networkAndNodeScanService *services.NetworkAndNodeScanService) *LiwascServer {
	return &LiwascServer{listenAddress, networkAndNodeScanService}
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

	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}
