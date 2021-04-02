package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/pojntfx/liwasc/pkg/providers"
)

type PortList struct {
	Ports []providers.Port

	expanded bool

	app.Compo
}

func (c *PortList) Render() app.UI {
	portsToDisplay := c.Ports
	if len(c.Ports) >= 3 && !c.expanded {
		portsToDisplay = c.Ports[:3]
	}

	return app.Div().
		Class("pf-c-label-group").
		Body(
			app.Div().
				Class("pf-c-label-group__main").
				Body(
					app.Ul().
						Class("pf-c-label-group__list").
						Aria("role", "list").
						Aria("label", "Ports of node").
						Body(
							app.Range(portsToDisplay).Slice(func(j int) app.UI {
								return app.Li().
									Class("pf-c-label-group__list-item").
									Body(
										app.Span().
											Class("pf-c-label").
											Body(
												app.
													Span().
													Class("pf-c-label__content").
													Text(
														fmt.Sprintf(
															"%v/%v (%v)",
															portsToDisplay[j].PortNumber,
															portsToDisplay[j].TransportProtocol,
															func() string {
																service := portsToDisplay[j].ServiceName
																if service == "" {
																	service = "Unregistered"
																}

																return service
															}(),
														),
													),
											),
									)
							}),
							app.If(
								// Only collapse if there are more than five ports
								len(c.Ports) >= 3,
								app.Li().
									Class("pf-c-label-group__list-item").
									Body(
										app.Button().
											Class("pf-c-label pf-m-overflow").
											OnClick(func(ctx app.Context, e app.Event) {
												e.Call("stopPropagation")

												c.dispatch(func() {
													c.expanded = !c.expanded
												})
											}).
											Body(
												app.Span().
													Class("pf-c-label__content").
													Body(
														app.If(
															c.expanded,
															app.Text(
																fmt.Sprintf("%v less", len(c.Ports)-3),
															),
														).Else(
															app.Text(
																fmt.Sprintf("%v more", len(c.Ports)-3),
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

func (c *PortList) dispatch(action func()) {
	action()

	c.Update()
}
