package components

import (
	"context"
	"time"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/pojntfx/go-app-grpc-chat-frontend-web/pkg/websocketproxy"
	api "github.com/pojntfx/liwasc/pkg/api/proto/v1"
	"github.com/pojntfx/liwasc/pkg/providers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Home struct {
	app.Compo
}

func (c *Home) Render() app.UI {
	return &providers.ConfigurationProvider{
		StoragePrefix:       "liwasc.configuration",
		StateQueryParameter: "state",
		CodeQueryParameter:  "code",
		Children: func(cpcp providers.SetupProviderChildrenProps) app.UI {
			// This div is required so that there are no authorization loops
			return app.Div().
				TabIndex(-1).
				Class("pf-x-ws-router").
				Body(
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
									return &SetupShell{
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
									MetadataService:        api.NewMetadataServiceClient(conn),
									NodeAndPortScanService: api.NewNodeAndPortScanServiceClient(conn),
									NodeWakeService:        api.NewNodeWakeServiceClient(conn),
									Children: func(dpcp providers.DataProviderChildrenProps) app.UI {
										// Data shell
										return &DataShell{
											Network:  dpcp.Network,
											UserInfo: ipcp.UserInfo,

											TriggerNetworkScan: dpcp.TriggerNetworkScan,
											StartNodeWake:      dpcp.StartNodeWake,
											Logout:             ipcp.Logout,

											Error:   dpcp.Error,
											Recover: dpcp.Recover,
											Ignore:  dpcp.Ignore,
										}
									},
								}
							},
						},
					).Else(
						// Configuration shell
						&SetupShell{
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
	}
}
