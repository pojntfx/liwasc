package shells

import (
	"fmt"
	"strconv"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/pojntfx/liwasc/pkg/components"
	"github.com/pojntfx/liwasc/pkg/providers"
)

type DataShell struct {
	app.Compo

	nodeScanTimeout    int64
	portScanTimeout    int64
	nodeScanMACAddress string

	nodeWakeTimeout    int64
	nodeWakeMACAddress string

	selectedMACAddress string

	userMenuExpanded        bool
	overflowMenuExpanded    bool
	aboutDialogOpen         bool
	settingsDialogOpen      bool
	notificationsDrawerOpen bool
	metadataDialogOpen      bool

	Network  providers.Network
	UserInfo oidc.UserInfo

	TriggerNetworkScan func(nodeScanTimeout int64, portScanTimeout int64, macAddress string)
	StartNodeWake      func(nodeWakeTimeout int64, macAddress string)
	Logout             func()

	Error   error
	Recover func()
}

const (
	// Names and IDs
	nodeScanTimeoutName    = "nodeScanTimeout"
	portScanTimeoutName    = "portScanTimeout"
	nodeScanMACAddressName = "nodeScanMACAddressTimeout"

	nodeWakeTimeoutName    = "nodeWakeTimeout"
	nodeWakeMACAddressName = "nodeWakeMACAddressTimeout"

	// Default values
	defaultNodeScanTimeout = 500
	defaultPortScanTimeout = 10
	allMACAddresses        = "ff:ff:ff:ff"

	defaultNodeWakeTimeout    = 600000
	defaultNodeWakeMACAddress = ""
)

func (c *DataShell) Render() app.UI {
	selectedNode := providers.Node{}
	if c.selectedMACAddress != "" {
		// Find selected node
		for _, candidate := range c.Network.Nodes {
			if candidate.MACAddress == c.selectedMACAddress {
				selectedNode = candidate

				break
			}
		}

		// If selected node could not be found, clear selected MAC address
		if selectedNode.MACAddress == "" {
			c.selectedMACAddress = ""
		}
	}

	return app.Div().
		Class("pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-page").
				ID("page-layout-horizontal-nav").
				Aria("hidden", c.aboutDialogOpen || c.settingsDialogOpen).
				Body(
					app.A().
						Class("pf-c-skip-to-content pf-c-button pf-m-primary").
						Href("#main-content-page-layout-horizontal-nav").
						Text(
							"Skip to content",
						),
					app.Header().
						Class("pf-c-page__header").
						Body(
							app.Div().
								Class("pf-c-page__header-brand").
								Body(
									app.A().
										Href("#").
										Class("pf-c-page__header-brand-link").
										Body(
											app.Img().
												Class("pf-c-brand pf-x-c-brand--nav").
												Src("/web/logo.svg").
												Alt("liwasc Logo"),
										),
								),
							app.Div().
								Class("pf-c-page__header-tools").
								Body(
									app.Div().
										Class("pf-c-page__header-tools-group").
										Body(
											app.Div().
												Class("pf-c-page__header-tools-group").
												Body(
													app.Div().
														Class(func() string {
															classes := "pf-c-page__header-tools-item"

															if c.notificationsDrawerOpen {
																classes += " pf-m-selected"
															}

															return classes
														}()).
														Body(
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("label", "Unread notifications").
																Aria("expanded", false).
																OnClick(func(ctx app.Context, e app.Event) {
																	c.dispatch(func() {
																		c.notificationsDrawerOpen = !c.notificationsDrawerOpen
																		c.settingsDialogOpen = false
																		c.overflowMenuExpanded = false
																	})
																}).
																Body(
																	app.Span().
																		Class("pf-c-notification-badge").
																		Body(
																			app.I().
																				Class("pf-icon-bell").
																				Aria("hidden", true),
																		),
																),
														),
													app.Div().
														Class("pf-c-page__header-tools-item pf-m-hidden pf-m-visible-on-lg").
														Body(
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("label", "Settings").
																OnClick(func(ctx app.Context, e app.Event) {
																	c.dispatch(func() {
																		c.settingsDialogOpen = true
																		c.overflowMenuExpanded = false
																	})
																}).
																Body(
																	app.I().
																		Class("fas fa-cog").
																		Aria("hidden", true),
																),
														),
													app.Div().Class("pf-c-page__header-tools-item").
														Body(
															app.Div().
																Class(func() string {
																	classes := "pf-c-dropdown"

																	if c.overflowMenuExpanded {
																		classes += " pf-m-expanded"
																	}

																	return classes
																}()).
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("page-default-nav-example-dropdown-kebab-1-button").
																		Aria("expanded", c.overflowMenuExpanded).Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v pf-u-display-none-on-lg").
																				Aria("hidden", true),
																			app.I().
																				Class("fas fa-question-circle pf-u-display-none pf-u-display-inline-block-on-lg").
																				Aria("hidden", true),
																		).OnClick(func(ctx app.Context, e app.Event) {
																		c.dispatch(func() {
																			c.overflowMenuExpanded = !c.overflowMenuExpanded
																			c.userMenuExpanded = false
																		})
																	}),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("aria-labelledby", "page-default-nav-example-dropdown-kebab-1-button").
																		Hidden(!c.overflowMenuExpanded).
																		Body(
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-button pf-c-dropdown__menu-item pf-u-display-none-on-lg").
																						Type("button").
																						OnClick(func(ctx app.Context, e app.Event) {
																							c.dispatch(func() {
																								c.settingsDialogOpen = true
																								c.overflowMenuExpanded = false
																							})
																						}).
																						Body(
																							app.Span().
																								Class("pf-c-button__icon pf-m-start").
																								Body(
																									app.I().
																										Class("fas fa-cog").
																										Aria("hidden", true),
																								),
																							app.Text("Settings"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider pf-u-display-none-on-lg").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("https://github.com/pojntfx/liwasc/wiki").
																						Text("Documentation").
																						Target("_blank"),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-button pf-c-dropdown__menu-item").
																						Type("button").
																						OnClick(func(ctx app.Context, e app.Event) {
																							c.dispatch(func() {
																								c.aboutDialogOpen = true
																								c.overflowMenuExpanded = false
																							})
																						}).
																						Text("About"),
																				),
																			app.Li().
																				Class("pf-c-divider pf-u-display-none-on-md").
																				Aria("role", "separator"),
																			app.Li().
																				Class("pf-u-display-none-on-md").
																				Body(
																					app.Button().
																						Class("pf-c-button pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Span().
																								Class("pf-c-button__icon pf-m-start").
																								Body(
																									app.I().
																										Class("fas fa-sign-out-alt").
																										Aria("hidden", true),
																								),
																							app.Text("Logout"),
																						).
																						OnClick(func(ctx app.Context, e app.Event) {
																							go c.Logout()
																						}),
																				),
																		),
																),
														),
													app.Div().
														Class("pf-c-page__header-tools-item pf-m-hidden pf-m-visible-on-md").
														Body(
															app.Div().
																Class(func() string {
																	classes := "pf-c-dropdown"

																	if c.userMenuExpanded {
																		classes += " pf-m-expanded"
																	}

																	return classes
																}()).
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("page-layout-horizontal-nav-dropdown-kebab-2-button").
																		Aria("expanded", c.userMenuExpanded).
																		Type("button").
																		Body(
																			app.Span().
																				Class("pf-c-dropdown__toggle-text").
																				Text(c.UserInfo.Email),
																			app.
																				Span().
																				Class("pf-c-dropdown__toggle-icon").
																				Body(
																					app.I().
																						Class("fas fa-caret-down").
																						Aria("hidden", true),
																				),
																		).OnClick(func(ctx app.Context, e app.Event) {
																		c.dispatch(func() {
																			c.userMenuExpanded = !c.userMenuExpanded
																			c.overflowMenuExpanded = false
																		})
																	}),
																	app.Ul().
																		Class("pf-c-dropdown__menu").
																		Aria("labelledby", "page-layout-horizontal-nav-dropdown-kebab-2-button").
																		Hidden(!c.userMenuExpanded).
																		Body(
																			app.Li().Body(
																				app.Button().
																					Class("pf-c-button pf-c-dropdown__menu-item").
																					Type("button").
																					Body(
																						app.Span().
																							Class("pf-c-button__icon pf-m-start").
																							Body(
																								app.I().
																									Class("fas fa-sign-out-alt").
																									Aria("hidden", true),
																							),
																						app.Text("Logout"),
																					).
																					OnClick(func(ctx app.Context, e app.Event) {
																						go c.Logout()
																					}),
																			),
																		),
																),
														),
												),
											app.Img().Class("pf-c-avatar").Src("https://www.gravatar.com/avatar/db856df33fa4f4bce441819f604c90d5?s=150").Alt("Avatar image"),
										),
								),
						),
					app.Div().
						Class("pf-c-page__drawer").
						Body(
							app.Div().
								Class(func() string {
									classes := "pf-c-drawer"

									if c.notificationsDrawerOpen {
										classes += " pf-m-expanded"
									}

									return classes
								}()).
								Body(
									app.Div().
										Class("pf-c-drawer__main").
										Body(
											app.Div().
												Class("pf-c-drawer__content").
												Body(
													app.Div().Class("pf-c-drawer__body").Body(
														app.Main().
															Class("pf-c-page__main pf-u-h-100").
															ID("main-content-page-layout-horizontal-nav").
															TabIndex(-1).
															Body(
																app.Section().
																	Class("pf-c-page__main-section pf-m-no-padding").
																	Body(
																		// Primary-detail
																		&components.Inspector{
																			Open: c.selectedMACAddress != "",
																			Close: func() {
																				c.dispatch(func() {
																					c.selectedMACAddress = ""
																				})
																			},
																			StartNodeWake: func() {
																				go c.StartNodeWake(c.nodeWakeTimeout, selectedNode.MACAddress)
																			},
																			TriggerNetworkScan: func() {
																				go c.TriggerNetworkScan(c.nodeScanTimeout, c.portScanTimeout, selectedNode.MACAddress)
																			},
																			Header: []app.UI{
																				// Data status
																				&components.Status{
																					Error:   c.Error,
																					Recover: c.Recover,
																				},
																				// Toolbar
																				app.Div().
																					Class("pf-c-toolbar pf-m-page-insets").
																					Body(
																						app.Div().
																							Class("pf-c-toolbar__content").
																							Body(
																								app.Div().
																									Class("pf-c-toolbar__content-section pf-m-nowrap pf-u-display-none pf-u-display-flex-on-lg").
																									Body(
																										app.Div().
																											Class("pf-c-toolbar__item").
																											Body(
																												// Data actions
																												&components.ProgressButton{
																													Loading: c.Network.NodeScanRunning,
																													Icon:    "fas fa-rocket",
																													Text:    "Trigger Scan",

																													OnClick: func(ctx app.Context, e app.Event) {
																														go c.TriggerNetworkScan(c.nodeScanTimeout, c.portScanTimeout, "")
																													},
																												},
																											),
																										app.Div().
																											Class("pf-c-toolbar__item").
																											Body(
																												app.Div().
																													Class("pf-c-label-group pf-m-category").
																													Body(
																														app.Div().
																															Class("pf-c-label-group__main").
																															Body(
																																app.Span().
																																	Class("pf-c-label-group__label").
																																	Aria("hidden", true).
																																	ID("last-scan").
																																	Body(
																																		app.I().
																																			Class("fas fa-history pf-u-mr-xs").
																																			Aria("hidden", true),
																																		app.Text("Last Scan"),
																																	),
																																app.Ul().
																																	Class("pf-c-label-group__list").
																																	Aria("role", "list").
																																	Aria("labelledby", "last-scan").
																																	Body(
																																		app.Li().
																																			Class("pf-c-label-group__list-item").
																																			Body(
																																				app.Span().
																																					Class("pf-c-label").
																																					Body(
																																						app.Span().
																																							Class("pf-c-label__content").
																																							Body(
																																								app.Text(c.Network.LastNodeScanDate),
																																							),
																																					),
																																			),
																																	),
																															),
																													),
																											),
																										app.Div().Class("pf-c-toolbar__item pf-m-pagination").Body(
																											app.Div().
																												Class("pf-c-label-group pf-m-category pf-u-mr-md").
																												Body(
																													app.Div().
																														Class("pf-c-label-group__main").
																														Body(
																															app.Span().
																																Class("pf-c-label-group__label").
																																Aria("hidden", true).
																																ID("subnets").
																																Body(
																																	app.I().
																																		Class("fas fa-network-wired pf-u-mr-xs").
																																		Aria("hidden", true),
																																	app.Text("Subnets"),
																																),
																															app.Ul().
																																Class("pf-c-label-group__list").
																																Aria("role", "list").
																																Aria("labelledby", "subnets").
																																Body(
																																	app.Range(c.Network.ScannerMetadata.Subnets).Slice(func(i int) app.UI {
																																		return app.Li().
																																			Class("pf-c-label-group__list-item").
																																			Body(
																																				app.Span().
																																					Class("pf-c-label").
																																					Body(
																																						app.Span().
																																							Class("pf-c-label__content").
																																							Body(
																																								app.Text(c.Network.ScannerMetadata.Subnets[i]),
																																							),
																																					),
																																			)
																																	}),
																																),
																														),
																												),
																											app.Div().
																												Class("pf-c-label-group pf-m-category").
																												Body(
																													app.Div().
																														Class("pf-c-label-group__main").
																														Body(
																															app.Span().
																																Class("pf-c-label-group__label").
																																Aria("hidden", true).
																																ID("device").
																																Body(
																																	app.I().
																																		Class("fas fa-microchip pf-u-mr-xs").
																																		Aria("hidden", true),
																																	app.Text("Device"),
																																),
																															app.Ul().
																																Class("pf-c-label-group__list").
																																Aria("role", "list").
																																Aria("labelledby", "device").
																																Body(
																																	app.Li().
																																		Class("pf-c-label-group__list-item").
																																		Body(
																																			app.Span().
																																				Class("pf-c-label").
																																				Body(
																																					app.Span().
																																						Class("pf-c-label__content").
																																						Body(
																																							app.Text(c.Network.ScannerMetadata.Device),
																																						),
																																				),
																																		),
																																),
																														),
																												),
																										),
																									),
																								app.Div().
																									Class("pf-c-toolbar__content-section pf-m-nowrap pf-u-display-flex pf-u-display-none-on-lg").
																									Body(
																										app.Div().
																											Class("pf-c-toolbar__item").
																											Body(
																												// Data actions
																												&components.ProgressButton{
																													Loading: c.Network.NodeScanRunning,
																													Icon:    "fas fa-rocket",
																													Text:    "Trigger Scan",

																													OnClick: func(ctx app.Context, e app.Event) {
																														go c.TriggerNetworkScan(c.nodeScanTimeout, c.portScanTimeout, "")
																													},
																												},
																											),
																										app.Div().
																											Class("pf-c-toolbar__item pf-m-pagination").
																											Body(
																												app.Button().
																													Class("pf-c-button pf-m-plain").
																													Type("button").
																													Aria("label", "Metadata").
																													OnClick(func(ctx app.Context, e app.Event) {
																														c.dispatch(func() {
																															c.metadataDialogOpen = true
																														})
																													}).
																													Body(
																														app.I().
																															Class("fas fa-info-circle").
																															Aria("hidden", true),
																													),
																											),
																									),
																							),
																					),
																			},
																			Body:
																			// Data output
																			app.Table().
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
																										Text("Powered On"),
																									app.Th().
																										Aria("role", "columnheader").
																										Scope("col").
																										Text("MAC Address"),
																									app.Th().
																										Aria("role", "columnheader").
																										Scope("col").
																										Text("IP Address"),
																									app.Th().
																										Aria("role", "columnheader").
																										Scope("col").
																										Text("Vendor"),
																									app.Th().
																										Aria("role", "columnheader").
																										Scope("col").
																										Text("Services and Ports"),
																								),
																						),
																					app.TBody().
																						Class("pf-x-u-border-t-0").
																						Aria("role", "rowgroup").
																						Body(
																							app.If(
																								len(c.Network.Nodes) == 0 && !c.Network.NodeScanRunning,
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
																																				&components.ProgressButton{
																																					Loading: c.Network.NodeScanRunning,
																																					Icon:    "fas fa-rocket",
																																					Text:    "Trigger Scan",

																																					OnClick: func(ctx app.Context, e app.Event) {
																																						go c.TriggerNetworkScan(c.nodeScanTimeout, c.portScanTimeout, "")
																																					},
																																				},
																																			),
																																	),
																															),
																													),
																											),
																									),
																							).Else(
																								app.Range(c.Network.Nodes).Slice(func(i int) app.UI {
																									return app.Tr().
																										Class(func() string {
																											classes := "pf-m-hoverable"

																											if len(c.Network.Nodes) >= i && c.Network.Nodes[i].MACAddress == c.selectedMACAddress {
																												classes += " pf-m-selected"
																											}

																											return classes
																										}()).
																										Aria("role", "row").
																										OnClick(func(ctx app.Context, e app.Event) {
																											c.dispatch(func() {
																												// Reset selected node
																												if c.selectedMACAddress == c.Network.Nodes[i].MACAddress {
																													c.selectedMACAddress = ""

																													return
																												}

																												// Set selected node
																												c.selectedMACAddress = c.Network.Nodes[i].MACAddress
																											})
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
																																c.Network.Nodes[i].PoweredOn,
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
																															&components.Controlled{
																																Component: app.Input().
																																	Class("pf-c-switch__input").
																																	ID(fmt.Sprintf("node-row-%v", i)).
																																	Aria("label", "Node is off").
																																	Name(fmt.Sprintf("node-row-%v", i)).
																																	Type("checkbox").
																																	Checked(c.Network.Nodes[i].PoweredOn).
																																	Disabled(c.Network.Nodes[i].PoweredOn).
																																	OnClick(func(ctx app.Context, e app.Event) {
																																		e.Call("stopPropagation")

																																		go c.StartNodeWake(c.nodeWakeTimeout, c.Network.Nodes[i].MACAddress)
																																	}),
																																Properties: map[string]interface{}{
																																	"checked":  c.Network.Nodes[i].PoweredOn,
																																	"disabled": c.Network.Nodes[i].PoweredOn,
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
																																		c.Network.Nodes[i].NodeWakeRunning,
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
																																		c.Network.Nodes[i].NodeWakeRunning,
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
																												Text(c.Network.Nodes[i].MACAddress),
																											app.Td().
																												Aria("role", "cell").
																												DataSet("label", "IP Address").
																												Text(c.Network.Nodes[i].IPAddress),
																											app.Td().
																												Aria("role", "cell").
																												DataSet("label", "Vendor").
																												Text(func() string {
																													vendor := c.Network.Nodes[i].Vendor
																													if vendor == "" {
																														vendor = "Unregistered"
																													}

																													return vendor
																												}()),
																											app.Td().
																												Aria("role", "cell").
																												DataSet("label", "Services and Ports").
																												Body(
																													&components.ProgressButton{
																														Loading: c.Network.Nodes[i].PortScanRunning,
																														Icon:    "fas fa-sync",

																														OnClick: func(ctx app.Context, e app.Event) {
																															e.Call("stopPropagation")

																															go c.TriggerNetworkScan(c.nodeScanTimeout, c.portScanTimeout, c.Network.Nodes[i].MACAddress)
																														},
																													},
																													app.If(
																														len(c.Network.Nodes[i].Ports) > 0,
																														&components.PortList{
																															Ports: c.Network.Nodes[i].Ports,
																														},
																													).ElseIf(
																														c.Network.Nodes[i].PortScanRunning,
																														app.Text("No open ports found yet."),
																													).Else(
																														app.Text("No open ports found."),
																													),
																												),
																										)
																								}),
																							),
																						),
																				),
																			Node: selectedNode,
																		},
																	),
															),
													),
												),
											app.Div().
												Class("pf-c-drawer__panel").
												Body(
													app.Div().
														Class("pf-c-drawer__body pf-m-no-padding").
														Body(
															app.Div().
																Class("pf-c-notification-drawer").
																Body(
																	app.Div().
																		Class("pf-c-notification-drawer__header").
																		Body(
																			app.H1().
																				Class("pf-c-notification-drawer__header-title").
																				Text("Events"),
																		),
																	app.Div().Class("pf-c-notification-drawer__body").Body(
																		app.Ul().Class("pf-c-notification-drawer__list").Body(
																			app.Range(c.Network.Events).Slice(func(i int) app.UI {
																				return app.Li().Class("pf-c-notification-drawer__list-item pf-m-read pf-m-info").Body(
																					app.Div().Class("pf-c-notification-drawer__list-item-description").Text(
																						c.Network.Events[len(c.Network.Events)-1-i].Message,
																					),
																					app.Div().Class("pf-c-notification-drawer__list-item-timestamp").Text(
																						c.Network.Events[len(c.Network.Events)-1-i].Time,
																					),
																				)
																			}),
																		),
																	),
																),
														),
												),
										),
								),
						),
				),
			app.Div().
				Class(func() string {
					classes := "pf-c-backdrop"

					if !c.aboutDialogOpen {
						classes += " pf-u-display-none"
					}

					return classes
				}()).
				Body(
					app.Div().
						Class("pf-l-bullseye").
						Body(
							app.Div().
								Class("pf-c-about-modal-box").
								Aria("role", "dialog").
								Aria("modal", true).
								Aria("labelledby", "about-modal-title").
								Body(
									app.Div().
										Class("pf-c-about-modal-box__brand").
										Body(
											app.Img().
												Class("pf-c-about-modal-box__brand-image").
												Src("/web/logo.svg").
												Alt("liwasc Logo"),
										),
									app.Div().
										Class("pf-c-about-modal-box__close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close dialog").
												OnClick(func(ctx app.Context, e app.Event) {
													c.dispatch(func() {
														c.aboutDialogOpen = false
													})
												}).
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
									app.Div().
										Class("pf-c-about-modal-box__header").
										Body(
											app.H1().
												Class("pf-c-title pf-m-4xl").
												ID("about-modal-title").
												Text("liwasc"),
										),
									app.Div().Class("pf-c-about-modal-box__hero"),
									app.Div().
										Class("pf-c-about-modal-box__content").
										Body(
											app.Div().
												Class("pf-c-content").
												Body(
													app.Dl().
														Class("pf-c-content").
														Body(
															app.Dl().Body(
																app.Dt().Text("Frontend version"),
																app.Dd().Text("main"),
																app.Dt().Text("Backend version"),
																app.Dd().Text("main"),
															),
														),
												),
											app.P().
												Class("pf-c-about-modal-box__strapline").
												Text("Copyright © 2021 Felix Pojtinger and contributors (SPDX-License-Identifier: AGPL-3.0)"),
										),
								),
						),
				),
			app.Div().
				Class(func() string {
					classes := "pf-c-backdrop"

					if !c.settingsDialogOpen {
						classes += " pf-u-display-none"
					}

					return classes
				}()).
				Body(
					app.Div().
						Class("pf-l-bullseye").
						Body(
							app.Div().
								Class("pf-c-modal-box pf-m-sm").
								Aria("modal", true).
								Aria("labelledby", "modal-scroll-title").
								Aria("describedby", "modal-scroll-description").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Close dialog").
										OnClick(func(ctx app.Context, e app.Event) {
											c.dispatch(func() {
												c.settingsDialogOpen = false
											})
										}).
										Body(
											app.I().
												Class("fas fa-times").
												Aria("hidden", true),
										),
									app.Header().
										Class("pf-c-modal-box__header").
										Body(
											app.H1().
												Class("pf-c-modal-box__title").
												ID("modal-scroll-title").
												Text("Settings"),
										),
									app.Div().
										Class("pf-c-modal-box__body").
										Body(
											app.Form().
												Class("pf-c-form").
												ID("settings").
												Body(
													// Node Scan Timeout Input
													&components.FormGroup{
														Label: app.
															Label().
															For(nodeScanTimeoutName).
															Class("pf-c-form__label").
															Body(
																app.
																	Span().
																	Class("pf-c-form__label-text").
																	Text("Node Scan Timeout (in ms)"),
															),
														Input: &components.Controlled{
															Component: app.
																Input().
																Name(nodeScanTimeoutName).
																ID(nodeScanTimeoutName).
																Type("number").
																Required(true).
																Min(1).
																Step(1).
																Placeholder(strconv.Itoa(defaultNodeScanTimeout)).
																Class("pf-c-form-control").
																OnInput(func(ctx app.Context, e app.Event) {
																	v, err := strconv.Atoi(ctx.JSSrc.Get("value").String())
																	if err != nil || v == 0 {
																		c.Update()

																		return
																	}

																	c.nodeScanTimeout = int64(v)

																	c.Update()
																}),
															Properties: map[string]interface{}{
																"value": c.nodeScanTimeout,
															},
														},
														Required: true,
													},
													// Port Scan Timeout Input
													&components.FormGroup{
														Label: app.
															Label().
															For(portScanTimeoutName).
															Class("pf-c-form__label").
															Body(
																app.
																	Span().
																	Class("pf-c-form__label-text").
																	Text("Port Scan Timeout (in ms)"),
															),
														Input: &components.Controlled{
															Component: app.
																Input().
																Name(portScanTimeoutName).
																ID(portScanTimeoutName).
																Type("number").
																Required(true).
																Min(1).
																Step(1).
																Placeholder(strconv.Itoa(defaultPortScanTimeout)).
																Class("pf-c-form-control").
																OnInput(func(ctx app.Context, e app.Event) {
																	v, err := strconv.Atoi(ctx.JSSrc.Get("value").String())
																	if err != nil || v == 0 {
																		c.Update()

																		return
																	}

																	c.portScanTimeout = int64(v)

																	c.Update()
																}),
															Properties: map[string]interface{}{
																"value": c.portScanTimeout,
															},
														},
														Required: true,
													},
													// Node Wake Timeout Input
													&components.FormGroup{
														Label: app.
															Label().
															For(nodeWakeTimeoutName).
															Class("pf-c-form__label").
															Body(
																app.
																	Span().
																	Class("pf-c-form__label-text").
																	Text("Node Wake Timeout (in ms)"),
															),
														Input: &components.Controlled{
															Component: app.
																Input().
																Name(nodeWakeTimeoutName).
																ID(nodeWakeTimeoutName).
																Type("number").
																Required(true).
																Min(1).
																Step(1).
																Placeholder(strconv.Itoa(defaultNodeWakeTimeout)).
																Class("pf-c-form-control").
																OnInput(func(ctx app.Context, e app.Event) {
																	v, err := strconv.Atoi(ctx.JSSrc.Get("value").String())
																	if err != nil || v == 0 {
																		c.Update()

																		return
																	}

																	c.nodeWakeTimeout = int64(v)

																	c.Update()
																}),
															Properties: map[string]interface{}{
																"value": c.nodeWakeTimeout,
															},
														},
														Required: true,
													},
												).OnSubmit(func(ctx app.Context, e app.Event) {
												e.PreventDefault()

												c.dispatch(func() {
													c.settingsDialogOpen = false
												})
											}),
										),
									app.Footer().
										Class("pf-c-modal-box__footer").
										Body(
											app.Button().
												Class("pf-c-button pf-m-primary").
												Type("submit").
												Form("settings").
												Text("Save"),
											app.Button().
												Class("pf-c-button pf-m-link").
												Type("button").
												OnClick(func(ctx app.Context, e app.Event) {
													c.dispatch(func() {
														c.settingsDialogOpen = false
													})
												}).
												Text("Cancel"),
										),
								),
						),
				),
			app.Div().
				Class(func() string {
					classes := "pf-c-backdrop pf-u-display-none-on-lg"

					if !c.metadataDialogOpen {
						classes += " pf-u-display-none"
					}

					return classes
				}()).
				Body(
					app.Div().
						Class("pf-l-bullseye").
						Body(
							app.Div().
								Class("pf-c-modal-box pf-m-sm").
								Aria("modal", true).
								Aria("labelledby", "modal-scroll-title").
								Aria("describedby", "modal-scroll-description").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Close dialog").
										OnClick(func(ctx app.Context, e app.Event) {
											c.dispatch(func() {
												c.metadataDialogOpen = false
											})
										}).
										Body(
											app.I().
												Class("fas fa-times").
												Aria("hidden", true),
										),
									app.Header().
										Class("pf-c-modal-box__header").
										Body(
											app.H1().
												Class("pf-c-modal-box__title").
												ID("modal-scroll-title").
												Text("Metadata"),
										),
									app.Div().
										Class("pf-c-modal-box__body").
										Body(
											app.Dl().
												Class("pf-c-description-list").
												Body(
													app.Div().
														Class("pf-c-description-list__group").
														Body(
															app.Dt().
																Class("pf-c-description-list__term").
																Body(
																	app.Span().
																		Class("pf-c-description-list__text").
																		ID("last-scan-mobile").
																		Body(
																			app.I().
																				Class("fas fa-history pf-u-mr-xs").
																				Aria("hidden", true),
																			app.Text("Last Scan"),
																		),
																),
															app.Dd().
																Class("pf-c-description-list__description").
																Body(
																	app.Div().
																		Class("pf-c-description-list__text").
																		Body(
																			app.Ul().
																				Class("pf-c-label-group__list").
																				Aria("role", "list").
																				Aria("labelledby", "last-scan-mobile").
																				Body(
																					app.Li().
																						Class("pf-c-label-group__list-item").
																						Body(
																							app.Span().
																								Class("pf-c-label").
																								Body(
																									app.Span().
																										Class("pf-c-label__content").
																										Body(
																											app.Text(c.Network.LastNodeScanDate),
																										),
																								),
																						),
																				),
																		),
																),
														),
													app.Div().
														Class("pf-c-description-list__group").
														Body(
															app.Dt().
																Class("pf-c-description-list__term").
																Body(
																	app.Span().
																		Class("pf-c-description-list__text").
																		ID("subnets-mobile").
																		Body(
																			app.I().
																				Class("fas fa-network-wired pf-u-mr-xs").
																				Aria("hidden", true),
																			app.Text("Subnets"),
																		),
																),
															app.Dd().
																Class("pf-c-description-list__description").
																Body(
																	app.Div().
																		Class("pf-c-description-list__text").
																		Body(
																			app.Ul().
																				Class("pf-c-label-group__list").
																				Aria("role", "list").
																				Aria("labelledby", "subnets-mobile").
																				Body(
																					app.Range(c.Network.ScannerMetadata.Subnets).Slice(func(i int) app.UI {
																						return app.Li().
																							Class("pf-c-label-group__list-item").
																							Body(
																								app.Span().
																									Class("pf-c-label").
																									Body(
																										app.Span().
																											Class("pf-c-label__content").
																											Body(
																												app.Text(c.Network.ScannerMetadata.Subnets[i]),
																											),
																									),
																							)
																					}),
																				),
																		),
																),
														),
													app.Div().
														Class("pf-c-description-list__group").
														Body(
															app.Dt().
																Class("pf-c-description-list__term").
																Body(
																	app.Span().
																		Class("pf-c-description-list__text").
																		ID("device-mobile").
																		Body(
																			app.I().
																				Class("fas fa-microchip pf-u-mr-xs").
																				Aria("hidden", true),
																			app.Text("Device"),
																		),
																),
															app.Dd().
																Class("pf-c-description-list__description").
																Body(
																	app.Dd().
																		Class("pf-c-description-list__description").
																		Body(
																			app.Ul().
																				Class("pf-c-label-group__list").
																				Aria("role", "list").
																				Aria("labelledby", "device-mobile").
																				Body(
																					app.Li().
																						Class("pf-c-label-group__list-item").
																						Body(
																							app.Span().
																								Class("pf-c-label").
																								Body(
																									app.Span().
																										Class("pf-c-label__content").
																										Body(
																											app.Text(c.Network.ScannerMetadata.Device),
																										),
																								),
																						),
																				),
																		),
																),
														),
												),
										),
								),
						),
				),
		)
}

func (c *DataShell) OnMount(context app.Context) {
	// Initialize node scan form
	c.nodeScanTimeout = defaultNodeScanTimeout
	c.portScanTimeout = defaultPortScanTimeout
	c.nodeScanMACAddress = allMACAddresses

	// Initialize node wake form
	c.nodeWakeTimeout = defaultNodeWakeTimeout
	c.nodeWakeMACAddress = defaultNodeWakeMACAddress
}

func (c *DataShell) dispatch(action func()) {
	action()

	c.Update()
}
