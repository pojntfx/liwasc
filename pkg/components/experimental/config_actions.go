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
				Text("Backend URL: "),
			&Controlled{
				Component: app.
					Input().
					Name(backendURLName).
					ID(backendURLName).
					Type("url").
					Required(true).
					Placeholder(backendURLPlaceholder).
					OnInput(func(ctx app.Context, e app.Event) {
						c.SetBackendURL(ctx.JSSrc.Get("value").String())
					}),
				Value: c.BackendURL,
			},
			app.Br(),
			// OIDC Issuer Input
			app.
				Label().
				For(oidcIssuerName).
				Text("OIDC Issuer: "),
			&Controlled{
				Component: app.
					Input().
					Name(oidcIssuerName).
					ID(oidcIssuerName).
					Type("url").
					Required(true).
					Placeholder(oidcIssuerPlaceholder).
					OnInput(func(ctx app.Context, e app.Event) {
						c.SetOIDCIssuer(ctx.JSSrc.Get("value").String())
					}),
				Value: c.OIDCIssuer,
			},
			app.Br(),
			// OIDC Client ID
			app.
				Label().
				For(oidcClientIDName).
				Text("OIDC Client ID: "),
			&Controlled{
				Component: app.
					Input().
					Name(oidcClientIDName).
					ID(oidcClientIDName).
					Type("text").
					Required(true).
					OnInput(func(ctx app.Context, e app.Event) {
						c.SetOIDCClientID(ctx.JSSrc.Get("value").String())
					}),
				Value: c.OIDCClientID,
			},
			app.Br(),
			// OIDC Redirect URL
			app.
				Label().
				For(oidcRedirectURLName).
				Text("OIDC Redirect URL: "),
			&Controlled{
				Component: app.
					Input().
					Name(oidcRedirectURLName).
					ID(oidcRedirectURLName).
					Type("url").
					Required(true).
					Placeholder(oidcRedirectURLPlaceholder).
					OnInput(func(ctx app.Context, e app.Event) {
						c.SetOIDCRedirectURL(ctx.JSSrc.Get("value").String())
					}),
				Value: c.OIDCRedirectURL,
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
