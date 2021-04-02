package components

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type Property struct {
	app.Compo

	Key   string
	Value string
}

func (c *Property) Render() app.UI {
	return app.Div().
		Class("pf-c-description-list__group").
		Body(
			app.Dt().
				Class("pf-c-description-list__term").
				Body(
					app.Span().
						Class("pf-c-description-list__text").
						Text(c.Key),
				),
			app.Dd().
				Class("pf-c-description-list__description").
				Body(
					app.Div().
						Class("pf-c-description-list__text").
						Body(
							app.If(
								c.Value == "",
								app.Text("Unregistered"),
							).Else(
								app.Text(c.Value),
							),
						),
				),
		)
}
