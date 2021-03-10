package experimental

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type StatusComponent struct {
	app.Compo

	Error   error
	Recover func()
}

func (c *StatusComponent) Render() app.UI {
	// Display the error message if error != nil
	errorMessage := ""
	if c.Error != nil {
		errorMessage = c.Error.Error()
	}

	return app.If(c.Error != nil, app.Div().Body(
		app.P().Text(errorMessage),
		app.If(c.Recover != nil, app.Button().Text("Recover").OnClick(func(ctx app.Context, e app.Event) {
			c.Recover()
		})),
	)).Else(app.Span())
}
