package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/grozauf/VKgroups/internal/assets"
	"github.com/grozauf/VKgroups/internal/oauth"
	"github.com/grozauf/VKgroups/internal/router"
	"github.com/headzoo/surf"
	"github.com/martinlindhe/inputbox"
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

	bow := surf.NewBrowser()
	err := bow.Open(conf.AuthURL("state"))
	if err != nil {
		panic(err)
	}

	fmt.Println("Url: ", bow.Url())

	login, _ := inputbox.InputBox("LOGIN", "Enter your login", "")
	pass, _ := inputbox.InputBox("PASSWORD", "Enter your password", "")

	fm, _ := bow.Form("form#login_submit")
	fm.Input("email", login)
	fm.Input("pass", pass)
	if fm.Submit() != nil {
		panic(err)
	}
	fmt.Println("Url: ", bow.Url())

	code, _ := inputbox.InputBox("CODE", "Type a code", "")

	fm, _ = bow.Form("form")
	fm.Input("code", code)
	if fm.Submit() != nil {
		panic(err)
	}

	fmt.Println("Url: ", bow.Url())
	fmt.Println("Body: ", bow.Body())

	// if we got form, so vk wants captcha key
	fm, _ = bow.Form("form")
	if fm != nil {
		captcha_key, _ := inputbox.InputBox("CAPTCHA", "Type a key", "")

		fm.Input("captcha_key", captcha_key)
		if fm.Submit() != nil {
			panic(err)
		}
		fmt.Println("Url: ", bow.Url())
		fmt.Println("Body: ", bow.Body())
	}

	authUrlQuery := bow.Url().Query()["authorize_url"]
	authUrl, _ := url.QueryUnescape(authUrlQuery[0])

	authUrl = strings.Replace(authUrl, "#", "?", 1)
	u, _ := url.Parse(authUrl)
	token := u.Query()["access_token"][0]
	userId := u.Query()["user_id"][0]

	fmt.Println("Got token: ", token, " user_id: ", userId)

	conf.SetToken(token)
	conf.SetUser(userId)

	leaveUrl := "https://api.vk.com/method/groups.leave?group_id=141391486&v=5.131&access_token=" + token
	err = bow.Open(leaveUrl)
	fmt.Println("After leave: ", bow.Body())

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
