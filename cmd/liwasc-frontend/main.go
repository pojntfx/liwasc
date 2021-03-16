// +build js,wasm

package main

import (
	"context"
	"time"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/go-app-grpc-chat-frontend-web/pkg/websocketproxy"
	components "github.com/pojntfx/liwasc/pkg/components"
	proto "github.com/pojntfx/liwasc/pkg/proto/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	// Define the routes
	app.Route("/",
		// Config provider
		&components.ConfigProviderComponent{
			StoragePrefix: "liwasc.config",
			Children: func(cpcp components.ConfigProviderChildrenProps) app.UI {
				// This div is required so that there are no authorization loops
				return app.Div().Body(
					app.If(cpcp.Ready,
						// Login provider
						&components.LoginProviderComponent{
							Issuer:        cpcp.OIDCIssuer,
							ClientID:      cpcp.OIDCClientID,
							RedirectURL:   cpcp.OIDCRedirectURL,
							HomeURL:       "/",
							Scopes:        []string{"profile", "email"},
							StoragePrefix: "liwasc.login",
							Children: func(lpcp components.LoginProviderChildrenProps) app.UI {
								// Login actions and status
								if lpcp.Error != nil {
									return &components.ConfigActionsComponent{
										BackendURL:      cpcp.BackendURL,
										OIDCIssuer:      cpcp.OIDCIssuer,
										OIDCClientID:    cpcp.OIDCClientID,
										OIDCRedirectURL: cpcp.OIDCRedirectURL,

										SetBackendURL:      cpcp.SetBackendURL,
										SetOIDCIssuer:      cpcp.SetOIDCIssuer,
										SetOIDCClientID:    cpcp.SetOIDCClientID,
										SetOIDCRedirectURL: cpcp.SetOIDCRedirectURL,
										ApplyConfig:        cpcp.ApplyConfig,

										Error: lpcp.Error,
									}
								}

								// Login placeholder
								if lpcp.IDToken == "" || lpcp.UserInfo.Email == "" {
									return app.P().Text("Authorizing ...")
								}

								// gRPC Client
								conn, err := grpc.Dial(cpcp.BackendURL, grpc.WithContextDialer(websocketproxy.NewWebSocketProxyClient(time.Minute).Dialer), grpc.WithInsecure())
								if err != nil {
									panic(err)
								}

								// Data provider
								return &components.DataProviderComponent{
									AuthenticatedContext:   metadata.AppendToOutgoingContext(context.Background(), "X-Liwasc-Authorization", lpcp.IDToken),
									MetadataService:        proto.NewMetadataServiceClient(conn),
									NodeAndPortScanService: proto.NewNodeAndPortScanServiceClient(conn),
									NodeWakeService:        proto.NewNodeWakeServiceClient(conn),
									Children: func(dpcp components.DataProviderChildrenProps) app.UI {
										return &components.DataActionsComponent{
											Network:  dpcp.Network,
											UserInfo: lpcp.UserInfo,

											TriggerNetworkScan: dpcp.TriggerNetworkScan,
											StartNodeWake:      dpcp.StartNodeWake,
											Logout:             lpcp.Logout,

											Error:   dpcp.Error,
											Recover: dpcp.Recover,
										}
									},
								}
							},
						},
					).Else(
						// Config actions and status
						&components.ConfigActionsComponent{
							BackendURL:      cpcp.BackendURL,
							OIDCIssuer:      cpcp.OIDCIssuer,
							OIDCClientID:    cpcp.OIDCClientID,
							OIDCRedirectURL: cpcp.OIDCRedirectURL,

							SetBackendURL:      cpcp.SetBackendURL,
							SetOIDCIssuer:      cpcp.SetOIDCIssuer,
							SetOIDCClientID:    cpcp.SetOIDCClientID,
							SetOIDCRedirectURL: cpcp.SetOIDCRedirectURL,
							ApplyConfig:        cpcp.ApplyConfig,

							Error: cpcp.Error,
						},
					),
				)
			},
		},
	)

	// Start the app
	app.Run()
}
