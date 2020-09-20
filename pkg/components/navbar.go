package components

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type NavbarComponent struct {
	app.Compo

	CurrentUserEmail       string
	CurrentUserDisplayName string
	UserMenuOpen           bool
	OnUserMenuToggle       func(ctx app.Context, e app.Event)
}

func (c *NavbarComponent) Render() app.UI {
	return app.Header().Class("pf-c-page__header").Body(
		app.Div().Class("pf-c-page__header-brand").Body(
			app.A().Href("#").Class("pf-c-page__header-brand-link").Body(
				app.Img().Class("pf-c-brand").Src("https://www.patternfly.org/assets/images/PF-Masthead-Logo.svg").Alt("liwasc logo"),
			),
		),
		app.Div().Class("pf-c-page__header-tools").Body(
			app.Div().Class(fmt.Sprintf("pf-c-dropdown %v", func() string {
				if c.UserMenuOpen {
					return "pf-m-expanded"
				}

				return ""
			}())).Body(
				app.Button().Class("pf-c-dropdown__toggle pf-m-plain").Body(
					app.Span().Class("pf-c-dropdown__toggle-image pf-u-mr-0 pf-u-mr-sm-on-md").Body(
						app.Img().Class("pf-c-avatar").Src(fmt.Sprintf("https://www.gravatar.com/avatar/%v", func() string {
							hasher := md5.New()

							hasher.Write([]byte(c.CurrentUserEmail))

							return hex.EncodeToString(hasher.Sum(nil))
						}())).Alt("User avatar"),
					),
					app.Span().Class("pf-c-dropdown__toggle-text pf-u-display-none pf-u-display-flex-on-md").Text(c.CurrentUserDisplayName),
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
						),
					),
				),
			),
		),
	)
}
