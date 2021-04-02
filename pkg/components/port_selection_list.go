package components

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/pojntfx/liwasc/pkg/providers"
)

type PortSelectionList struct {
	app.Compo

	Ports           []providers.Port
	SelectedPort    string
	SetSelectedPort func(string)
}

func (c *PortSelectionList) Render() app.UI {
	return app.Ul().
		Class("pf-c-data-list pf-u-my-lg").
		ID("ports-in-inspector").
		Aria("role", "list").
		Aria("label", "Ports").
		Body(
			app.Range(c.Ports).Slice(func(i int) app.UI {
				return app.Li().
					Class(func() string {
						classes := "pf-c-data-list__item pf-m-selectable"

						if c.SelectedPort == GetPortID(c.Ports[i]) {
							classes += " pf-m-selected"
						}

						return classes
					}()).
					Aria("labelledby", "ports-in-inspector").
					TabIndex(0).
					OnClick(func(ctx app.Context, e app.Event) {
						// Reset selected port
						if c.SelectedPort == GetPortID(c.Ports[i]) {
							c.SetSelectedPort("")

							return
						}

						// Set selected port
						c.SetSelectedPort(GetPortID(c.Ports[i]))
					}).
					Body(
						app.Div().Class("pf-c-data-list__item-row").Body(
							app.Div().Class("pf-c-data-list__item-content").Body(
								app.Div().Class("pf-c-data-list__cell").Text(GetPortID(c.Ports[i])),
							),
						),
					)
			}),
		)
}
