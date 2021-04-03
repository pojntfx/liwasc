package components

import (
	"fmt"
	"net/url"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type Property struct {
	app.Compo

	Key   string
	Icon  string
	Value string
	Link  bool
}

func (c *Property) Render() app.UI {
	return app.Div().
		Class("pf-c-description-list__group").
		Body(
			app.Dt().
				Class("pf-c-description-list__term").
				Body(
					app.If(
						c.Icon == "",
						app.Span().
							Class("pf-c-description-list__text").
							Text(c.Key),
					).Else(
						app.Span().
							Class("pf-c-description-list__text").
							Body(
								app.I().Class(fmt.Sprintf("%v pf-u-mr-xs", c.Icon)).Aria("hidden", true),
								app.Text(c.Key),
							),
					),
				),
			app.Dd().
				Class("pf-c-description-list__description").
				Body(
					app.Div().
						Class("pf-c-description-list__text").
						Body(
							app.If(
								c.Value == "",
								app.Text("Unregistered"),
							).Else(
								app.If(
									c.Link,
									app.A().
										Target("_blank").
										Href(
											fmt.Sprintf(
												"https://duckduckgo.com/?q=%v",
												url.QueryEscape(c.Value),
											),
										).
										Text(c.Value),
								).Else(
									app.Text(c.Value),
								),
							),
						),
				),
		)
}
