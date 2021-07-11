package oauth

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type Endpoint struct {
	AuthURL string
}

var VKEndpoint = Endpoint{
	AuthURL: "https://oauth.vk.com/authorize",
}

type Config interface {
	AuthURL(state string) string
	SetToken(token string)
	Token() string
	SetUser(userId string)
	User() string
}

func NewConfig(clientId string, clientSecret string, redirectUrl string, scopes []string, ep Endpoint) Config {
	return &config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  redirectUrl,
		Scopes:       scopes,
		Endpoint:     ep,
	}
}

type config struct {
	// ClientID is the application's ID.
	ClientID string

	// ClientSecret is the application's secret.
	ClientSecret string

	// Endpoint contains the resource server's token endpoint
	// URLs. These are constants specific to each server and are
	// often available via site-specific packages, such as
	// google.Endpoint or github.Endpoint.
	Endpoint Endpoint

	// RedirectURL is the URL to redirect users going through
	// the OAuth flow, after the resource owner's URLs.
	RedirectURL string

	// Scope specifies optional requested permissions.
	Scopes []string

	//
	token string

	//
	userId string
}

func (c config) AuthURL(state string) string {
	var scopeList string
	for i, scope := range c.Scopes {
		scopeList = scopeList + scope
		if i != len(c.Scopes)-1 {
			scopeList = scopeList + ","
		}
	}

	url := fmt.Sprintf(
		"%s?client_id=%s&display=page&redirect_uri=%s&scope=%s&response_type=token&v=5.131&state=%s",
		c.Endpoint.AuthURL, c.ClientID, c.RedirectURL, scopeList, state,
	)
	log.Info().Msgf("Auth url: %s", url)
	return url
}

func (c *config) SetToken(token string) {
	c.token = token
}

func (c config) Token() string {
	return c.token
}

func (c *config) SetUser(userId string) {
	c.userId = userId
}

func (c config) User() string {
	return c.userId
}
