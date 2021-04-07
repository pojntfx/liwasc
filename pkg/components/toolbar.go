package components

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type Toolbar struct {
	app.Compo

	NodeScanRunning        bool
	TriggerFullNetworkScan func()

	LastNodeScanDate string
	Subnets          []string
	Device           string

	ToggleMetadataDialogOpen func()
}

func (c *Toolbar) Render() app.UI {
	return app.Div().
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
										Loading: c.NodeScanRunning,
										Icon:    "fas fa-rocket",
										Text:    "Trigger Scan",

										OnClick: func(ctx app.Context, e app.Event) {
											c.TriggerFullNetworkScan()
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
																					app.Text(c.LastNodeScanDate),
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
														app.Range(c.Subnets).Slice(func(i int) app.UI {
															return app.Li().
																Class("pf-c-label-group__list-item").
																Body(
																	app.Span().
																		Class("pf-c-label").
																		Body(
																			app.Span().
																				Class("pf-c-label__content").
																				Body(
																					app.Text(c.Subnets[i]),
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
																				app.Text(c.Device),
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
										Loading: c.NodeScanRunning,
										Icon:    "fas fa-rocket",
										Text:    "Trigger Scan",

										OnClick: func(ctx app.Context, e app.Event) {
											c.TriggerFullNetworkScan()
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
											c.ToggleMetadataDialogOpen()
										}).
										Body(
											app.I().
												Class("fas fa-info-circle").
												Aria("hidden", true),
										),
								),
						),
				),
		)
}
