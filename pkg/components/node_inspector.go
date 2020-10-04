package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/liwasc-frontend-web/pkg/models"
)

type NodeInspectorComponent struct {
	app.Compo

	Node                 *models.Node
	ServicesAndPortsOpen bool
	DetailsOpen          bool
	SearchValue          string

	OnServicesAndPortsToggle func(ctx app.Context, e app.Event)
	OnDetailsToggle          func(ctx app.Context, e app.Event)
	OnSearchChange           func(string)
	OnReScanClick            func(ctx app.Context, e app.Event)
	OnServiceClick           func(int)
}

func (c *NodeInspectorComponent) Render() app.UI {
	return app.Div().Body(
		app.Dl().Class("pf-c-description-list pf-m-2-col pf-u-mb-md").Body(
			&DefinitionComponent{
				Title:   "IP Address",
				Icon:    "fas fa-globe",
				Content: app.Text(c.Node.IPAddress),
			},
			&DefinitionComponent{
				Title: "Vendor",
				Icon:  "fas fa-store-alt",
				Content: func() app.UI {
					if c.Node.Vendor != "" {
						return &SearchLinkComponent{Topic: c.Node.Vendor}
					}

					return app.Text("Unknown vendor")
				}(),
			},
		),
		&ExpandableSectionComponent{
			Open:          c.ServicesAndPortsOpen,
			OnToggleClick: c.OnServicesAndPortsToggle,
			Title:         "Services and Ports",
			Content: app.Div().Class("pf-u-text-align-center").Body(
				app.Div().Class("pf-c-search-input pf-u-mb-md").Body(
					app.Span().Class("pf-c-search-input__text").Body(
						app.Span().Class("pf-c-search-input__icon").Body(
							app.I().Class("fas fa-search fa-fw"),
						),
					),
					app.Input().Class("pf-c-search-input__text-input").Type("search").
						Placeholder("Find by name, port or protocol").
						Value(c.SearchValue).
						OnInput(func(ctx app.Context, e app.Event) { c.OnSearchChange(e.Get("target").Get("value").String()) }),
				),
				app.Ul().Class("pf-c-data-list pf-u-mb-md").Body(
					app.Range(c.Node.Services).Slice(func(i int) app.UI {
						return app.Li().Class("pf-c-data-list__item pf-m-selectable").Body(
							app.Div().Class("pf-c-data-list__item-row").Body(
								app.Div().Class("pf-c-data-list__item-content").Body(
									app.Div().Class("pf-c-data-list__cell pf-u-display-flex pf-u-justify-content-space-between").Body(
										app.Span().Text(func() string {
											serviceName := c.Node.Services[i].ServiceName

											if serviceName == "" {
												return "Non-Registered service"
											}

											return serviceName
										}()),
										app.Span().Class("pf-u-ml-md").Text(fmt.Sprintf("%v/%v", c.Node.Services[i].PortNumber, c.Node.Services[i].TransportProtocol)),
									),
								),
							).OnClick(func(ctx app.Context, e app.Event) {
								c.OnServiceClick(i)
							}),
						)
					}),
				),
				app.Button().Class("pf-c-button pf-m-secondary").Body(
					app.Span().Class("pf-c-button__icon pf-m-start").Body(
						app.I().Class("fas fa-sync"),
					),
					app.Text("Re-Scan For Ports"),
				),
			),
		},
		&ExpandableSectionComponent{
			Open:          c.DetailsOpen,
			OnToggleClick: c.OnDetailsToggle,
			Title:         "Details",
			Content: app.Dl().Class("pf-c-description-list pf-m-2-col pf-u-mb-md").Body(
				&DefinitionComponent{
					Title: "Registry",
					Icon:  "fas fa-list",
					Content: app.Text(func() string {
						if c.Node.Registry != "" {
							return c.Node.Registry
						}

						return "Unknown registry"
					}()),
				},
				&DefinitionComponent{
					Title: "Organization",
					Icon:  "fas fa-university",
					Content: func() app.UI {
						if c.Node.Organization != "" {
							return &SearchLinkComponent{Topic: c.Node.Organization}
						}

						return app.Text("Unknown organization")
					}(),
				},
				&DefinitionComponent{
					Title: "Registered Address",
					Icon:  "fas fa-map-marker-alt",
					Content: func() app.UI {
						if c.Node.Organization != "" {
							return &SearchLinkComponent{Topic: c.Node.Address}
						}

						return app.Text("Unknown address")
					}(),
				},
				&DefinitionComponent{
					Title: "Visible Address",
					Icon:  "fas fa-binoculars",
					Content: app.Div().Body(
						app.I().Class(fmt.Sprintf("fas %v pf-u-mr-xs", func() string {
							if c.Node.Visible {
								return "fas fa-eye"
							}

							return "fas fa-eye-slash"
						}())),
						app.Text(func() string {
							if c.Node.Visible {
								return "Visible"
							}

							return "Hidden"
						}()),
					),
				},
			),
		},
	)
}
