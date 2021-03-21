package shells

import (
	"strconv"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
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

	userMenuExpanded     bool
	overflowMenuExpanded bool
	aboutDialogOpen      bool
	settingsDialogOpen   bool

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
	defaultPortScanTimeout = 50
	allMACAddresses        = "ff:ff:ff:ff"

	defaultNodeWakeTimeout    = 1000
	defaultNodeWakeMACAddress = ""
)

func (c *DataShell) Render() app.UI {
	return app.Div().Body(
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
													Class("pf-c-page__header-tools-item").
													Body(
														app.Button().
															Class("pf-c-button pf-m-plain").
															Type("button").
															Aria("label", "Unread notifications").
															Aria("expanded", false).
															Body(
																app.I().
																	Class("pf-icon-bell").
																	Aria("hidden", true),
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
				app.Main().
					Class("pf-c-page__main").
					TabIndex(-1).
					ID("main-content-page-layout-horizontal-nav").
					Body(
						app.Section().Class("pf-c-page__main-section pf-m-limit-width").Body(
							app.Div().Class("pf-c-page__main-section").Body(
								// Toolbar
								app.Div().
									Class("pf-c-toolbar pf-m-page-insets").
									Body(
										app.Div().
											Class("pf-c-toolbar__content").
											Body(
												app.Div().
													Class("pf-c-toolbar__content-section pf-m-nowrap").
													Body(
														app.Div().
															Class("pf-c-toolbar__item").
															Body(
																// Data actions
																app.
																	Button().
																	Type("submit").
																	Class(func() string {
																		classes := "pf-c-button pf-m-primary"

																		if c.Network.NodeScanRunning {
																			classes += " pf-m-progress pf-m-in-progress"
																		}

																		return classes
																	}()).
																	OnClick(func(ctx app.Context, e app.Event) {
																		go c.TriggerNetworkScan(c.nodeScanTimeout, c.portScanTimeout, "")
																	}).
																	Body(
																		app.If(c.Network.NodeScanRunning,
																			app.Span().
																				Class("pf-c-button__progress").
																				Body(
																					app.Span().
																						Class("pf-c-spinner pf-m-md").
																						Aria("role", "progressbar").
																						Aria("valuetext", "Loading...").
																						Body(
																							app.Span().Class("pf-c-spinner__clipper"),
																							app.Span().Class("pf-c-spinner__lead-ball"),
																							app.Span().Class("pf-c-spinner__tail-ball"),
																						),
																				)).Else(
																			app.Span().
																				Class("pf-c-button__icon pf-m-start").
																				Body(
																					app.I().
																						Class("fas fa-rocket").
																						Aria("hidden", true),
																				),
																		),
																		app.Text("Trigger Scan"),
																	),
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
											),
									),
								// Data status
								&components.Status{
									Error:   c.Error,
									Recover: c.Recover,
								},
								// Data output
								app.Ul().
									Class("pf-c-list").
									Body(
										&components.JSONDisplay{
											Object: c.Network.Events,
										},
										app.Range(c.Network.Nodes).Slice(func(i int) app.UI {
											return app.Li().Body(
												app.
													Button().
													Type("button").
													Class(func() string {
														classes := "pf-c-button pf-m-primary"

														if c.Network.Nodes[i].PortScanRunning {
															classes += " pf-m-progress pf-m-in-progress"
														}

														return classes
													}()).
													OnClick(func(ctx app.Context, e app.Event) {
														go c.TriggerNetworkScan(c.nodeScanTimeout, c.portScanTimeout, c.Network.Nodes[i].MACAddress)
													}).
													Body(
														app.If(c.Network.Nodes[i].PortScanRunning,
															app.Span().
																Class("pf-c-button__progress").
																Body(
																	app.Span().
																		Class("pf-c-spinner pf-m-md").
																		Aria("role", "progressbar").
																		Aria("valuetext", "Loading...").
																		Body(
																			app.Span().Class("pf-c-spinner__clipper"),
																			app.Span().Class("pf-c-spinner__lead-ball"),
																			app.Span().Class("pf-c-spinner__tail-ball"),
																		),
																)),
														app.Text("Scan this node"),
													),
												app.
													Button().
													Type("button").
													Class("pf-c-button pf-m-primary").
													OnClick(func(ctx app.Context, e app.Event) {
														go c.StartNodeWake(c.nodeWakeTimeout, c.Network.Nodes[i].MACAddress)
													}).
													Text("Wake this node"),
												&components.JSONDisplay{
													Object: c.Network.Nodes[i],
												},
											)
										}),
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
											Text("Copyright © 2021 Felicitas Pojtinger and contributors (SPDX-License-Identifier: AGPL-3.0)"),
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
														Value: c.nodeScanTimeout,
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
														Value: c.portScanTimeout,
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
														Value: c.nodeWakeTimeout,
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
