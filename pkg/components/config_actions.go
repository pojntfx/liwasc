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

	Error error
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
	// Display the error message if error != nil
	errorMessage := ""
	if c.Error != nil {
		errorMessage = c.Error.Error()
	}

	return app.Div().
		Body(
			app.Div().
				Class("pf-c-background-image").
				Body(
					app.Raw(`<svg
  xmlns="http://www.w3.org/2000/svg"
  class="pf-c-background-image__filter"
  width="0"
  height="0"
>
  <filter id="image_overlay">
    <feColorMatrix
      type="matrix"
      values="1 0 0 0 0 1 0 0 0 0 1 0 0 0 0 0 0 0 1 0"
    ></feColorMatrix>
    <feComponentTransfer
      color-interpolation-filters="sRGB"
      result="duotone"
    >
      <feFuncR
        type="table"
        tableValues="0.086274509803922 0.43921568627451"
      ></feFuncR>
      <feFuncG
        type="table"
        tableValues="0.086274509803922 0.43921568627451"
      ></feFuncG>
      <feFuncB
        type="table"
        tableValues="0.086274509803922 0.43921568627451"
      ></feFuncB>
      <feFuncA type="table" tableValues="0 1"></feFuncA>
    </feComponentTransfer>
  </filter>
</svg>`),
				),
			app.Div().Class("pf-c-login").Body(
				app.Div().Class("pf-c-login__container").Body(
					app.Header().Class("pf-c-login__header").Body(
						app.Img().
							Class("pf-c-brand pc-x-c-brand--main").
							Src("/web/logo.svg").
							Alt("liwasc Logo"),
					),
					app.Main().Class("pf-c-login__main").Body(
						app.Header().Class("pf-c-login__main-header").Body(
							app.H1().Class("pf-c-title pf-m-3xl").Text(
								"Log in to liwasc",
							),
							app.P().Class("pf-c-login__main-header-desc").Text(
								"List, wake and scan nodes in a network.",
							),
						),
						app.Div().Class("pf-c-login__main-body").Body(
							app.Form().
								Class("pf-c-form").
								Body(
									// Error display
									app.If(c.Error != nil, app.P().
										Class("pf-c-form__helper-text pf-m-error").
										Aria("live", "polite").
										Body(
											app.Span().
												Class("pf-c-form__helper-text-icon").
												Body(
													app.I().
														Class("fas fa-exclamation-circle").
														Aria("hidden", true),
												),
											app.Text(errorMessage),
										),
									),
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
												Aria("invalid", c.Error != nil).
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
												Aria("invalid", c.Error != nil).
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
												Aria("invalid", c.Error != nil).
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
												Aria("invalid", c.Error != nil).
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
											app.
												Button().
												Type("submit").
												Class("pf-c-button pf-m-primary pf-m-block").
												Text("Log in"),
										),
								).OnSubmit(func(ctx app.Context, e app.Event) {
								e.PreventDefault()

								go c.ApplyConfig()
							}),
						),
						app.Footer().Class("pf-c-login__main-footer").Body(
							app.Div().Class("pf-c-login__main-footer-band").Body(
								app.P().Class("pf-c-login__main-footer-band-item").Body(
									app.Text("Not sure what to do? "),
									app.A().
										Href("https://github.com/pojntfx/liwasc/wiki").
										Target("_blank").
										Text("Get help."),
								),
							),
						),
					),
					app.Footer().Class("pf-c-login__footer").Body(
						app.P().Text(`liwasc is a high-performance network and port scanner. It can
quickly give you a overview of the nodes in your network, the
services that run on them and manage their power status.`),
						app.Ul().Class("pf-c-list pf-m-inline").Body(
							app.Li().Body(
								app.
									A().
									Target("_blank").
									Href("https://github.com/pojntfx/liwasc/blob/main/LICENSE").
									Text("License"),
							),
							app.Li().Body(
								app.
									A().
									Target("_blank").
									Href("https://github.com/pojntfx/liwasc").
									Text("Source Code"),
							),
							app.Li().Body(
								app.
									A().
									Target("_blank").
									Href("https://github.com/pojntfx/liwasc/wiki").
									Text("Documentation"),
							),
						),
					),
				),
			),
		)

}
