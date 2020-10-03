package main

import (
	"time"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/go-app-grpc-chat-frontend-web/pkg/websocketproxy"
	"github.com/pojntfx/liwasc-frontend-web/pkg/components"
	proto "github.com/pojntfx/liwasc-frontend-web/pkg/proto/generated"
	"google.golang.org/grpc"
)

func main() {
	proxy := websocketproxy.NewWebSocketProxyClient(time.Minute)

	conn, err := grpc.Dial("ws://stuttgart.felicitas.pojtinger.com:15124", grpc.WithContextDialer(proxy.Dialer), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	networkAndNodeScanServiceClient := proto.NewNetworkAndNodeScanServiceClient(conn)
	nodeWakeServiceClient := proto.NewNodeWakeServiceClient(conn)

	app.Route("/", &components.DataProviderComponent{
		NetworkAndNodeScanServiceClient: networkAndNodeScanServiceClient,
		NodeWakeServiceClient:           nodeWakeServiceClient,
		Children: func(dataProviderChildrenProps components.DataProviderChildrenProps) app.UI {
			return &components.AppComponent{
				UserMenuOpen: false,
				UserAvatar:   "https://www.gravatar.com/avatar/db856df33fa4f4bce441819f604c90d5",
				UserName:     "Felicitas Pojtinger",

				Subnets:         []string{"10.0.0.0/9", "192.168.0.0/27"},
				Device:          "eth0",
				NodeSearchValue: "",

				Nodes:        dataProviderChildrenProps.Nodes,
				SelectedNode: -1,

				InspectorOpen:        false,
				ServicesAndPortsOpen: true,
				DetailsOpen:          false,

				ServicesOpen:         false,
				SelectedService:      -1,
				InspectorSearchValue: "",
			}
		},
	})

	app.Run()
}
