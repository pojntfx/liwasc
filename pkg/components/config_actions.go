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
			&FormGroupComponent{
				Label: app.
					Label().
					For(oidcIssuerName).
					Class("pf-c-form__label").
					Body(
						app.
							Span().
							Class("pf-c-form__label-text").
							Text("OIDC Issuer"),
					),
				Input: &Controlled{
					Component: app.
						Input().
						Name(oidcIssuerName).
						ID(oidcIssuerName).
						Type("url").
						Required(true).
						Placeholder(oidcIssuerPlaceholder).
						Class("pf-c-form-control").
						OnInput(func(ctx app.Context, e app.Event) {
							c.SetOIDCIssuer(ctx.JSSrc.Get("value").String())
						}),
					Value: c.OIDCIssuer,
				},
				Required: true,
			},
			// OIDC Client ID
			&FormGroupComponent{
				Label: app.
					Label().
					For(oidcClientIDName).
					Class("pf-c-form__label").
					Body(
						app.
							Span().
							Class("pf-c-form__label-text").
							Text("OIDC Client ID"),
					),
				Input: &Controlled{
					Component: app.
						Input().
						Name(oidcClientIDName).
						ID(oidcClientIDName).
						Type("text").
						Required(true).
						Class("pf-c-form-control").
						OnInput(func(ctx app.Context, e app.Event) {
							c.SetOIDCClientID(ctx.JSSrc.Get("value").String())
						}),
					Value: c.OIDCClientID,
				},
				Required: true,
			},
			// OIDC Redirect URL
			&FormGroupComponent{
				Label: app.
					Label().
					For(oidcRedirectURLName).
					Class("pf-c-form__label").
					Body(
						app.
							Span().
							Class("pf-c-form__label-text").
							Text("OIDC Redirect URL"),
					),
				Input: &Controlled{
					Component: app.
						Input().
						Name(oidcRedirectURLName).
						ID(oidcRedirectURLName).
						Type("url").
						Required(true).
						Placeholder(oidcRedirectURLPlaceholder).
						Class("pf-c-form-control").
						OnInput(func(ctx app.Context, e app.Event) {
							c.SetOIDCRedirectURL(ctx.JSSrc.Get("value").String())
						}),
					Value: c.OIDCRedirectURL,
				},
				Required: true,
			},
			// Configuration Apply Trigger
			app.Div().
				Class("pf-c-form__group pf-m-action").
				Body(
					app.Div().
						Class("pf-c-form__actions").
						Body(
							app.
								Button().
								Type("submit").
								Class("pf-c-button pf-m-primary").
								Text("Apply Configuration"),
						),
				),
		).OnSubmit(func(ctx app.Context, e app.Event) {
		e.PreventDefault()

		go c.ApplyConfig()
	})
}
