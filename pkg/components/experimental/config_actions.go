package experimental

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type ConfigActionsComponent struct {
	app.Compo

	BackendURL      string
	OIDCIssuer      string
	OIDCClientID    string
	OIDCRedirectURL string

	SetBackendURL,
	SetOIDCIssuer,
	SetOIDCClientID,
	SetOIDCRedirectURL func(string)
	ApplyConfig func()
}

const (
	// Names and IDs
	backendURLName      = "backendURLName"
	oidcIssuerName      = "oidcIssuer"
	oidcClientIDName    = "oidcClientID"
	oidcRedirectURLName = "oidcRedirectURL"

	// Placeholders
	backendURLPlaceholder      = "ws://localhost:15124"
	oidcIssuerPlaceholder      = "https://pojntfx.eu.auth0.com/"
	oidcRedirectURLPlaceholder = "http://localhost:15125/"
)

func (c *ConfigActionsComponent) Render() app.UI {
	return app.Div().Body(
		app.Form().Body(
			// Backend URL Input
			app.
				Label().
				For(backendURLName).
				Text("Node Scan Timeout (in ms): "),
			&Controlled{
				Component: app.
					Input().
					Name(backendURLName).
					Type("url").
					Required(true).
					Placeholder(backendURLPlaceholder).
					OnInput(func(ctx app.Context, e app.Event) {
						c.SetBackendURL(ctx.JSSrc.Get("value").String())
					}),
				Value: c.BackendURL,
			},
			app.Br(),
			// Configuration Apply Trigger
			app.
				Input().
				Type("submit").
				Value("Apply Configuration"),
		).OnSubmit(func(ctx app.Context, e app.Event) {
			e.PreventDefault()

			go c.ApplyConfig()
		}),
	)
}
