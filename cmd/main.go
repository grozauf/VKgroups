package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/grozauf/VKgroups/internal/assets"
	"github.com/grozauf/VKgroups/internal/oauth"
	"github.com/grozauf/VKgroups/internal/router"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	conf := oauth.NewConfig(
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
		os.Getenv("REDIRECT_URL"),
		[]string{"groups"},
		oauth.VKEndpoint,
	)

	router := router.NewRouter(conf)

	srv := gin.Default()

	t, err := assets.LoadTemplate()
	if err != nil {
		panic(err)
	}

	srv.SetHTMLTemplate(t)
	srv.GET("/", router.Root)
	srv.GET("/fragment", router.Fragment)
	srv.GET("/groups", router.Groups)
	srv.POST("/delete", router.Delete)
	srv.Run(":8080")
}
