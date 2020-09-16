package components

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type FilterComponent struct {
	app.Compo
	Subnets []string
}

func (c *FilterComponent) Render() app.UI {
	return app.Div().Class("pf-c-toolbar").Body(
		app.Div().Class("pf-c-toolbar__content").Body(
			app.Div().Class("pf-c-toolbar__content-section").Body(
				app.Div().Class("pf-c-toolbar__item").Body(
					app.Div().Class("pf-c-label-group pf-m-category").Body(
						app.Span().Class("pf-c-label-group__label").Body(
							app.I().Class("fas fa-network-wired pf-u-mr-xs"),
							app.Text("Subnets"),
						),
						app.Ul().Class("pf-c-label-group__list").Body(
							app.Range(c.Subnets).Slice(func(i int) app.UI {
								return app.Li().Class("pf-c-label-group__list-item").Body(
									app.Span().Class("pf-c-label").Body(
										app.Span().Class("pf-c-label__content").Body(
											app.Text(c.Subnets[i]),
										),
									),
								)
							}),
						),
					),
				),
				app.Div().Class("pf-c-toolbar__item pf-m-pagination").Body(
					app.Button().Class("pf-c-button pf-m-primary").Body(
						app.Span().Class("pf-c-button__icon pf-m-start").Body(
							app.I().Class("fas fa-sync"),
						),
						app.Text("Re-Scan"),
					),
				),
			),
		),
	)
}
