package components

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type TooltipComponent struct {
	app.Compo

	Children app.UI
	Tooltip  app.UI
}

func (c *TooltipComponent) Render() app.UI {
	return app.Div().Class("pf-c-tooltip pf-m-bottom x__tooltip").Body(
		c.Children,
		app.Div().Class("x__tooltip__wrapper").Body(
			app.Div().Class("pf-c-tooltip__arrow x__tooltip__wrapper__arrow"),
			app.Div().Class("pf-c-tooltip__content x__tooltip__wrapper__content").Body(
				c.Tooltip,
			),
		),
	)
}
