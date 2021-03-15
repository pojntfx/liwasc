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
				return app.Div().Body(
					// Config status
					&components.StatusComponent{
						Error: cpcp.Error,
					},
					// Config actions
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
					},
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
								// Login placeholder
								if lpcp.IDToken == "" || lpcp.UserInfo.Email == "" {
									return app.P().Text("Authorizing ...")
								}

								// gRPC Client
								conn, err := grpc.Dial(cpcp.BackendURL, grpc.WithContextDialer(websocketproxy.NewWebSocketProxyClient(time.Minute).Dialer), grpc.WithInsecure())
								if err != nil {
									panic(err)
								}

								return app.Div().Body(
									// Login status
									&components.StatusComponent{
										Error:   lpcp.Error,
										Recover: lpcp.Recover,
									},
									// Login actions
									&components.LoginActionsComponent{
										Logout: lpcp.Logout,
									},
									// Login output
									&components.JSONOutputComponent{
										Object: struct {
											Email string
										}{
											Email: lpcp.UserInfo.Email,
										},
									},
									// Data provider
									&components.DataProviderComponent{
										AuthenticatedContext:   metadata.AppendToOutgoingContext(context.Background(), "X-Liwasc-Authorization", lpcp.IDToken),
										MetadataService:        proto.NewMetadataServiceClient(conn),
										NodeAndPortScanService: proto.NewNodeAndPortScanServiceClient(conn),
										NodeWakeService:        proto.NewNodeWakeServiceClient(conn),
										Children: func(dpcp components.DataProviderChildrenProps) app.UI {
											return app.Div().Body(
												// Data status
												&components.StatusComponent{
													Error:   dpcp.Error,
													Recover: dpcp.Recover,
												},
												// Data actions
												&components.DataActionsComponent{
													Network: dpcp.Network,

													TriggerNetworkScan: dpcp.TriggerNetworkScan,
													StartNodeWake:      dpcp.StartNodeWake,
												},
												// Data output
												&components.JSONOutputComponent{
													Object: dpcp.Network,
												},
											)
										},
									},
								)
							},
						},
					),
				)
			},
		},
	)

	// Start the app
	app.Run()
}
