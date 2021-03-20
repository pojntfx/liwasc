package components

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type Status struct {
	app.Compo

	Error   error
	Recover func()
}

func (c *Status) Render() app.UI {
	// Display the error message if error != nil
	errorMessage := ""
	if c.Error != nil {
		errorMessage = c.Error.Error()
	}

	return app.If(c.Error != nil, app.Div().
		Class("pf-c-alert pf-m-danger").
		Aria("label", "Error").
		Body(
			app.Div().
				Class("pf-c-alert__icon").
				Body(
					app.I().
						Class("fas fa-fw fa-exclamation-circle").
						Aria("hidden", true),
				),
			app.P().
				Class("pf-c-alert__title").
				Body(
					app.Strong().Body(
						app.Span().
							Class("pf-screen-reader").
							Text("Error"),
					),
					app.Text("Error"),
				),
			app.Div().
				Class("pf-c-alert__description").
				Body(
					app.P().Body(
						app.Code().
							Text(errorMessage),
					),
				),
			app.If(c.Recover != nil,
				app.Div().
					Class("pf-c-alert__action-group").
					Body(
						app.Button().
							Class("pf-c-button pf-m-link pf-m-inline").
							Type("button").
							Text("Recover").
							OnClick(func(ctx app.Context, e app.Event) {
								c.Recover()
							}),
					),
			),
		)).Else(app.Span())
}
