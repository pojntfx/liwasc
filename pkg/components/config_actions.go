package components

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
	return app.Form().
		Class("pf-c-form pf-m-horizontal").
		Body(
			// Backend URL Input
			&FormGroupComponent{
				Label: app.
					Label().
					For(backendURLName).
					Class("pf-c-form__label").
					Body(
						app.
							Span().
							Class("pf-c-form__label-text").
							Text("Backend URL"),
					),
				Input: &Controlled{
					Component: app.
						Input().
						Name(backendURLName).
						ID(backendURLName).
						Type("url").
						Required(true).
						Placeholder(backendURLPlaceholder).
						Class("pf-c-form-control").
						OnInput(func(ctx app.Context, e app.Event) {
							c.SetBackendURL(ctx.JSSrc.Get("value").String())
						}),
					Value: c.BackendURL,
				},
				Required: true,
			},
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
			// Configuration Apply Trigger
			app.
				Input().
				Type("submit").
				Value("Apply Configuration").
				Class("pf-c-button pf-m-primary"),
		).OnSubmit(func(ctx app.Context, e app.Event) {
		e.PreventDefault()

		go c.ApplyConfig()
	})
}
