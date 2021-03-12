package components

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"golang.org/x/oauth2"
)

const (
	oauth2TokenKey             = "oauth2Token"
	idTokenKey                 = "idToken"
	userInfoKey                = "userInfo"
	AUTHORIZATION_METADATA_KEY = "X-Liwasc-Authorization"
)

type OIDCLoginProviderChildrenProps struct {
	OAuth2Token oauth2.Token
	IDToken     string
	UserInfo    oidc.UserInfo
	Error       error

	Logout func()
}

type OIDCLoginProviderComponent struct {
	app.Compo

	Issuer      string
	ClientID    string
	RedirectURL string
	HomePath    string
	Scopes      []string

	LocalStoragePrefix string

	Children func(OIDCLoginProviderChildrenProps) app.UI

	oauth2Token oauth2.Token
	idToken     string
	userInfo    oidc.UserInfo
	err         error
}

func (c *OIDCLoginProviderComponent) Render() app.UI {
	c.upsertLogin()

	return c.Children(
		OIDCLoginProviderChildrenProps{
			OAuth2Token: c.oauth2Token,
			IDToken:     c.idToken,
			UserInfo:    c.userInfo,
			Error:       c.err,
			Logout:      func() { c.handleLogout(true) },
		},
	)
}

func (c *OIDCLoginProviderComponent) upsertLogin() {
	// Fetch current info from local storage
	oauth2Token, idToken, userInfo, err := c.getStateFromLocalStorage()
	if err != nil {
		c.invalidateLogin(err)

		return
	}

	// Create OIDC client
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, c.Issuer)
	if err != nil {
		c.invalidateLogin(err)

		return
	}

	config := oauth2.Config{
		ClientID:    c.ClientID,
		Endpoint:    provider.Endpoint(),
		RedirectURL: c.RedirectURL,
		Scopes:      append([]string{oidc.ScopeOpenID}, c.Scopes...),
	}

	if oauth2Token.AccessToken == "" || userInfo.Email == "" {
		if app.Window().URL().Query().Get("state") == "" {
			// If info could not be found in both local storage and the URL, login
			app.Navigate(config.AuthCodeURL(c.RedirectURL, oauth2.AccessTypeOffline))

			return
		}

		// Info could not be found in local storage but could be found in the URL, set local storage and navigate to home
		oauth2Token, err := config.Exchange(ctx, app.Window().URL().Query().Get("code"))
		if err != nil {
			c.invalidateLogin(err)

			return
		}

		idToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			c.invalidateLogin(err)

			return
		}

		userInfo, err := provider.UserInfo(ctx, oauth2.StaticTokenSource(oauth2Token))
		if err != nil {
			c.invalidateLogin(err)

			return
		}

		if err := c.setStateToLocalStorage(*oauth2Token, idToken, *userInfo); err != nil {
			c.invalidateLogin(err)

			return
		}

		_, _, _, err = c.getStateFromLocalStorage()
		if err != nil {
			c.invalidateLogin(err)

			return
		}

		c.Update()

		app.Navigate(c.HomePath)

		return
	}

	// Info could be found in local storage; set in state and update
	c.oauth2Token = oauth2Token
	c.idToken = idToken
	c.userInfo = userInfo

	c.registerTokenRefresh()

	return
}

func (c *OIDCLoginProviderComponent) registerTokenRefresh() {
	go func() {
		for {
			// Wait till token expires
			if c.oauth2Token.Expiry.After(time.Now()) {
				time.Sleep(c.oauth2Token.Expiry.Sub(time.Now()))
			}

			// Fetch new token
			tokenSource := oauth2.StaticTokenSource(&c.oauth2Token)

			newOAuth2Token, err := tokenSource.Token()
			if err != nil {
				c.invalidateLogin(err)

				return
			}

			newIDToken, ok := newOAuth2Token.Extra("id_token").(string)
			if !ok {
				c.invalidateLogin(err)

				return
			}

			// Set new token in local storage
			if err := c.setStateToLocalStorage(*newOAuth2Token, newIDToken, c.userInfo); err != nil {
				c.invalidateLogin(err)

				return
			}

			// Set new token in state
			c.oauth2Token = *newOAuth2Token
			c.idToken = newIDToken

			c.Update()
		}
	}()
}

func (c *OIDCLoginProviderComponent) handleLogout(withRedirect bool) {
	app.LocalStorage.Del(c.getKeyWithPrefix(oauth2TokenKey))
	app.LocalStorage.Del(c.getKeyWithPrefix(userInfoKey))

	c.Update()

	if withRedirect {
		app.Navigate(c.HomePath)
	}
}

func (c *OIDCLoginProviderComponent) getStateFromLocalStorage() (oauth2.Token, string, oidc.UserInfo, error) {
	oauth2Token := oauth2.Token{}
	idToken := ""
	userInfo := oidc.UserInfo{}
	if err := app.LocalStorage.Get(c.getKeyWithPrefix(oauth2TokenKey), &oauth2Token); err != nil {
		return oauth2.Token{}, "", oidc.UserInfo{}, err
	}
	if err := app.LocalStorage.Get(c.getKeyWithPrefix(idTokenKey), &idToken); err != nil {
		return oauth2.Token{}, "", oidc.UserInfo{}, err
	}
	if err := app.LocalStorage.Get(c.getKeyWithPrefix(userInfoKey), &userInfo); err != nil {
		return oauth2.Token{}, "", oidc.UserInfo{}, err
	}

	return oauth2Token, idToken, userInfo, nil
}

func (c *OIDCLoginProviderComponent) setStateToLocalStorage(oauth2Token oauth2.Token, idToken string, userInfo oidc.UserInfo) error {
	if err := app.LocalStorage.Set(c.getKeyWithPrefix(oauth2TokenKey), oauth2Token); err != nil {
		return err
	}

	if err := app.LocalStorage.Set(c.getKeyWithPrefix(idTokenKey), idToken); err != nil {
		return err
	}

	return app.LocalStorage.Set(c.getKeyWithPrefix(userInfoKey), userInfo)
}

func (c *OIDCLoginProviderComponent) invalidateLogin(err error) {
	log.Println("could not login", err)

	c.err = err

	c.handleLogout(false)
}

func (c *OIDCLoginProviderComponent) getKeyWithPrefix(key string) string {
	return fmt.Sprintf("%v.%v", c.LocalStoragePrefix, key)
}
