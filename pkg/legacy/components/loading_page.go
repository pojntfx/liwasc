package components

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type LoadingPageComponent struct {
	app.Compo

	Message string
}

func (c *LoadingPageComponent) Render() app.UI {
	return app.Div().Class("x__page--full").Body(
		app.Span().Class("pf-c-spinner pf-u-mb-lg").Body(
			app.Span().Class("pf-c-spinner__clipper"),
			app.Span().Class("pf-c-spinner__lead-ball"),
			app.Span().Class("pf-c-spinner__tail-ball"),
		),
		app.H1().Text(c.Message),
	)
}
