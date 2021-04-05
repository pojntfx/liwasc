package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/pojntfx/liwasc/pkg/providers"
)

type NodeTable struct {
	app.Compo

	Nodes              []providers.Node
	NodeScanRunning    bool
	SelectedMACAddress string

	SetSelectedMACAddress    func(macAddress string)
	TriggerFullNetworkScan   func()
	TriggerScopedNetworkScan func(macAddress string)
	StartNodeWake            func(macAddress string)
}

func (c *NodeTable) Render() app.UI {
	return app.Table().
		Class("pf-c-table pf-m-grid-md").
		Aria("role", "grid").
		Aria("label", "Nodes and their status").
		Body(
			app.THead().
				Body(
					app.Tr().
						Aria("role", "row").
						Body(
							app.Th().
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.I().Class("fas fa-plug pf-u-mr-xs").Aria("hidden", true),
									app.Text("Powered On"),
								),
							app.Th().
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.I().Class("fas fa-address-card pf-u-mr-xs").Aria("hidden", true),
									app.Text("MAC Address"),
								),
							app.Th().
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.I().Class("fas fa-network-wired pf-u-mr-xs").Aria("hidden", true),
									app.Text("IP Address"),
								),
							app.Th().
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.I().Class("fas fa-industry pf-u-mr-xs").Aria("hidden", true),
									app.Text("Vendor"),
								),
							app.Th().
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.I().Class("fas fa-server pf-u-mr-xs").Aria("hidden", true),
									app.Text("Ports and Services"),
								),
						),
				),
			app.TBody().
				Class("pf-x-u-border-t-0").
				Aria("role", "rowgroup").
				Body(
					app.If(
						len(c.Nodes) == 0 && !c.NodeScanRunning,
						app.Tr().
							Aria("role", "row").
							Body(
								app.Td().
									Aria("role", "cell").
									ColSpan(5).
									Body(
										app.Div().
											Class("pf-l-bullseye").
											Body(
												app.Div().
													Class("pf-c-empty-state pf-m-sm").
													Body(
														app.Div().
															Class("pf-c-empty-state__content").
															Body(
																app.I().
																	Class("fas fa-binoculars pf-c-empty-state__icon").
																	Aria("hidden", true),
																app.H2().
																	Class("pf-c-title pf-m-lg").
																	Text("No nodes here yet"),
																app.Div().
																	Class("pf-c-empty-state__body").
																	Text("Scan the network to find out what nodes are on it."),
																app.Div().
																	Class("pf-c-empty-state__primary").
																	Body(
																		// Data actions
																		&ProgressButton{
																			Loading: c.NodeScanRunning,
																			Icon:    "fas fa-rocket",
																			Text:    "Trigger Scan",

																			OnClick: func(ctx app.Context, e app.Event) {
																				c.TriggerFullNetworkScan()
																			},
																		},
																	),
															),
													),
											),
									),
							),
					).Else(
						app.Range(c.Nodes).Slice(func(i int) app.UI {
							return app.Tr().
								Class(func() string {
									classes := "pf-m-hoverable"

									if len(c.Nodes) >= i && c.Nodes[i].MACAddress == c.SelectedMACAddress {
										classes += " pf-m-selected"
									}

									return classes
								}()).
								Aria("role", "row").
								OnClick(func(ctx app.Context, e app.Event) {
									c.SetSelectedMACAddress(c.Nodes[i].MACAddress)
								}).
								Body(
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Powered On").
										Body(
											app.Label().
												Class("pf-c-switch pf-x-c-tooltip-wrapper").
												For(fmt.Sprintf("node-row-%v", i)).
												Body(
													app.If(
														c.Nodes[i].PoweredOn,
														app.Div().
															Class("pf-c-tooltip pf-x-c-tooltip pf-m-right").
															Aria("role", "tooltip").
															Body(
																app.Div().
																	Class("pf-c-tooltip__arrow"),
																app.Div().
																	Class("pf-c-tooltip__content").
																	Text("To turn this node off, please do so manually."),
															),
													),
													&Controlled{
														Component: app.Input().
															Class("pf-c-switch__input").
															ID(fmt.Sprintf("node-row-%v", i)).
															Aria("label", "Node is off").
															Name(fmt.Sprintf("node-row-%v", i)).
															Type("checkbox").
															Checked(c.Nodes[i].PoweredOn).
															Disabled(c.Nodes[i].PoweredOn).
															OnClick(func(ctx app.Context, e app.Event) {
																e.Call("stopPropagation")

																c.StartNodeWake(c.Nodes[i].MACAddress)
															}),
														Properties: map[string]interface{}{
															"checked":  c.Nodes[i].PoweredOn,
															"disabled": c.Nodes[i].PoweredOn,
														},
													},
													app.Span().
														Class("pf-c-switch__toggle").
														Body(
															app.Span().
																Class("pf-c-switch__toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-lightbulb").
																		Aria("hidden", true),
																),
														),
													app.Span().
														Class("pf-c-switch__label pf-m-on pf-l-flex pf-m-justify-content-center pf-m-align-items-center").
														ID(fmt.Sprintf("node-row-%v-on", i)).
														Aria("hidden", true).
														Body(
															app.If(
																c.Nodes[i].NodeWakeRunning,
																app.Span().
																	Class("pf-c-spinner pf-m-md").
																	Aria("role", "progressbar").
																	Aria("valuetext", "Loading...").
																	Body(
																		app.Span().Class("pf-c-spinner__clipper"),
																		app.Span().Class("pf-c-spinner__lead-ball"),
																		app.Span().Class("pf-c-spinner__tail-ball"),
																	),
															).Else(
																app.Text("On"),
															),
														),
													app.Span().
														Class("pf-c-switch__label pf-m-off pf-l-flex pf-m-justify-content-center pf-m-align-items-center").
														ID(fmt.Sprintf("node-row-%v-off", i)).
														Aria("hidden", true).
														Body(
															app.If(
																c.Nodes[i].NodeWakeRunning,
																app.Span().
																	Class("pf-c-spinner pf-m-md").
																	Aria("role", "progressbar").
																	Aria("valuetext", "Loading...").
																	Body(
																		app.Span().Class("pf-c-spinner__clipper"),
																		app.Span().Class("pf-c-spinner__lead-ball"),
																		app.Span().Class("pf-c-spinner__tail-ball"),
																	),
															).Else(
																app.Text("Off"),
															),
														),
												),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "MAC Address").
										Text(c.Nodes[i].MACAddress),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "IP Address").
										Text(c.Nodes[i].IPAddress),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Vendor").
										Text(func() string {
											vendor := c.Nodes[i].Vendor
											if vendor == "" {
												vendor = "Unregistered"
											}

											return vendor
										}()),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Ports and Services").
										Body(
											&ProgressButton{
												Loading: c.Nodes[i].PortScanRunning,
												Icon:    "fas fa-sync",

												OnClick: func(ctx app.Context, e app.Event) {
													e.Call("stopPropagation")

													c.TriggerScopedNetworkScan(c.Nodes[i].MACAddress)
												},
											},
											app.If(
												len(c.Nodes[i].Ports) > 0,
												&PortList{
													Ports: c.Nodes[i].Ports,
												},
											).ElseIf(
												c.Nodes[i].PortScanRunning,
												app.Text("No open ports found yet."),
											).Else(
												app.Text("No open ports found."),
											),
										),
								)
						}),
					),
				),
		)
}
