package components

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type OnOffSwitchComponent struct {
	app.Compo

	On       bool
	OnToggle func(ctx app.Context, e app.Event)
}

func (c *OnOffSwitchComponent) Render() app.UI {
	return app.Label().Class("pf-c-switch").Body(
		&OnOffSwitchCheckboxComponent{On: c.On, OnToggle: c.OnToggle},
		app.Span().Class("pf-c-switch__toggle").Body(
			app.Span().Class("pf-c-switch__toggle-icon").Body(
				app.I().Class("fas fa-lightbulb"),
			),
		),
		app.Span().Class("pf-c-switch__label pf-m-on").Body(app.Text("On")),
		app.Span().Class("pf-c-switch__label pf-m-off").Body(app.Text("Off")),
	)
}

type OnOffSwitchCheckboxComponent struct {
	app.Compo

	On       bool
	OnToggle func(ctx app.Context, e app.Event)
}

func (c *OnOffSwitchCheckboxComponent) Render() app.UI {
	app.Dispatch(func() {
		c.JSValue().Set("checked", c.On)
	})

	return app.Input().Class("pf-c-switch__input").Type("checkbox").Checked(c.On).OnChange(c.OnToggle)
}
