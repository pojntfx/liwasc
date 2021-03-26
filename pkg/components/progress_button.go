package components

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type ProgressButton struct {
	app.Compo

	Loading bool
	Icon    string
	Text    string

	OnClick func(ctx app.Context, e app.Event)
}

func (c *ProgressButton) Render() app.UI {
	return app.
		Button().
		Class(func() string {
			classes := "pf-c-button pf-m-primary"

			if c.Loading {
				classes += " pf-m-progress pf-m-in-progress"
			}

			return classes
		}()).
		OnClick(c.OnClick).
		Body(
			app.If(c.Loading,
				app.Span().
					Class("pf-c-button__progress").
					Body(
						app.Span().
							Class("pf-c-spinner pf-m-md").
							Aria("role", "progressbar").
							Aria("valuetext", "Loading...").
							Body(
								app.Span().Class("pf-c-spinner__clipper"),
								app.Span().Class("pf-c-spinner__lead-ball"),
								app.Span().Class("pf-c-spinner__tail-ball"),
							),
					)).Else(
				app.Span().
					Class("pf-c-button__icon pf-m-start").
					Body(
						app.I().
							Class(c.Icon).
							Aria("hidden", true),
					),
			),
			app.Text(c.Text),
		)
}