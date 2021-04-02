package components

import (
	"encoding/json"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type JSONDisplay struct {
	app.Compo

	Object interface{}
}

func (c *JSONDisplay) Render() app.UI {
	output, err := json.MarshalIndent(c.Object, "", "    ")
	if err != nil {
		panic(err)
	}

	return app.Div().Class("pf-c-code-editor pf-m-read-only").Body(
		app.Div().Class("pf-c-code-editor__header").Body(
			app.Div().Class("pf-c-code-editor__tab").Body(
				app.Span().Class("pf-c-code-editor__tab-icon").Body(
					app.I().Class("fas fa-code"),
				),
				app.Span().Class("pf-c-code-editor__tab-text").Text(
					"JSON",
				),
			),
		),
		app.Div().Class("pf-c-code-editor__main").Body(
			app.Div().Class("pf-c-code-editor__code pf-x-m-overflow-x-auto").Body(
				app.Pre().Text(string(output)),
			),
		),
	)
}
