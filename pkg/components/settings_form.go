package components

import (
	"strconv"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type SettingsForm struct {
	app.Compo

	NodeScanTimeout    int64
	SetNodeScanTimeout func(int64)

	PortScanTimeout    int64
	SetPortScanTimeout func(int64)

	NodeWakeTimeout    int64
	SetNodeWakeTimeout func(int64)

	Submit func()
}

const (
	// Names and IDs
	nodeScanTimeoutName    = "nodeScanTimeout"
	portScanTimeoutName    = "portScanTimeout"
	nodeScanMACAddressName = "nodeScanMACAddressTimeout"

	nodeWakeTimeoutName    = "nodeWakeTimeout"
	nodeWakeMACAddressName = "nodeWakeMACAddressTimeout"

	// Default values
	defaultNodeWakeTimeout = 600000
	defaultNodeScanTimeout = 500
	defaultPortScanTimeout = 10
)

func (c *SettingsForm) Render() app.UI {
	return app.Form().
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

							c.SetNodeScanTimeout(int64(v))
						}),
					Properties: map[string]interface{}{
						"value": c.NodeScanTimeout,
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

							c.SetPortScanTimeout(int64(v))
						}),
					Properties: map[string]interface{}{
						"value": c.PortScanTimeout,
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

							c.SetNodeWakeTimeout(int64(v))

							c.Update()
						}),
					Properties: map[string]interface{}{
						"value": c.NodeWakeTimeout,
					},
				},
				Required: true,
			},
		).OnSubmit(func(ctx app.Context, e app.Event) {
		e.PreventDefault()

		c.Submit()
	})
}
