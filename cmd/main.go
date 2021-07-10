package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/grozauf/VKgroups/internal/assets"
	"github.com/grozauf/VKgroups/internal/router"
	"github.com/rs/zerolog"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	conf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes:       []string{"groups"},
		Endpoint:     vk.Endpoint,
	}

	router := router.NewRouter(conf)

	srv := gin.Default()

	t, err := assets.LoadTemplate()
	if err != nil {
		panic(err)
	}

	srv.SetHTMLTemplate(t)
	srv.GET("/", router.Root)
	srv.GET("/groups", router.Groups)
	srv.Run(":8080")
}
