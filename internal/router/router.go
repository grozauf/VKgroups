package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grozauf/VKgroups/internal/groups"
	"github.com/grozauf/VKgroups/internal/oauth"

	"github.com/rs/zerolog/log"

	vkApi "github.com/go-vk-api/vk"
)

type Router interface {
	Root(c *gin.Context)
	Fragment(c *gin.Context)
	Groups(c *gin.Context)
	Delete(c *gin.Context)
}

func NewRouter(conf oauth.Config) Router {
	return &router{conf: conf}
}

type router struct {
	conf oauth.Config
}

func (r router) Root(c *gin.Context) {
	url := r.conf.AuthURL("state")
	c.HTML(
		http.StatusOK,
		"/templates/index.html",
		gin.H{
			"groupsUrl": url,
		},
	)
}

func (r router) Fragment(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"/templates/fragment.html",
		gin.H{},
	)
}

func (r router) Groups(c *gin.Context) {
	authToken := c.Request.URL.Query()["access_token"]

	client, err := vkApi.NewClientWithOptions(
		vkApi.WithToken(authToken[0]),
	)
	if err != nil {
		log.Error().Msgf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	var response groups.GroupList

	err = client.CallMethod(
		"users.getSubscriptions",
		vkApi.RequestParams{"extended": "1", "count": "200", "v": "5.131"},
		&response,
	)
	if err != nil {
		log.Error().Msgf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	//c.JSON(http.StatusOK, &response)
	c.HTML(
		http.StatusOK,
		"/templates/groups.html",
		gin.H{
			"list":  response,
			"token": authToken[0],
		},
	)
}

type groupsForm struct {
	Groups []string `form:"groups[]"`
}

func (r router) Delete(c *gin.Context) {
	token := c.PostForm("token")

	client, err := vkApi.NewClientWithOptions(
		vkApi.WithToken(token),
	)
	if err != nil {
		log.Error().Msgf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var form groupsForm
	c.ShouldBind(&form)
	for groupId := range form.Groups {
		err = client.CallMethod(
			"groups.leave",
			vkApi.RequestParams{"group_id": groupId, "v": "5.131", "state": "state"},
			nil,
		)
		if err != nil {
			log.Error().Msgf("Error: %v", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

	}
}
