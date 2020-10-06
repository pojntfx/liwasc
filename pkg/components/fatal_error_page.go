package components

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type FatalErrorPageComponent struct {
	app.Compo

	Header             string
	Description        string
	StackTraceLanguage string
	StackTraceContent  string

	Actions []app.UI
}

func (c *FatalErrorPageComponent) Render() app.UI {
	return app.Div().Class("x__page--full").Body(
		app.Div().Class("pf-c-alert pf-m-danger").Body(
			app.Div().Class("pf-c-alert__icon").Body(
				app.I().Class("fas fa-fw fa-exclamation-circle"),
			),
			app.P().Class("pf-c-alert__title").Body(
				app.Strong().Text(c.Header),
			),
			app.Div().Class("pf-c-alert__description").Body(
				app.P().Class("pf-u-mb-sm").Text(c.Description),
				app.Div().Class("pf-c-code-editor pf-m-read-only").Body(
					app.Div().Class("pf-c-code-editor__header").Body(
						app.Div().Class("pf-c-code-editor__tab").Body(
							app.Span().Class("pf-c-code-editor__tab-icon").Body(
								app.I().Class("fas fa-code"),
							),
							app.Span().Class("pf-c-code-editor__tab-text").Text(c.StackTraceLanguage),
						),
					),
					app.Div().Class("pf-c-code-editor__main").Body(
						app.Div().Class("pf-c-code-editor__code").Body(
							app.Pre().Class("pf-c-code-editor__code-pre").Text(c.StackTraceContent),
						),
					),
				),
			),
			app.Div().Class("pf-c-alert__action-group").Body(
				c.Actions...,
			),
		),
	)
}
