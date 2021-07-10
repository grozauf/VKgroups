module github.com/grozauf/VKgroups

go 1.16

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/go-vk-api/vk v0.0.0-20200129183856-014d9b8adc96
	github.com/rs/zerolog v1.23.0
	golang.org/x/oauth2 v0.0.0-20210628180205-a41e5a781914
)

replace github.com/go-vk-api/vk => ../vk
