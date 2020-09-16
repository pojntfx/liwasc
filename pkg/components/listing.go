package components

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type ListingComponent struct {
	app.Compo
	Nodes      []ListingNode
	OnRowClick func(int)
}

func (c *ListingComponent) Render() app.UI {
	return app.Table().Class("pf-c-table pf-m-grid-md").Body(
		app.THead().Body(
			app.Tr().Body(
				app.Th().Body(app.Text("Powered On")),
				app.Th().Body(app.Text("MAC Address")),
				app.Th().Body(app.Text("IP Address")),
				app.Th().Body(app.Text("Vendor")),
				app.Th().Body(app.Text("Servers and Ports")),
			),
		),
		app.TBody().Body(
			app.Range(c.Nodes).Slice(func(i int) app.UI {
				return app.Tr().Body(
					app.Td().Body(
						app.Label().Class("pf-c-switch").Body(
							app.Input().Class("pf-c-switch__input").Type("checkbox").Checked(c.Nodes[i].PoweredOn),
							app.Span().Class("pf-c-switch__toggle").Body(
								app.Span().Class("pf-c-switch__toggle-icon").Body(
									app.I().Class("fas fa-lightbulb"),
								),
							),
						),
					),
					app.Td().OnClick(func(ctx app.Context, e app.Event) { c.OnRowClick(i) }).Body(app.Text(c.Nodes[i].MACAddress)),
					app.Td().OnClick(func(ctx app.Context, e app.Event) { c.OnRowClick(i) }).Body(app.Text(c.Nodes[i].IPAddress)),
					app.Td().OnClick(func(ctx app.Context, e app.Event) { c.OnRowClick(i) }).Body(app.Text(c.Nodes[i].Vendor)),
					app.Td().OnClick(func(ctx app.Context, e app.Event) { c.OnRowClick(i) }).Body(
						app.Div().Class("pf-c-label-group").Body(
							app.Ul().Class("pf-c-label-group__list").Body(
								app.Range(c.Nodes[i].ServicesAndPorts).Slice(func(i2 int) app.UI {
									return app.Li().Class("pf-c-label-group__list-item").Body(
										app.Span().Class("pf-c-label").Body(
											app.Span().Class("pf-c-label__content").Body(
												app.Text(c.Nodes[i].ServicesAndPorts[i2]),
											),
										),
									)
								}),
							),
						),
					),
				)
			}),
		),
	)
}

type ListingNode struct {
	PoweredOn        bool
	MACAddress       string
	IPAddress        string
	Vendor           string
	ServicesAndPorts []string
}
