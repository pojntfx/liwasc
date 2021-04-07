package components

import (
	"strconv"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/maxence-charriere/go-app/v8/pkg/app"
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
	portFilter              string
	selectedPort            string

	Network  providers.Network
	UserInfo oidc.UserInfo

	TriggerNetworkScan func(nodeScanTimeout int64, portScanTimeout int64, macAddress string)
	StartNodeWake      func(nodeWakeTimeout int64, macAddress string)
	Logout             func()

	Error   error
	Recover func()
	Ignore  func()
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

		// If selected node could not be found, clear selected MAC address and port filter
		if selectedNode.MACAddress == "" {
			c.selectedMACAddress = ""
			c.portFilter = ""
			c.selectedPort = ""
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
					&Navbar{
						NotificationsDrawerOpen: c.notificationsDrawerOpen,
						ToggleNotificationsDrawerOpen: func() {
							c.dispatch(func() {
								c.notificationsDrawerOpen = !c.notificationsDrawerOpen
								c.settingsDialogOpen = false
								c.overflowMenuExpanded = false
							})
						},

						ToggleSettings: func() {
							c.dispatch(func() {
								c.settingsDialogOpen = true
								c.overflowMenuExpanded = false
							})
						},
						ToggleAbout: func() {
							c.dispatch(func() {
								c.aboutDialogOpen = true
								c.overflowMenuExpanded = false
							})
						},

						OverflowMenuExpanded: c.overflowMenuExpanded,
						ToggleOverflowMenuExpanded: func() {
							c.dispatch(func() {
								c.overflowMenuExpanded = !c.overflowMenuExpanded
								c.userMenuExpanded = false
							})
						},

						UserMenuExpanded: c.userMenuExpanded,
						ToggleUserMenuExpanded: func() {
							c.dispatch(func() {
								c.userMenuExpanded = !c.userMenuExpanded
								c.overflowMenuExpanded = false
							})
						},

						UserEmail: c.UserInfo.Email,
						Logout: func() {
							go c.Logout()
						},
					},
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
																		&Inspector{
																			Open: c.selectedMACAddress != "",
																			Close: func() {
																				c.dispatch(func() {
																					c.selectedMACAddress = ""
																					c.portFilter = ""
																					c.selectedPort = ""
																				})
																			},
																			StartNodeWake: func() {
																				go c.StartNodeWake(c.nodeWakeTimeout, selectedNode.MACAddress)
																			},
																			TriggerNetworkScan: func() {
																				go c.TriggerNetworkScan(c.nodeScanTimeout, c.portScanTimeout, selectedNode.MACAddress)
																			},
																			Header: []app.UI{
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
																												&ProgressButton{
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
																												&ProgressButton{
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
																			Body: &NodeTable{
																				Nodes:              c.Network.Nodes,
																				NodeScanRunning:    c.Network.NodeScanRunning,
																				SelectedMACAddress: c.selectedMACAddress,

																				SetSelectedMACAddress: func(s string) {
																					c.dispatch(func() {
																						// Reset port filter and selected port
																						c.portFilter = ""
																						c.selectedPort = ""

																						// Reset selected node
																						if c.selectedMACAddress == s {
																							c.selectedMACAddress = ""

																							return
																						}

																						// Set selected node
																						c.selectedMACAddress = s
																					})
																				},
																				TriggerFullNetworkScan: func() {
																					go c.TriggerNetworkScan(c.nodeScanTimeout, c.portScanTimeout, "")
																				},
																				TriggerScopedNetworkScan: func(s string) {
																					go c.TriggerNetworkScan(c.nodeScanTimeout, c.portScanTimeout, s)
																				},
																				StartNodeWake: func(s string) {
																					go c.StartNodeWake(c.nodeWakeTimeout, s)
																				},
																			},
																			Node:       selectedNode,
																			PortFilter: c.portFilter,
																			SetPortFilter: func(s string) {
																				c.dispatch(func() {
																					c.portFilter = s
																				})
																			},
																			SelectedPort: c.selectedPort,
																			SetSelectedPort: func(s string) {
																				c.dispatch(func() {
																					c.selectedPort = s
																				})
																			},
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
			app.Ul().
				Class("pf-c-alert-group pf-m-toast").
				Body(
					app.If(
						c.Error != nil,
						app.Li().
							Class("pf-c-alert-group__item").
							Body(
								&Status{
									Error:       c.Error,
									ErrorText:   "Fatal Error",
									Recover:     c.Recover,
									RecoverText: "Reconnect",
									Ignore:      c.Ignore,
								},
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
													&FormGroup{
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
														Input: &Controlled{
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
													&FormGroup{
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
														Input: &Controlled{
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
													&FormGroup{
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
														Input: &Controlled{
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
