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

	conn, err := grpc.Dial("ws://stuttgart.felicitas.pojtinger.com:15124", grpc.WithContextDialer(proxy.Dialer), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	networkAndNodeScanServiceClient := proto.NewNetworkAndNodeScanServiceClient(conn)
	nodeWakeServiceClient := proto.NewNodeWakeServiceClient(conn)
	metadataServiceClient := proto.NewMetadataServiceClient(conn)

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

				return &components.DataProviderComponent{
					IDToken: loginProviderChildrenProps.IDToken,

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

							TriggerNetworkScan: func() {
								protoNetworkScanTriggerMessage := &proto.NetworkScanTriggerMessage{
									NetworkScanTimeout: 1000,
									NodeScanTimeout:    500,
								}

								dataProviderChildrenProps.TriggerNetworkScan(protoNetworkScanTriggerMessage)
							},
						}
					},
				}
			},
		},
	)

	app.Run()
}
