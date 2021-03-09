package experimental

import (
	"github.com/coreos/go-oidc"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type LoginProviderChildrenProps struct {
	IDToken  string
	UserInfo oidc.UserInfo

	Logout func()

	Error   error
	Recover func()
}

type LoginProviderComponent struct {
	app.Compo

	Issuer      string
	ClientID    string
	RedirectURL string
	Children    func(LoginProviderChildrenProps) app.UI

	idToken  string
	userInfo oidc.UserInfo

	err error
}

func (c *LoginProviderComponent) Render() app.UI {
	return c.Children(LoginProviderChildrenProps{
		IDToken:  c.idToken,
		UserInfo: c.userInfo,

		Logout: c.logout,

		Error:   c.err,
		Recover: c.recover,
	})
}

func (c *LoginProviderComponent) logout() {
	// TODO: Implement
}

func (c *LoginProviderComponent) recover() {
	// TODO: Implement
}

func (c *LoginProviderComponent) dispatch(action func()) {
	action()

	c.Update()
}

func (c *LoginProviderComponent) OnMount(context app.Context) {
	// Initialize state
	c.dispatch(func() {
		c.idToken = ""
		c.userInfo = oidc.UserInfo{}
	})
}
