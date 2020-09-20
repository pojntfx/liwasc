package components

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type NavbarComponent struct {
	app.Compo

	CurrentUserEmail string
}

func (c *NavbarComponent) Render() app.UI {
	return app.Header().Class("pf-c-page__header").Body(
		app.Div().Class("pf-c-page__header-brand").Body(
			app.A().Href("#").Class("pf-c-page__header-brand-link").Body(
				app.Img().Class("pf-c-brand").Src("https://www.patternfly.org/assets/images/PF-Masthead-Logo.svg").Alt("liwasc logo"),
			),
		),
		app.Div().Class("pf-c-page__header-tools").Body(
			app.Img().Class("pf-c-avatar").Src(fmt.Sprintf("https://www.gravatar.com/avatar/%v", func() string {
				hasher := md5.New()

				hasher.Write([]byte(c.CurrentUserEmail))

				return hex.EncodeToString(hasher.Sum(nil))
			}())).Alt("User avatar"),
		),
	)
}
