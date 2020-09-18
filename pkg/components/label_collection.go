package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type LabelCollectionComponent struct {
	app.Compo
	Icon   string
	Title  string
	Labels []string
}

func (c *LabelCollectionComponent) Render() app.UI {
	return app.Div().Class("pf-c-label-group pf-m-category pf-u-w-100").Body(
		app.Span().Class("pf-c-label-group__label").Body(
			app.I().Class(fmt.Sprintf("%v pf-u-mr-xs", c.Icon)),
			app.Text(c.Title),
		),
		app.Ul().Class("pf-c-label-group__list").Body(
			app.Range(c.Labels).Slice(func(i int) app.UI {
				return app.Li().Class("pf-c-label-group__list-item").Body(
					app.Span().Class("pf-c-label").Body(
						app.Span().Class("pf-c-label__content").Body(
							app.Text(c.Labels[i]),
						),
					),
				)
			}),
		),
	)
}
