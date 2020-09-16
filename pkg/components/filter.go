package components

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type FilterComponent struct {
	app.Compo
	Subnets []string
	Device  string
}

func (c *FilterComponent) Render() app.UI {
	return app.Div().Class("pf-c-toolbar").Body(
		app.Div().Class("pf-c-toolbar__content").Body(
			app.Div().Class("pf-c-toolbar__content-section").Body(
				app.Div().Class("pf-c-toolbar__item").Body(
					&LabelCollectionComponent{Icon: "fas fa-network-wired", Title: "Subnets", Labels: c.Subnets},
				),
				app.Div().Class("pf-c-toolbar__item").Body(
					&LabelCollectionComponent{Icon: "fas fa-microchip", Title: "Device", Labels: []string{c.Device}},
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
