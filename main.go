package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

func auth(c *gin.Context) {
	ctx := context.Background()
	// получаем код от API VK из квери стринга
	authCode := c.Request.URL.Query()["code"]
	// меняем код на access токен
	tok, err := conf.Exchange(ctx, authCode[0])
	if err != nil {
		log.Fatal().Msgf("Fatal error: %v", err)
	}

	apiUrl := fmt.Sprintf("https://api.vk.com/method/users.getSubscriptions?extended=1&count=200&v=5.131&access_token=%s", tok.AccessToken)

	log.Debug().Msgf("get subs from url: %s", apiUrl)

	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Fatal().Msgf("Fatal error: %v", err)
	}
	defer resp.Body.Close()

	// headers

	for name, values := range resp.Header {
		c.Writer.Header()[name] = values
	}

	// status (must come after setting headers and before copying body)

	c.Writer.WriteHeader(resp.StatusCode)

	io.Copy(c.Writer, resp.Body)
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
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v\n", url)

	r := gin.New()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", root)
	r.GET("/auth", auth)
	r.Run(":8080")
}
