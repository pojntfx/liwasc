package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type DefinitionComponent struct {
	app.Compo
	Icon    string
	Title   string
	Content app.UI
}

func (c *DefinitionComponent) Render() app.UI {
	return app.Div().Class("pf-c-description-list__group").Body(
		app.Dt().Class("pf-c-description-list__term").Body(
			app.Span().Class("pf-c-description-list__text").Body(
				app.I().Class(fmt.Sprintf("%v pf-u-mr-xs", c.Icon)),
				app.Text(c.Title),
			),
		),
		app.Dd().Class("pf-c-description-list__description").Body(
			app.Div().Class("pf-c-description-list__text").Body(
				c.Content,
			),
		),
	)
}
