package assets

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets9f0ae851af9adb08e3242917d7d52dc270c7bae2 = "<!DOCTYPE html>\n<html>\n{{template \"header\" .}}\n<body>\n\t<a href={{ .groupsUrl }}> Получить список групп!\t</a>\n</body>\n</html>"
var _Assets7c932e997237213be967a49691ce1726e11d3921 = "{{define \"header\"}}\n<head>\n    <link rel=\"stylesheet\" href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css\"\n          integrity=\"sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh\"\n          crossorigin=\"anonymous\">\n    <script src=\"https://code.jquery.com/jquery-3.4.1.slim.min.js\"\n            integrity=\"sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n\"\n            crossorigin=\"anonymous\"></script>\n    <script src=\"https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js\"\n            integrity=\"sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo\"\n            crossorigin=\"anonymous\"></script>\n    <script src=\"https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/js/bootstrap.min.js\"\n            integrity=\"sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6\"\n            crossorigin=\"anonymous\"></script>\n</head>\n{{end}}"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"templates"}, "/templates": []string{"index.html", "headers.html"}}, map[string]*assets.File{
	"/templates/index.html": &assets.File{
		Path:     "/templates/index.html",
		FileMode: 0x1b4,
		Mtime:    time.Unix(1625917808, 1625917808331077955),
		Data:     []byte(_Assets9f0ae851af9adb08e3242917d7d52dc270c7bae2),
	}, "/templates/headers.html": &assets.File{
		Path:     "/templates/headers.html",
		FileMode: 0x1b4,
		Mtime:    time.Unix(1625917731, 1625917731937953190),
		Data:     []byte(_Assets7c932e997237213be967a49691ce1726e11d3921),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001fd,
		Mtime:    time.Unix(1625916945, 1625916945989513123),
		Data:     nil,
	}, "/templates": &assets.File{
		Path:     "/templates",
		FileMode: 0x800001fd,
		Mtime:    time.Unix(1625917602, 1625917602284368058),
		Data:     nil,
	}}, "")
