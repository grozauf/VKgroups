package assets

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets9f0ae851af9adb08e3242917d7d52dc270c7bae2 = "<html>\n\t<a href={{ .authUrl }}> Получить список групп!\t</a>\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"templates"}, "/templates": []string{"index.html"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001fd,
		Mtime:    time.Unix(1625916945, 1625916945989513123),
		Data:     nil,
	}, "/templates": &assets.File{
		Path:     "/templates",
		FileMode: 0x800001fd,
		Mtime:    time.Unix(1625897060, 1625897060484177424),
		Data:     nil,
	}, "/templates/index.html": &assets.File{
		Path:     "/templates/index.html",
		FileMode: 0x1b4,
		Mtime:    time.Unix(1625852583, 1625852583883513001),
		Data:     []byte(_Assets9f0ae851af9adb08e3242917d7d52dc270c7bae2),
	}}, "")
