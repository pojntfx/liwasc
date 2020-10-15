package servers

import (
	"net"
	"sync"

	"github.com/pojntfx/go-app-grpc-chat-backend/pkg/websocketproxy"
	proto "github.com/pojntfx/liwasc/pkg/proto/generated"
	"github.com/pojntfx/liwasc/pkg/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type LiwascServer struct {
	listenAddress          string
	webSocketListenAddress string

	nodeAndPortScanService *services.NodeAndPortScanPortService
	metadataService     *services.MetadataService
	nodeWakeService     *services.NodeWakeService
}

func NewLiwascServer(
	listenAddress string,
	webSocketListenAddress string,

	nodeAndPortScanService *services.NodeAndPortScanPortService,
	metadataService *services.MetadataService,
	nodeWakeService *services.NodeWakeService,
) *LiwascServer {
	return &LiwascServer{
		listenAddress:          listenAddress,
		webSocketListenAddress: webSocketListenAddress,

		nodeAndPortScanService: nodeAndPortScanService,
		metadataService:     metadataService,
		nodeWakeService:     nodeWakeService,
	}
}

func (s *LiwascServer) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.listenAddress)
	if err != nil {
		return err
	}

	proxy := websocketproxy.NewWebSocketProxyServer(s.webSocketListenAddress)
	webSocketListener, err := proxy.Listen()
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	reflection.Register(server)
	proto.RegisterNodeAndPortScanServiceServer(server, s.nodeAndPortScanService)
	proto.RegisterMetadataServiceServer(server, s.metadataService)
	proto.RegisterNodeWakeServiceServer(server, s.nodeWakeService)

	doneChan := make(chan struct{})
	errChan := make(chan error)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		wg.Wait()

		close(doneChan)
	}()

	go func() {
		if err := server.Serve(listener); err != nil {
			errChan <- err
		}

		wg.Done()
	}()

	go func() {
		if err := server.Serve(webSocketListener); err != nil {
			errChan <- err
		}

		wg.Done()
	}()

	select {
	case <-doneChan:
		return nil
	case err := <-errChan:
		return err
	}
}
