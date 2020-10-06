package main

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/go-app-grpc-chat-frontend-web/pkg/websocketproxy"
	"github.com/pojntfx/liwasc-frontend-web/pkg/components"
	proto "github.com/pojntfx/liwasc-frontend-web/pkg/proto/generated"
	"google.golang.org/grpc"
)

func main() {
	proxy := websocketproxy.NewWebSocketProxyClient(time.Minute)

	conn, err := grpc.Dial("ws://stuttgart.felix.pojtinger.com:15124", grpc.WithContextDialer(proxy.Dialer), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	networkAndNodeScanServiceClient := proto.NewNetworkAndNodeScanServiceClient(conn)
	nodeWakeServiceClient := proto.NewNodeWakeServiceClient(conn)
	metadataServiceClient := proto.NewMetadataServiceClient(conn)

	app.Route("/",
		&components.OIDCLoginProviderComponent{
			Issuer:       app.Getenv("LIWASC_OIDC_ISSUER"),
			ClientID:     app.Getenv("LIWASC_OIDC_CLIENT_ID"),
			ClientSecret: app.Getenv("LIWASC_OIDC_CLIENT_SECRET"),
			RedirectURL:  app.Getenv("LIWASC_OIDC_REDIRECT_URL"),
			HomePath:     "/",
			Scopes:       []string{"profile", "email"},

			LocalStoragePrefix: "liwasc",

			Children: func(loginProviderChildrenProps components.OIDCLoginProviderChildrenProps) app.UI {
				// TODO: Add better error and login displays
				if loginProviderChildrenProps.Error != nil {
					return app.Text(fmt.Sprintf("An error occurred: %v", err))
				}

				if loginProviderChildrenProps.OAuth2Token.AccessToken == "" || loginProviderChildrenProps.UserInfo.Email == "" {
					return app.Text("Logging you in ...")
				}

				return &components.DataProviderComponent{
					AccessToken: loginProviderChildrenProps.OAuth2Token.AccessToken,

					NetworkAndNodeScanServiceClient: networkAndNodeScanServiceClient,
					NodeWakeServiceClient:           nodeWakeServiceClient,
					MetadataServiceClient:           metadataServiceClient,
					Children: func(dataProviderChildrenProps components.DataProviderChildrenProps) app.UI {
						return &components.AppComponent{
							UserAvatar: fmt.Sprintf("https://www.gravatar.com/avatar/%x", md5.Sum([]byte(loginProviderChildrenProps.UserInfo.Email))),
							UserName:   loginProviderChildrenProps.UserInfo.Email,

							Logout: loginProviderChildrenProps.Logout,

							Subnets:         dataProviderChildrenProps.Subnets,
							Device:          dataProviderChildrenProps.Device,
							NodeSearchValue: "",

							Nodes: dataProviderChildrenProps.Nodes,

							InspectorSearchValue: "",

							Connected: dataProviderChildrenProps.Connected,
							Scanning:  dataProviderChildrenProps.Scanning,
						}
					},
				}
			},
		},
	)

	app.Run()
}
