package components

import (
	"fmt"
	"strings"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/pojntfx/liwasc/pkg/providers"
)

type Inspector struct {
	app.Compo

	Open               bool
	Close              func()
	StartNodeWake      func()
	TriggerNetworkScan func()

	Header          []app.UI
	Body            app.UI
	Node            providers.Node
	PortFilter      string
	SetPortFilter   func(string)
	SelectedPort    string
	SetSelectedPort func(string)

	portsAndServicesOpen bool
	detailsOpen          bool
}

func (c *Inspector) Render() app.UI {
	// Filter ports with port filter
	filteredPorts := append([]providers.Port{}, c.Node.Ports...) // Copy to prevent changes from parent
	if c.PortFilter != "" {
		filteredPorts = []providers.Port{}

		for _, port := range c.Node.Ports {
			if strings.Contains(GetPortID(port), c.PortFilter) {
				filteredPorts = append(filteredPorts, port)
			}
		}
	}

	// Get the selected port
	selectedPort := providers.Port{
		PortNumber: -1,
	}
	for _, port := range filteredPorts {
		if GetPortID(port) == c.SelectedPort {
			selectedPort = port

			break
		}
	}

	// Reset selected port if it does not exist anymore
	if selectedPort.PortNumber == -1 {
		c.SelectedPort = ""
	}

	return app.Div().
		Class(func() string {
			classes := "pf-c-drawer pf-m-inline-on-2xl"

			if c.Open {
				classes += " pf-m-expanded"
			}

			return classes
		}()).
		Body(
			app.Div().Class("pf-c-drawer__section").Body(
				append(
					c.Header,
					app.Div().
						Class("pf-c-divider").
						Aria("role", "separator"),
				)...,
			),
			app.Div().Class("pf-c-drawer__main").Body(
				// Content
				app.Div().
					Class("pf-c-drawer__content pf-m-no-background pf-x-m-overflow-x-hidden").
					Body(
						app.Div().
							Class("pf-c-drawer__body pf-m-padding").
							Body(c.Body),
					),
				// Panel
				app.Div().Class("pf-c-drawer__panel").Body(
					app.Div().
						Class("pf-c-drawer__body").
						Body(
							app.Div().
								Class("pf-c-drawer__head").
								Body(
									app.If(
										c.SelectedPort == "",
										app.Span().
											Text(fmt.Sprintf("Node %v", c.Node.MACAddress)),
									).Else(
										app.Div().
											Body(
												app.Button().
													Class("pf-c-button pf-m-plain pf-u-mr-md").
													Type("button").
													Aria("label", "Close port inspector").
													OnClick(func(ctx app.Context, e app.Event) {
														c.SetSelectedPort("")
													}).Body(
													app.I().Class("fas fa-arrow-left").Aria("hidden", true),
												),
												app.Span().
													Text(fmt.Sprintf("Port %v", GetPortID(selectedPort))),
											),
									),
									app.Div().
										Class("pf-c-drawer__actions").
										Body(
											app.If(
												c.SelectedPort == "",
												app.Label().
													Class("pf-c-switch pf-x-c-tooltip-wrapper pf-x-c-power-switch pf-u-mr-md").
													For("selected-node-inspector-power").
													Body(
														app.If(
															c.Node.PoweredOn,
															app.Div().
																Class("pf-c-tooltip pf-x-c-tooltip pf-x-c-tooltip--bottom pf-m-bottom").
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
																ID("selected-node-inspector-power").
																Aria("label", "Node is off").
																Name("selected-node-inspector-power").
																Type("checkbox").
																Checked(c.Node.PoweredOn).
																Disabled(c.Node.PoweredOn).
																OnClick(func(ctx app.Context, e app.Event) {
																	e.Call("stopPropagation")

																	c.StartNodeWake()
																}),
															Properties: map[string]interface{}{
																"checked":  c.Node.PoweredOn,
																"disabled": c.Node.PoweredOn,
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
															ID("selected-node-inspector-power-on").
															Aria("hidden", true).
															Body(
																app.If(
																	c.Node.NodeWakeRunning,
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
															ID("selected-node-inspector-power-off").
															Aria("hidden", true).
															Body(
																app.If(
																	c.Node.NodeWakeRunning,
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
											app.Div().
												Class("pf-c-drawer__close").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "Close inspector").
														OnClick(func(ctx app.Context, e app.Event) {
															c.Close()
														}).Body(
														app.I().Class("fas fa-times").Aria("hidden", true),
													),
												),
										),
								),
						),
					app.Div().
						Class("pf-c-drawer__body").
						Body(
							app.If(
								c.SelectedPort == "",
								app.Dl().
									Class("pf-c-description-list pf-m-2-col").
									Body(
										&Property{
											Key:   "IP Address",
											Value: c.Node.IPAddress,
										},
										&Property{
											Key:   "Vendor",
											Value: c.Node.Vendor,
										},
									),
								&ExpandableSection{
									Open: c.portsAndServicesOpen,
									OnToggle: func() {
										c.Defer(func(_ app.Context) {
											c.portsAndServicesOpen = !c.portsAndServicesOpen

											c.Update()
										})
									},
									Title:       "Ports and Services",
									ClosedTitle: "Hide ports and services",
									OpenTitle:   "Show ports and services",
									Body: []app.UI{
										&ProgressButton{
											Loading:   c.Node.PortScanRunning,
											Icon:      "fas fa-sync",
											Text:      "Trigger Port Scan",
											Secondary: true,
											Classes:   "pf-u-w-100",

											OnClick: func(ctx app.Context, e app.Event) {
												e.Call("stopPropagation")

												c.TriggerNetworkScan()
											},
										},
										app.Div().
											Class("pf-c-input-group pf-u-mt-lg").
											Body(
												&Controlled{
													Component: app.
														Input().
														Type("search").
														Placeholder("Service name or port number").
														Class("pf-c-form-control").
														Aria("label", "Service name or port number").
														OnInput(func(ctx app.Context, e app.Event) {
															c.SetPortFilter(ctx.JSSrc.Get("value").String())
														}),
													Properties: map[string]interface{}{
														"value": c.PortFilter,
													},
												},
												app.Button().
													Class("pf-c-button pf-m-control").
													Type("button").
													Aria("label", "Search button for service name or port number").Body(
													app.I().
														Class("fas fa-search").
														Aria("hidden", true),
												),
											),
										app.If(
											len(filteredPorts) > 0,
											&PortSelectionList{
												Ports:           filteredPorts,
												SelectedPort:    c.SelectedPort,
												SetSelectedPort: c.SetSelectedPort,
											},
										).
											ElseIf(
												c.PortFilter != "",
												app.Div().Class("pf-u-mt-lg").Text("No open ports found for this filter."),
											).
											ElseIf(
												c.Node.PortScanRunning,
												app.Div().Class("pf-u-mt-lg").Text("No open ports found yet."),
											).
											Else(
												app.Div().Class("pf-u-mt-lg").Text("No open ports found."),
											),
									},
								},
								&ExpandableSection{
									Open: c.detailsOpen,
									OnToggle: func() {
										c.Defer(func(_ app.Context) {
											c.detailsOpen = !c.detailsOpen

											c.Update()
										})
									},
									Title:       "Details",
									ClosedTitle: "Hide details",
									OpenTitle:   "Show details",
									Body: []app.UI{
										app.Dl().
											Class("pf-c-description-list pf-m-2-col").
											Body(
												&Property{
													Key:   "Registry",
													Value: c.Node.Registry,
												},
												&Property{
													Key:   "Organization",
													Value: c.Node.Organization,
												},
												&Property{
													Key:   "Address",
													Value: c.Node.Address,
												},
												&Property{
													Key:   "Visible",
													Value: fmt.Sprintf("%v", c.Node.Visible),
												},
											),
									},
								},
							).Else(
								app.Dl().
									Class("pf-c-description-list pf-m-2-col").
									Body(
										&Property{
											Key:   "Description",
											Value: selectedPort.Description,
										},
										&Property{
											Key:   "Assignee",
											Value: selectedPort.Assignee,
										},
										&Property{
											Key:   "Contact",
											Value: selectedPort.Contact,
										},
										&Property{
											Key:   "Registration Date",
											Value: selectedPort.RegistrationDate,
										},
										&Property{
											Key:   "Modification Date",
											Value: selectedPort.ModificationDate,
										},
										&Property{
											Key:   "Reference",
											Value: selectedPort.Reference,
										},
										&Property{
											Key:   "Service Code",
											Value: selectedPort.ServiceCode,
										},
										&Property{
											Key:   "Unauthorized Use Reported",
											Value: selectedPort.UnauthorizedUseReported,
										},
										&Property{
											Key:   "Assignment Notes",
											Value: selectedPort.AssignmentNotes,
										},
									),
							),
						),
				),
			),
		)
}

func (c *Inspector) OnMount(ctx app.Context) {
	c.portsAndServicesOpen = true
}

func GetPortID(port providers.Port) string {
	return fmt.Sprintf(
		"%v/%v (%v)",
		port.PortNumber,
		port.TransportProtocol,
		func() string {
			service := port.ServiceName
			if service == "" {
				service = "Unregistered"
			}

			return service
		}(),
	)
}
