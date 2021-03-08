package main

import (
	"context"
	"time"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/go-app-grpc-chat-frontend-web/pkg/websocketproxy"
	"github.com/pojntfx/liwasc-frontend-web/pkg/components"
	"github.com/pojntfx/liwasc-frontend-web/pkg/components/experimental"
	proto "github.com/pojntfx/liwasc-frontend-web/pkg/proto/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	proxy := websocketproxy.NewWebSocketProxyClient(time.Minute)

	conn, err := grpc.Dial("ws://localhost:15124", grpc.WithContextDialer(proxy.Dialer), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	metadataService := proto.NewMetadataServiceClient(conn)
	nodeAndPortScanService := proto.NewNodeAndPortScanServiceClient(conn)

	app.Route("/",
		&components.OIDCLoginProviderComponent{
			Issuer:      app.Getenv("LIWASC_OIDC_ISSUER"),
			ClientID:    app.Getenv("LIWASC_OIDC_CLIENT_ID"),
			RedirectURL: app.Getenv("LIWASC_OIDC_REDIRECT_URL"),
			HomePath:    "/",
			Scopes:      []string{"profile", "email"},

			LocalStoragePrefix: "liwasc",

			Children: func(loginProviderChildrenProps components.OIDCLoginProviderChildrenProps) app.UI {
				if loginProviderChildrenProps.OAuth2Token.AccessToken == "" || loginProviderChildrenProps.UserInfo.Email == "" {
					return &components.LoadingPageComponent{
						Message: "Logging you in ...",
					}
				}

				if loginProviderChildrenProps.Error != nil {
					return &components.FatalErrorPageComponent{
						Header:             "Oh no! A fatal error occured.",
						Description:        "The following message might be of help; more details might be in the console:",
						StackTraceLanguage: "Go Stacktrace",
						StackTraceContent:  loginProviderChildrenProps.Error.Error(),
						Actions: []app.UI{
							app.Button().Class("pf-c-button pf-m-primary").Body(
								app.Span().Class("pf-c-button__icon pf-m-start").Body(
									app.I().Class("fas fa-sync-alt"),
								),
								app.Text("Restart liwasc"),
							).OnClick(func(ctx app.Context, e app.Event) { app.Reload() }),
						},
					}
				}

				return &experimental.DataProviderComponent{
					AuthenticatedContext:   metadata.AppendToOutgoingContext(context.Background(), components.AUTHORIZATION_METADATA_KEY, loginProviderChildrenProps.IDToken),
					MetadataService:        metadataService,
					NodeAndPortScanService: nodeAndPortScanService,
					Children: func(dpcp experimental.DataProviderChildrenProps) app.UI {
						return app.Div().Body(
							&experimental.ActionsComponent{
								Nodes: dpcp.Network.Nodes,

								TriggerNetworkScan: dpcp.TriggerNetworkScan,
							},
							&experimental.JSONOutputComponent{
								Object: dpcp.Network,
							},
						)
					},
				}

				// return &components.DataProviderComponent{
				// 	IDToken: loginProviderChildrenProps.IDToken,

				// 	MetadataServiceClient:        metadataServiceClient,
				// 	NodeAndPortScanServiceClient: nodeAndPortScanServiceClient,
				// 	Children: func(dataProviderChildrenProps components.DataProviderChildrenProps) app.UI {
				// 		return &components.AppComponent{
				// 			UserAvatar: fmt.Sprintf("https://www.gravatar.com/avatar/%x", md5.Sum([]byte(loginProviderChildrenProps.UserInfo.Email))),
				// 			UserName:   loginProviderChildrenProps.UserInfo.Email,

				// 			Logout: loginProviderChildrenProps.Logout,

				// 			Subnets:         dataProviderChildrenProps.Subnets,
				// 			Device:          dataProviderChildrenProps.Device,
				// 			NodeSearchValue: "",

				// 			Nodes: dataProviderChildrenProps.Nodes,

				// 			InspectorSearchValue: "",

				// 			Connected: dataProviderChildrenProps.Connected,
				// 			Scanning:  dataProviderChildrenProps.Scanning,

				// 			TriggerNodeScan: func() {
				// 				protoNodeScanTriggerMessage := &proto.NodeScanStartMessage{
				// 					NodeScanTimeout: 500,
				// 					PortScanTimeout: 50,
				// 					MACAddress:      "",
				// 				}

				// 				go dataProviderChildrenProps.TriggerNodeScan(protoNodeScanTriggerMessage)
				// 			},
				// 		}
				// 	},
				// }
			},
		},
	)

	app.Run()
}
