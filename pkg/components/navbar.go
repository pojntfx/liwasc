package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type NavbarComponent struct {
	app.Compo

	UserMenuOpen bool
	UserAvatar   string
	UserName     string

	Connected bool

	OnUserMenuToggle func(ctx app.Context, e app.Event)
	OnSignOutClick   func(ctx app.Context, e app.Event)
}

func (c *NavbarComponent) Render() app.UI {
	return app.Header().Class("pf-c-page__header").Body(
		app.Div().Class("pf-c-page__header-brand").Body(
			app.A().Href("#").Class("pf-c-page__header-brand-link").Body(
				app.Img().Class("pf-c-brand pf-u-py-sm x__brand--svg").Src("/web/logo.svg").Alt("liwasc logo"),
			),
		),
		app.Div().Class("pf-c-page__header-tools").Body(
			app.Div().Class("pf-c-page__header-tools-group").Body(
				app.Div().Class("pf-c-page__header-tools-item ").Body(
					&TooltipComponent{
						Children: app.I().Class(fmt.Sprintf(
							"fas %v %v",
							func() string {
								if c.Connected {
									return "fa-satellite-dish"
								}

								return "fa-exclamation-triangle"
							}(),
							func() string {
								if c.Connected {
									return "x__icon--pulsating"
								}

								return "x__icon--danger"
							}())),
						Tooltip: app.Div().Body(
							app.If(
								c.Connected,
								app.I().Class("fas fa-check-circle pf-u-mr-sm"),
							).Else(
								app.I().Class("fas fa-times pf-u-mr-sm"),
							),
							app.If(
								c.Connected,
								app.Text("You are connected to the real-time node stream. No need to reload!"),
							).Else(
								app.Text("You are not connected to the real-time node stream. Please check the console for errors and reload!"),
							),
						),
					},
				),
			),
			app.Div().Class(fmt.Sprintf("pf-c-dropdown %v", func() string {
				if c.UserMenuOpen {
					return "pf-m-expanded"
				}

				return ""
			}())).Body(
				app.Button().Class("pf-c-dropdown__toggle pf-m-plain").Body(
					app.Span().Class("pf-c-dropdown__toggle-image pf-u-mr-0 pf-u-mr-sm-on-md").Body(
						app.Img().Class("pf-c-avatar").Src(c.UserAvatar).Alt("User avatar"),
					),
					app.Span().Class("pf-c-dropdown__toggle-text pf-u-display-none pf-u-display-flex-on-md").Text(c.UserName),
					app.Span().Class("pf-c-dropdown__toggle-icon").Body(
						app.I().Class("fas fa-caret-down"),
					),
				).OnClick(c.OnUserMenuToggle),
				app.Ul().Class("pf-c-dropdown__menu pf-m-align-right").Hidden(!c.UserMenuOpen).Body(
					app.Li().Body(
						app.Button().Class("pf-c-dropdown__menu-item pf-m-icon").Body(
							app.Span().Class("pf-c-dropdown__menu-item-icon").Body(
								app.I().Class("fas fa-sign-out-alt"),
							),
							app.Span().Text("Sign Out"),
						).OnClick(c.OnSignOutClick),
					),
				),
			),
		),
	)
}
