module github.com/grozauf/VKgroups

go 1.16

require (
	github.com/PuerkitoBio/goquery v1.7.0 // indirect
	github.com/gin-gonic/gin v1.7.2
	github.com/go-vk-api/vk v0.0.0-20200129183856-014d9b8adc96
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/headzoo/surf v1.0.0
	github.com/headzoo/ut v0.0.0-20181013193318-a13b5a7a02ca // indirect
	github.com/jessevdk/go-assets v0.0.0-20160921144138-4f4301a06e15
	github.com/martinlindhe/inputbox v0.0.0-20210326232244-b26136a79ad0
	github.com/rs/zerolog v1.23.0
	golang.org/x/sys v0.0.0-20210320140829-1e4c9ba3b0c4 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/headzoo/surf.v1 v1.0.0
)

replace github.com/go-vk-api/vk => ../vk
