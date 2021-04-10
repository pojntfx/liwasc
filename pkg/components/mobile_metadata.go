package components

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type MobileMetadata struct {
	app.Compo

	LastNodeScanDate string
	Subnets          []string
	Device           string
}

func (c *MobileMetadata) Render() app.UI {
	return app.Dl().
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
																	app.Text(c.LastNodeScanDate),
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
																	app.Text(c.Device),
																),
														),
												),
										),
								),
						),
				),
		)
}
