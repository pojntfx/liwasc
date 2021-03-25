// +build js

package main

import (
	"context"
	"time"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/go-app-grpc-chat-frontend-web/pkg/websocketproxy"
	proto "github.com/pojntfx/liwasc/pkg/proto/generated"
	"github.com/pojntfx/liwasc/pkg/providers"
	"github.com/pojntfx/liwasc/pkg/shells"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	// Define the routes
	app.Route("/",
		// Configuration provider
		&providers.ConfigurationProvider{
			StoragePrefix:       "liwasc.configuration",
			StateQueryParameter: "state",
			CodeQueryParameter:  "code",
			Children: func(cpcp providers.ConfigurationProviderChildrenProps) app.UI {
				// This div is required so that there are no authorization loops
				return app.Div().Body(
					app.If(cpcp.Ready,
						// Identity provider
						&providers.IdentityProvider{
							Issuer:        cpcp.OIDCIssuer,
							ClientID:      cpcp.OIDCClientID,
							RedirectURL:   cpcp.OIDCRedirectURL,
							HomeURL:       "/",
							Scopes:        []string{"profile", "email"},
							StoragePrefix: "liwasc.identity",
							Children: func(ipcp providers.IdentityProviderChildrenProps) app.UI {
								// Configuration shell
								if ipcp.Error != nil {
									return &shells.ConfigurationShell{
										LogoSrc:          "/web/logo.svg",
										Title:            "Log in to liwasc",
										ShortDescription: "List, wake and scan nodes in a network.",
										LongDescription: `liwasc is a high-performance network and port scanner. It can
quickly give you a overview of the nodes in your network, the
services that run on them and manage their power status.`,
										HelpLink: "https://github.com/pojntfx/liwasc/wiki",
										Links: map[string]string{
											"License":       "https://github.com/pojntfx/liwasc/blob/main/LICENSE",
											"Source Code":   "https://github.com/pojntfx/liwasc",
											"Documentation": "https://github.com/pojntfx/liwasc/wiki",
										},

										BackendURL:      cpcp.BackendURL,
										OIDCIssuer:      cpcp.OIDCIssuer,
										OIDCClientID:    cpcp.OIDCClientID,
										OIDCRedirectURL: cpcp.OIDCRedirectURL,

										SetBackendURL:      cpcp.SetBackendURL,
										SetOIDCIssuer:      cpcp.SetOIDCIssuer,
										SetOIDCClientID:    cpcp.SetOIDCClientID,
										SetOIDCRedirectURL: cpcp.SetOIDCRedirectURL,
										ApplyConfig:        cpcp.ApplyConfig,

										Error: ipcp.Error,
									}
								}

								// Configuration placeholder
								if ipcp.IDToken == "" || ipcp.UserInfo.Email == "" {
									return app.P().Text("Authorizing ...")
								}

								// gRPC Client
								conn, err := grpc.Dial(cpcp.BackendURL, grpc.WithContextDialer(websocketproxy.NewWebSocketProxyClient(time.Minute).Dialer), grpc.WithInsecure())
								if err != nil {
									panic(err)
								}

								// Data provider
								return &providers.DataProvider{
									AuthenticatedContext:   metadata.AppendToOutgoingContext(context.Background(), "X-Liwasc-Authorization", ipcp.IDToken),
									MetadataService:        proto.NewMetadataServiceClient(conn),
									NodeAndPortScanService: proto.NewNodeAndPortScanServiceClient(conn),
									NodeWakeService:        proto.NewNodeWakeServiceClient(conn),
									Children: func(dpcp providers.DataProviderChildrenProps) app.UI {
										// Data shell
										return &shells.DataShell{
											Network:  dpcp.Network,
											UserInfo: ipcp.UserInfo,

											TriggerNetworkScan: dpcp.TriggerNetworkScan,
											StartNodeWake:      dpcp.StartNodeWake,
											Logout:             ipcp.Logout,

											Error:   dpcp.Error,
											Recover: dpcp.Recover,
										}
									},
								}
							},
						},
					).Else(
						// Configuration shell
						&shells.ConfigurationShell{
							LogoSrc:          "/web/logo.svg",
							Title:            "Log in to liwasc",
							ShortDescription: "List, wake and scan nodes in a network.",
							LongDescription: `liwasc is a high-performance network and port scanner. It can
quickly give you a overview of the nodes in your network, the
services that run on them and manage their power status.`,
							HelpLink: "https://github.com/pojntfx/liwasc/wiki",
							Links: map[string]string{
								"License":       "https://github.com/pojntfx/liwasc/blob/main/LICENSE",
								"Source Code":   "https://github.com/pojntfx/liwasc",
								"Documentation": "https://github.com/pojntfx/liwasc/wiki",
							},

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
