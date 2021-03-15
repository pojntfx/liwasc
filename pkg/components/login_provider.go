package components

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"golang.org/x/oauth2"
)

const (
	oauth2TokenKey = "oauth2Token"
	idTokenKey     = "idToken"
	userInfoKey    = "userInfo"

	stateQueryParameter = "state"
	codeQueryParameter  = "code"

	idTokenExtraKey = "id_token"
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

	Issuer        string
	ClientID      string
	RedirectURL   string
	HomeURL       string
	Scopes        []string
	StoragePrefix string
	Children      func(LoginProviderChildrenProps) app.UI

	oauth2Token oauth2.Token
	idToken     string
	userInfo    oidc.UserInfo

	err error
}

func (c *LoginProviderComponent) Render() app.UI {
	c.authorize()

	return c.Children(
		LoginProviderChildrenProps{
			IDToken:  c.idToken,
			UserInfo: c.userInfo,

			Logout: func() {
				c.logout(true)
			},

			Error:   c.err,
			Recover: c.recover,
		},
	)
}

func (c *LoginProviderComponent) panic(err error) {
	c.dispatch(func() {
		// Set the error
		c.err = err
	})

	// Prevent infinite retries
	time.Sleep(time.Second)
}

func (c *LoginProviderComponent) recover() {
	c.dispatch(func() {
		// Clear the error
		c.err = nil
	})

	// Logout
	c.logout(false)
}

func (c *LoginProviderComponent) dispatch(action func()) {
	action()

	c.Update()
}

func (c *LoginProviderComponent) watch() {
	for {
		// Wait till token expires
		if c.oauth2Token.Expiry.After(time.Now()) {
			time.Sleep(c.oauth2Token.Expiry.Sub(time.Now()))
		}

		// Fetch new OAuth2 token
		oauth2Token, err := oauth2.StaticTokenSource(&c.oauth2Token).Token()
		if err != nil {
			c.panic(err)

			return
		}

		// Parse ID token
		idToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			c.panic(err)

			return
		}

		// Persist state in storage
		if err := c.persist(*oauth2Token, idToken, c.userInfo); err != nil {
			c.panic(err)

			return
		}

		// Set the login state
		c.dispatch(func() {
			c.oauth2Token = *oauth2Token
			c.idToken = idToken
		})
	}
}

func (c *LoginProviderComponent) logout(withRedirect bool) {
	// Remove from storage
	c.clear()

	// Reload the app
	if withRedirect {
		app.Reload()
	}
}

func (c *LoginProviderComponent) rehydrate() (oauth2.Token, string, oidc.UserInfo, error) {
	// Read state from storage
	oauth2Token := oauth2.Token{}
	idToken := ""
	userInfo := oidc.UserInfo{}

	if err := app.LocalStorage.Get(c.getKey(oauth2TokenKey), &oauth2Token); err != nil {
		return oauth2.Token{}, "", oidc.UserInfo{}, err
	}
	if err := app.LocalStorage.Get(c.getKey(idTokenKey), &idToken); err != nil {
		return oauth2.Token{}, "", oidc.UserInfo{}, err
	}
	if err := app.LocalStorage.Get(c.getKey(userInfoKey), &userInfo); err != nil {
		return oauth2.Token{}, "", oidc.UserInfo{}, err
	}

	return oauth2Token, idToken, userInfo, nil
}

func (c *LoginProviderComponent) persist(oauth2Token oauth2.Token, idToken string, userInfo oidc.UserInfo) error {
	// Write state to storage
	if err := app.LocalStorage.Set(c.getKey(oauth2TokenKey), oauth2Token); err != nil {
		return err
	}
	if err := app.LocalStorage.Set(c.getKey(idTokenKey), idToken); err != nil {
		return err
	}
	return app.LocalStorage.Set(c.getKey(userInfoKey), userInfo)
}

func (c *LoginProviderComponent) clear() {
	// Remove from storage
	app.LocalStorage.Del(c.getKey(oauth2TokenKey))
	app.LocalStorage.Del(c.getKey(idTokenKey))
	app.LocalStorage.Del(c.getKey(userInfoKey))

	// Remove cookies
	app.Window().Get("document").Set("cookie", "")
}

func (c *LoginProviderComponent) getKey(key string) string {
	// Get a prefixed key
	return fmt.Sprintf("%v.%v", c.StoragePrefix, key)
}

func (c *LoginProviderComponent) authorize() {
	// Read state from storage
	oauth2Token, idToken, userInfo, err := c.rehydrate()
	if err != nil {
		c.panic(err)

		return
	}

	// Create the OIDC provider
	provider, err := oidc.NewProvider(context.Background(), c.Issuer)
	if err != nil {
		c.panic(err)

		return
	}

	// Create the OAuth2 config
	config := &oauth2.Config{
		ClientID:    c.ClientID,
		RedirectURL: c.RedirectURL,
		Endpoint:    provider.Endpoint(),
		Scopes:      append([]string{oidc.ScopeOpenID}, c.Scopes...),
	}

	// Log in
	if oauth2Token.AccessToken == "" || userInfo.Email == "" {
		// Logged out state, info neither in storage nor in URL: Redirect to login
		if app.Window().URL().Query().Get(stateQueryParameter) == "" {
			app.Navigate(config.AuthCodeURL(c.RedirectURL, oauth2.AccessTypeOffline))

			return
		}

		// Intermediate state, info is in URL: Parse OAuth2 token
		oauth2Token, err := config.Exchange(context.Background(), app.Window().URL().Query().Get(codeQueryParameter))
		if err != nil {
			c.panic(err)

			return
		}

		// Parse ID token
		idToken, ok := oauth2Token.Extra(idTokenExtraKey).(string)
		if !ok {
			c.panic(err)

			return
		}

		// Parse user info
		userInfo, err := provider.UserInfo(context.Background(), oauth2.StaticTokenSource(oauth2Token))
		if err != nil {
			c.panic(err)

			return
		}

		// Persist state in storage
		if err := c.persist(*oauth2Token, idToken, *userInfo); err != nil {
			c.panic(err)

			return
		}

		// Test validity of storage
		if _, _, _, err = c.rehydrate(); err != nil {
			c.panic(err)

			return
		}

		// Update and navigate to home URL
		c.Update()
		app.Navigate(c.HomeURL)

		return
	}

	// Validation state

	// Create the OIDC config
	oidcConfig := &oidc.Config{
		ClientID: c.ClientID,
	}

	// Create the OIDC verifier and validate the token (i.e. check for it's expiry date)
	verifier := provider.Verifier(oidcConfig)
	if _, err := verifier.Verify(context.Background(), idToken); err != nil {
		// Invalid token; clear and re-authorize
		c.clear()
		c.authorize()

		return
	}

	// Logged in state

	// Set the login state
	c.oauth2Token = oauth2Token
	c.idToken = idToken
	c.userInfo = userInfo

	// Watch and renew token once expired
	go c.watch()

	return
}
