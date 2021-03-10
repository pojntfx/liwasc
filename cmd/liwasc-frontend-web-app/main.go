package main

import (
	"context"
	"time"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/go-app-grpc-chat-frontend-web/pkg/websocketproxy"
	"github.com/pojntfx/liwasc-frontend-web/pkg/components/experimental"
	proto "github.com/pojntfx/liwasc-frontend-web/pkg/proto/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	// Define the routes
	app.Route("/",
		// Config provider
		&experimental.ConfigProviderComponent{
			Children: func(cpcp experimental.ConfigProviderChildrenProps) app.UI {
				return app.Div().Body(
					// Config status
					&experimental.StatusComponent{
						Error: cpcp.Error,
					},
					// Config actions
					&experimental.ConfigActionsComponent{
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
						&experimental.LoginProviderComponent{
							Issuer:        cpcp.OIDCIssuer,
							ClientID:      cpcp.OIDCClientID,
							RedirectURL:   cpcp.OIDCRedirectURL,
							HomeURL:       "/",
							Scopes:        []string{"profile", "email"},
							StoragePrefix: "liwasc",
							Children: func(lpcp experimental.LoginProviderChildrenProps) app.UI {
								// Login placeholder
								if lpcp.IDToken == "" || lpcp.UserInfo.Email == "" {
									return app.P().Text("Authorizing ...")
								}

								// Connect to the backend
								conn, err := grpc.Dial(app.Getenv("LIWASC_BACKEND_URL"), grpc.WithContextDialer(websocketproxy.NewWebSocketProxyClient(time.Minute).Dialer), grpc.WithInsecure())
								if err != nil {
									panic(err)
								}
								defer conn.Close()

								return app.Div().Body(
									// Login status
									&experimental.StatusComponent{
										Error:   lpcp.Error,
										Recover: lpcp.Recover,
									},
									// Data provider
									&experimental.DataProviderComponent{
										AuthenticatedContext:   metadata.AppendToOutgoingContext(context.Background(), "X-Liwasc-Authorization", lpcp.IDToken),
										MetadataService:        proto.NewMetadataServiceClient(conn),
										NodeAndPortScanService: proto.NewNodeAndPortScanServiceClient(conn),
										NodeWakeService:        proto.NewNodeWakeServiceClient(conn),
										Children: func(dpcp experimental.DataProviderChildrenProps) app.UI {
											return app.Div().Body(
												// Data actions
												&experimental.DataActionsComponent{
													Nodes: dpcp.Network.Nodes,

													TriggerNetworkScan: dpcp.TriggerNetworkScan,
													StartNodeWake:      dpcp.StartNodeWake,
												},
												// Data status
												&experimental.StatusComponent{
													Error:   dpcp.Error,
													Recover: dpcp.Recover,
												},
												// Data output
												&experimental.JSONOutputComponent{
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
