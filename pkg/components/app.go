package components

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type AppComponent struct {
	app.Compo
}

func NewAppComponent() *AppComponent {
	return &AppComponent{}
}

func (c *AppComponent) Render() app.UI {
	return app.Button().Class("pf-c-button pf-m-primary").Body(
		app.Span().Class("pf-c-button__icon pf-m-start").Body(
			app.I().Class("fas fa-plus-circle"),
		),
		app.Text("Re-Scan"),
	)
}
