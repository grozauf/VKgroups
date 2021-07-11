module github.com/grozauf/VKgroups

go 1.16

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/go-vk-api/vk v0.0.0-20200129183856-014d9b8adc96
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/jessevdk/go-assets v0.0.0-20160921144138-4f4301a06e15
	github.com/kr/pretty v0.1.0 // indirect
	github.com/rs/zerolog v1.23.0
	golang.org/x/sys v0.0.0-20210320140829-1e4c9ba3b0c4 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

replace github.com/go-vk-api/vk => ../vk
