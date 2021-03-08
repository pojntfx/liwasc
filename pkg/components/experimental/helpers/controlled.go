package helpers

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type Controlled struct {
	app.Compo

	Component app.UI
	Value     interface{}
}

func (c *Controlled) Render() app.UI {
	app.Dispatch(func() {
		c.JSValue().Set("value", c.Value)
	})

	return c.Component
}
