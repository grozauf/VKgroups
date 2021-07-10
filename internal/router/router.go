package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grozauf/VKgroups/internal/groups"
	"golang.org/x/oauth2"

	"github.com/rs/zerolog/log"

	vkApi "github.com/go-vk-api/vk"
)

type Router interface {
	Root(c *gin.Context)
	Groups(c *gin.Context)
}

func NewRouter(conf *oauth2.Config) Router {
	return &router{conf: conf}
}

type router struct {
	conf *oauth2.Config
}

func (r router) Root(c *gin.Context) {
	url := r.conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.HTML(
		http.StatusOK,
		"/templates/index.html",
		gin.H{
			"authUrl": url,
		},
	)
}

func (r router) Groups(c *gin.Context) {
	ctx := context.Background()
	// получаем код от API VK из квери стринга
	authCode := c.Request.URL.Query()["code"]

	// меняем код на access токен
	tok, err := r.conf.Exchange(ctx, authCode[0])
	if err != nil {
		log.Fatal().Msgf("Fatal error: %v", err)
	}
	client, err := vkApi.NewClientWithOptions(
		vkApi.WithToken(tok.AccessToken),
	)
	if err != nil {
		log.Fatal().Msgf("Fatal error: %v", err)
	}
	var response groups.GroupList

	err = client.CallMethod(
		"users.getSubscriptions",
		vkApi.RequestParams{"extended": "1", "count": "200", "v": "5.131"},
		&response,
	)
	if err != nil {
		log.Fatal().Msgf("Fatal error: %v", err)
	}

	c.JSON(http.StatusOK, &response)
}
