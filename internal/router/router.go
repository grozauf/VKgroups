package router

import (
	"fmt"
	"net/http"
	"time"

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

	c.HTML(
		http.StatusOK,
		"/templates/index.html",
		gin.H{
			"groupsUrl": fmt.Sprintf("/groups?access_token=%s&user_id=%s", r.conf.Token(), r.conf.User()),
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
	authToken := c.Request.URL.Query()["access_token"][0]
	userId := c.Request.URL.Query()["user_id"][0]

	client, err := vkApi.NewClientWithOptions(
		vkApi.WithToken(authToken),
	)
	if err != nil {
		log.Error().Msgf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	var response groups.GroupList

	err = client.CallMethod(
		"groups.get",
		vkApi.RequestParams{"user_id": userId, "extended": "1", "count": "200", "v": "5.131"},
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
			"token": authToken,
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
	for _, groupId := range form.Groups {
		log.Info().Msgf("Leave group with id: %s", groupId)
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
		log.Info().Msg("sleep 4 seconds...")
		time.Sleep(time.Second * 4)
	}
}
