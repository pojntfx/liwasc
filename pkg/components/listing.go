package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type ListingComponent struct {
	app.Compo
	Nodes             []ListingNode
	OnRowClick        func(int)
	SelectedNode      int
	OnNodePowerToggle func(int)
}

func (c *ListingComponent) Render() app.UI {
	headers := []string{"Powered On", "MAC Address", "IP Address", "Vendor", "Servers and Ports"}

	return app.Table().Class("pf-c-table pf-m-grid-md").Body(
		app.THead().Body(
			app.Tr().Body(
				app.Th().Body(app.Text(headers[0])),
				app.Th().Body(app.Text(headers[1])),
				app.Th().Body(app.Text(headers[2])),
				app.Th().Body(app.Text(headers[3])),
				app.Th().Body(app.Text(headers[4])),
			),
		),
		app.TBody().Body(
			app.Range(c.Nodes).Slice(func(i int) app.UI {
				return app.Tr().Class(fmt.Sprintf("x__table__col--selectable %v", func() string {
					if i == c.SelectedNode {
						return "x__table__col--selectable--selected"
					}

					return ""
				}())).Body(
					app.Td().DataSet("label", headers[0]).Class(fmt.Sprintf("%v", func() string {
						if i == c.SelectedNode {
							return "x__table__col--selectable--selected--first"
						}

						return ""
					}())).Body(
						&OnOffSwitchComponent{On: c.Nodes[i].PoweredOn, OnToggle: func(ctx app.Context, e app.Event) { c.OnNodePowerToggle(i) }},
					),
					app.Td().DataSet("label", headers[1]).Class("x__table__col--selectable").OnClick(func(ctx app.Context, e app.Event) { c.OnRowClick(i) }).Body(app.Text(c.Nodes[i].MACAddress)),
					app.Td().DataSet("label", headers[2]).Class("x__table__col--selectable").OnClick(func(ctx app.Context, e app.Event) { c.OnRowClick(i) }).Body(app.Text(c.Nodes[i].IPAddress)),
					app.Td().DataSet("label", headers[3]).Class("x__table__col--selectable").OnClick(func(ctx app.Context, e app.Event) { c.OnRowClick(i) }).Body(app.Text(c.Nodes[i].Vendor)),
					app.Td().DataSet("label", headers[4]).Class("x__table__col--selectable").OnClick(func(ctx app.Context, e app.Event) { c.OnRowClick(i) }).Body(
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
