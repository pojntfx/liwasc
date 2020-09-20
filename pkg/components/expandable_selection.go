package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type ExpandableSectionComponent struct {
	app.Compo

	Open     bool
	OnToggle func(ctx app.Context, e app.Event)
	Title    string
	Content  app.UI
}

func (c *ExpandableSectionComponent) Render() app.UI {
	return app.Div().Class(fmt.Sprintf("pf-c-expandable-section pf-u-mb-md %v", func() string {
		if c.Open {
			return "pf-m-expanded"
		}

		return ""
	}())).Body(
		app.Button().Class("pf-c-expandable-section__toggle").Body(
			app.Span().Class("pf-c-expandable-section__toggle-icon").Body(
				app.I().Class("fas fa-angle-right"),
			),
			app.Span().Class("pf-c-expandable-section__toggle-text").Body(
				app.Text(c.Title),
			),
		).OnClick(c.OnToggle),
		&ExpandableSectionComponentContent{Content: c.Content, Open: c.Open},
	)
}

type ExpandableSectionComponentContent struct {
	app.Compo

	Content app.UI
	Open    bool
}

func (c *ExpandableSectionComponentContent) Render() app.UI {
	app.Dispatch(func() {
		c.JSValue().Set("hidden", !c.Open)
	})

	return app.Div().Class("pf-c-expandable-section__content").Hidden(!c.Open).Body(
		c.Content,
	)
}
