package experimental

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"golang.org/x/oauth2"
)

const (
	stateQueryParameter = "state"
	codeQueryParameter  = "code"
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
	HomeURL     string
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

func (c *LoginProviderComponent) OnMount(ctx app.Context) {
	// Initialize state
	c.dispatch(func() {
		c.idToken = ""
		c.userInfo = oidc.UserInfo{}
	})

	// Create the provider
	provider, err := oidc.NewProvider(context.Background(), c.Issuer)
	if err != nil {
		panic(err)
	}

	// Create the OAuth2 config
	config := &oauth2.Config{
		ClientID:    c.ClientID,
		RedirectURL: c.RedirectURL,
		Endpoint:    provider.Endpoint(),
		Scopes:      []string{oidc.ScopeOpenID, "profile", "email"},
	}

	// If state query parameter not in query, redirect
	if app.Window().URL().Query().Get(stateQueryParameter) == "" {
		app.Navigate(config.AuthCodeURL(stateQueryParameter, oauth2.AccessTypeOffline))

		return
	}

	// Parse token
	token, err := config.Exchange(context.Background(), app.Window().URL().Query().Get(codeQueryParameter))
	if err != nil {
		panic(err)
	}

	// Parse user info
	userInfo, err := provider.UserInfo(context.Background(), oauth2.StaticTokenSource(token))
	if err != nil {
		panic(err)
	}

	// Set the login state
	c.dispatch(func() {
		c.idToken = token.Extra("id_token").(string)
		c.userInfo = *userInfo
	})

	// Navigate to home URL
	app.Navigate(c.HomeURL)
}
