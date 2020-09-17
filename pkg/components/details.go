package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type DetailsComponent struct {
	app.Compo
	Open    bool
	OnClose func(ctx app.Context, e app.Event)
	Title   string
	Main    app.UI
	Details app.UI
}

func (c *DetailsComponent) Render() app.UI {
	return app.Div().Class(fmt.Sprintf("pf-c-drawer %v", func() string {
		if c.Open {
			return "pf-m-expanded"
		}

		return ""
	}())).Body(app.Div().Class("pf-c-drawer__main").Body(
		app.Div().Class("pf-c-drawer__content").Body(
			app.Div().Class("pf-c-drawer__body").Body(c.Main),
		),
		app.Div().Class("pf-c-drawer__panel").Body(
			app.Div().Class("pf-c-drawer__body").Body(
				app.Div().Class("pf-c-drawer__head").Body(
					app.Span().Body(app.Text(c.Title)),
					app.Div().Class("pf-c-drawer__actions").Body(
						app.Div().Class("pf-c-drawer__close").Body(
							app.Button().Class("pf-c-button pf-m-plain").Body(
								app.I().Class("fas fa-times"),
							).OnClick(c.OnClose),
						),
					),
				),
			),
		),
	))
}
