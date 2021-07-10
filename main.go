package main

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	vkApi "github.com/go-vk-api/vk"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
)

var (
	conf *oauth2.Config
)

func root(c *gin.Context) {
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"authUrl": url,
		},
	)
}

func groups(c *gin.Context) {
	ctx := context.Background()
	// получаем код от API VK из квери стринга
	authCode := c.Request.URL.Query()["code"]

	// меняем код на access токен
	tok, err := conf.Exchange(ctx, authCode[0])
	if err != nil {
		log.Fatal().Msgf("Fatal error: %v", err)
	}
	client, err := vkApi.NewClientWithOptions(
		vkApi.WithToken(tok.AccessToken),
	)
	if err != nil {
		log.Fatal().Msgf("Fatal error: %v", err)
	}
	var response struct {
		Count  int64 `json:"count"`
		Groups []struct {
			Id           int64  `json:"id"`
			Name         string `json:"name"`
			ScreenName   string `json:"screen_name"`
			IsClosed     int64  `json:"is_closed"`
			Deactivated  string `json:"deactivated"`
			IsAdmin      int64  `json:"is_admin"`
			AdminLevel   int64  `json:"admin_level"`
			IsMember     int64  `json:"is_member"`
			IsAdvertiser int64  `json:"is_advertiser"`
			InvitedBy    int64  `json:"invited_by"`
			Type         string `json:"type"`
			Photo50      string `json:"photo_50"`
			Photo100     string `json:"photo_100"`
			Photo200     string `json:"photo_200"`
		} `json:"items"`
	}

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

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	conf = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes:       []string{"groups"},
		Endpoint:     vk.Endpoint,
	}

	r := gin.New()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", root)
	r.GET("/groups", groups)
	r.Run(":8080")
}
