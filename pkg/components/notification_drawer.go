package components

import "github.com/maxence-charriere/go-app/v8/pkg/app"

type Notification struct {
	Message string
	Time    string
}

type NotificationDrawer struct {
	app.Compo

	Notifications []Notification
}

func (c *NotificationDrawer) Render() app.UI {
	return app.Div().
		Class("pf-c-notification-drawer").
		Body(
			app.Div().
				Class("pf-c-notification-drawer__header").
				Body(
					app.H1().
						Class("pf-c-notification-drawer__header-title").
						Text("Events"),
				),
			app.Div().Class("pf-c-notification-drawer__body").Body(
				app.Ul().Class("pf-c-notification-drawer__list").Body(
					app.Range(c.Notifications).Slice(func(i int) app.UI {
						return app.Li().Class("pf-c-notification-drawer__list-item pf-m-read pf-m-info").Body(
							app.Div().Class("pf-c-notification-drawer__list-item-description").Text(
								c.Notifications[len(c.Notifications)-1-i].Message,
							),
							app.Div().Class("pf-c-notification-drawer__list-item-timestamp").Text(
								c.Notifications[len(c.Notifications)-1-i].Time,
							),
						)
					}),
				),
			),
		)
}
