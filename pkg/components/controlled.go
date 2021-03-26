package components

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type Controlled struct {
	app.Compo

	Component  app.UI
	Properties map[string]interface{}
}

func (c *Controlled) Render() app.UI {
	for key, value := range c.Properties {
		app.Dispatch(func() {
			if c.JSValue() != nil {
				c.JSValue().Set(key, value)
			}
		})
	}

	return c.Component
}
