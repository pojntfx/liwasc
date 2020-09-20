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
			app.Div().Class("pf-c-toolbar__content-section pf-u-flex-direction-column pf-u-flex-direction-row-on-lg").Body(
				app.Div().Class("pf-c-toolbar__item pf-u-w-100 pf-u-w-initial-on-lg pf-u-mb-sm pf-u-mb-0-on-lg").Body(
					&LabelCollectionComponent{Icon: "fas fa-network-wired", Title: "Subnets", Labels: c.Subnets},
				),
				app.Div().Class("pf-c-toolbar__item pf-u-w-100 pf-u-w-initial-on-lg pf-u-mb-sm pf-u-mb-0-on-lg").Body(
					&LabelCollectionComponent{Icon: "fas fa-microchip", Title: "Device", Labels: []string{c.Device}},
				),
				app.Div().Class("pf-c-toolbar__group pf-m-filter-group pf-u-w-100 pf-u-w-initial-on-lg pf-u-ml-auto-on-lg pf-u-mb-sm pf-u-mb-0-on-lg").Body(
					app.Div().Class("pf-c-toolbar__item pf-m-search-filter pf-u-w-100").Body(
						app.Div().Class("pf-c-search-input").Body(
							app.Span().Class("pf-c-search-input__text").Body(
								app.Span().Class("pf-c-search-input__icon").Body(
									app.I().Class("fas fa-search fa-fw"),
								),
							),
							app.Input().Class("pf-c-search-input__text-input").Type("search").Placeholder("Find by MAC, IP, ..."),
						),
					),
				),
				app.Div().Class("pf-c-divider pf-m-vertical pf-m-inset-md"),
				app.Div().Class("pf-c-toolbar__item").Body(
					app.Button().Class("pf-c-button pf-m-primary").Body(
						app.Span().Class("pf-c-button__icon pf-m-start").Body(
							app.I().Class("fas fa-rocket"),
						),
						app.Text("Trigger Scan"),
					),
				),
			),
		),
	)
}
