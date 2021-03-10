package experimental

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type ConfigProviderChildrenProps struct {
	BackendURL      string
	OIDCIssuer      string
	OIDCClientID    string
	OIDCRedirectURL string

	SetBackendURL,
	SetOIDCIssuer,
	SetOIDCClientID,
	SetOIDCRedirectURL func(string)
	ApplyConfig func()

	Error error
}

type ConfigProviderComponent struct {
	app.Compo

	Children func(ConfigProviderChildrenProps) app.UI

	backendURL      string
	oidcIssuer      string
	oidcClientID    string
	oidcRedirectURL string

	err error
}

func (c *ConfigProviderComponent) Render() app.UI {
	return c.Children(ConfigProviderChildrenProps{
		BackendURL:      c.backendURL,
		OIDCIssuer:      c.oidcIssuer,
		OIDCClientID:    c.oidcClientID,
		OIDCRedirectURL: c.oidcRedirectURL,

		SetBackendURL: func(s string) {
			c.dispatch(func() {
				c.backendURL = s
			})
		},
		SetOIDCIssuer: func(s string) {
			c.dispatch(func() {
				c.oidcIssuer = s
			})
		},
		SetOIDCClientID: func(s string) {
			c.dispatch(func() {
				c.oidcClientID = s
			})
		},
		SetOIDCRedirectURL: func(s string) {
			c.dispatch(func() {
				c.oidcRedirectURL = s
			})
		},
		ApplyConfig: func() {
			c.validate()
		},

		Error: c.err,
	})
}

func (c *ConfigProviderComponent) invalidate(err error) {
	// Set the error state
	c.err = err

	c.Update()
}

func (c *ConfigProviderComponent) dispatch(action func()) {
	action()

	c.Update()
}

func (c *ConfigProviderComponent) validate() {
	// Validate fields
	if c.oidcIssuer == "" {
		c.invalidate(errors.New("invalid OIDC issuer"))

		return
	}

	if c.oidcClientID == "" {
		c.invalidate(errors.New("invalid OIDC client ID"))

		return
	}

	if _, err := url.ParseRequestURI(c.backendURL); err != nil {
		c.invalidate(fmt.Errorf("invalid backend URL: %v", err))

		return
	}

	if _, err := url.ParseRequestURI(c.oidcRedirectURL); err != nil {
		c.invalidate(fmt.Errorf("invalid OIDC redirect URL: %v", err))

		return
	}

	// If all are valid, set ready state
	c.dispatch(func() {
		c.err = nil
	})
}
