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
	Ready           bool

	SetBackendURL,
	SetOIDCIssuer,
	SetOIDCClientID,
	SetOIDCRedirectURL func(string)
	ApplyConfig func()

	Error error
}

type ConfigProviderComponent struct {
	app.Compo

	StoragePrefix string
	Children      func(ConfigProviderChildrenProps) app.UI

	backendURL      string
	oidcIssuer      string
	oidcClientID    string
	oidcRedirectURL string
	ready           bool

	err error
}

const (
	backendURLKey      = "backendURL"
	oidcIssuerKey      = "oidcIssuer"
	oidcClientIDKey    = "oidcClientID"
	oidcRedirectURLKey = "oidcRedirectURL"
)

func (c *ConfigProviderComponent) Render() app.UI {
	return c.Children(ConfigProviderChildrenProps{
		BackendURL:      c.backendURL,
		OIDCIssuer:      c.oidcIssuer,
		OIDCClientID:    c.oidcClientID,
		OIDCRedirectURL: c.oidcRedirectURL,
		Ready:           c.ready,

		SetBackendURL: func(s string) {
			c.dispatch(func() {
				c.ready = false
				c.backendURL = s
			})
		},
		SetOIDCIssuer: func(s string) {
			c.dispatch(func() {
				c.ready = false
				c.oidcIssuer = s
			})
		},
		SetOIDCClientID: func(s string) {
			c.dispatch(func() {
				c.ready = false
				c.oidcClientID = s
			})
		},
		SetOIDCRedirectURL: func(s string) {
			c.dispatch(func() {
				c.ready = false
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
	c.ready = false

	c.Update()
}

func (c *ConfigProviderComponent) dispatch(action func()) {
	action()

	c.Update()
}

func (c *ConfigProviderComponent) validate() {
	// Validate fields
	if c.oidcClientID == "" {
		c.invalidate(errors.New("invalid OIDC client ID"))

		return
	}

	if _, err := url.ParseRequestURI(c.oidcIssuer); err != nil {
		c.invalidate(fmt.Errorf("invalid OIDC issuer: %v", err))

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

	// Persist state
	if err := c.persist(); err != nil {
		c.invalidate(err)

		return
	}

	// If all are valid, set ready state
	c.dispatch(func() {
		c.err = nil
		c.ready = true
	})
}

func (c *ConfigProviderComponent) persist() error {
	// Write state to storage
	if err := app.LocalStorage.Set(c.getKey(backendURLKey), c.backendURL); err != nil {
		return err
	}
	if err := app.LocalStorage.Set(c.getKey(oidcIssuerKey), c.oidcIssuer); err != nil {
		return err
	}
	if err := app.LocalStorage.Set(c.getKey(oidcClientIDKey), c.oidcClientID); err != nil {
		return err
	}

	return app.LocalStorage.Set(c.getKey(oidcRedirectURLKey), c.oidcRedirectURL)
}

func (c *ConfigProviderComponent) rehydrateFromURL() bool {
	// Read the values from the URL
	query := app.Window().URL().Query()

	backendURL := query.Get(backendURLKey)
	oidcIssuer := query.Get(oidcIssuerKey)
	oidcClientID := query.Get(oidcClientIDKey)
	oidcRedirectURL := query.Get(oidcRedirectURLKey)

	// If all values are set, set them in the data provider
	if backendURL != "" && oidcIssuer != "" && oidcClientID != "" && oidcRedirectURL != "" {
		c.dispatch(func() {
			c.backendURL = backendURL
			c.oidcIssuer = oidcIssuer
			c.oidcClientID = oidcClientID
			c.oidcRedirectURL = oidcRedirectURL
		})

		return true
	}

	return false
}

func (c *ConfigProviderComponent) getKey(key string) string {
	// Get a prefixed key
	return fmt.Sprintf("%v.%v", c.StoragePrefix, key)
}

func (c *ConfigProviderComponent) OnMount(context app.Context) {
	// Initialize state
	c.backendURL = ""
	c.oidcIssuer = ""
	c.oidcClientID = ""
	c.oidcRedirectURL = ""
	c.ready = false

	// Rehydrate from URL
	rehydratedFromURL := c.rehydrateFromURL()

	// If rehydrated from URL, validate & apply
	if rehydratedFromURL {
		// Auto-apply if configured
		c.validate()
	}
}
