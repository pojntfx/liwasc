package components

import (
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
	// Default values
	allMACAddresses           = "ff:ff:ff:ff"
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

	// Gather notifications
	notifications := []Notification{}
	for _, event := range c.Network.Events {
		notifications = append(notifications, Notification{
			Message: event.Message,
			Time:    event.Time.String(),
		})
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
																				&Toolbar{
																					NodeScanRunning: c.Network.NodeScanRunning,
																					TriggerFullNetworkScan: func() {

																						go c.TriggerNetworkScan(c.nodeScanTimeout, c.portScanTimeout, "")
																					},

																					LastNodeScanDate: c.Network.LastNodeScanDate.String(),
																					Subnets:          c.Network.ScannerMetadata.Subnets,
																					Device:           c.Network.ScannerMetadata.Device,

																					ToggleMetadataDialogOpen: func() {
																						c.dispatch(func() {
																							c.metadataDialogOpen = true
																						})
																					},
																				},
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
															&NotificationDrawer{
																Notifications: notifications,
															},
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
			&AboutModal{
				Open: c.aboutDialogOpen,
				Close: func() {
					c.dispatch(func() {
						c.aboutDialogOpen = false
					})
				},

				ID: "about-modal-title",

				LogoSrc: "/web/logo.svg",
				LogoAlt: "liwasc Logo",
				Title:   "liwasc",

				Body: app.Dl().
					Body(
						app.Dt().Text("Frontend version"),
						app.Dd().Text("main"),
						app.Dt().Text("Backend version"),
						app.Dd().Text("main"),
					),
				Footer: "Copyright © 2021 Felix Pojtinger and contributors (SPDX-License-Identifier: AGPL-3.0)",
			},
			&Modal{
				Open: c.settingsDialogOpen,
				Close: func() {
					c.dispatch(func() {
						c.settingsDialogOpen = false
					})
				},

				ID: "settings-modal-title",

				Title: "Settings",
				Body: []app.UI{
					&SettingsForm{
						NodeScanTimeout: c.nodeScanTimeout,
						SetNodeScanTimeout: func(i int64) {
							c.dispatch(func() {
								c.nodeScanTimeout = i
							})
						},

						PortScanTimeout: c.portScanTimeout,
						SetPortScanTimeout: func(i int64) {
							c.dispatch(func() {
								c.portScanTimeout = i
							})
						},

						NodeWakeTimeout: c.nodeWakeTimeout,
						SetNodeWakeTimeout: func(i int64) {
							c.dispatch(func() {
								c.nodeWakeTimeout = i
							})
						},

						Submit: func() {
							c.dispatch(func() {
								c.settingsDialogOpen = false
							})
						},
					},
				},
				Footer: []app.UI{
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
				},
			},
			&Modal{
				Open: c.metadataDialogOpen,
				Close: func() {
					c.dispatch(func() {
						c.metadataDialogOpen = false
					})
				},

				ID: "metadata-modal-title",

				Title: "Metadata",
				Body: []app.UI{
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
				},
			},
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
