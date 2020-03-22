package main

import (
	"fmt"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"os"
	"path"
)

func basic() string {
	return "Hello World!"
}

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	ePath, _ := os.Executable()
	fmt.Println(ePath)
	fmt.Println(path.Dir(ePath))

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "copy",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(basic)
	app.Bind(datas)
	app.Run()
}

func datas() string {
	return `[{
	date: '2016-05-07',
	name: 'zhangyiyang',
	address: '/ddd/ddd/ddd/d/dd'
},{
	date: '2016-05-17',
	name: 'zhangyiyang1',
	address: '/ddd/ddd/ddd/d/dd222'
},{
	date: '2016-05-27',
	name: 'zhangyiyang2',
	address: '/ddd/ddd/ddd/d/11122'
},{
	date: '2016-05-18',
	name: 'zhangyiyang3',
	address: '/ddd/ddd/ddd/d/dd123'
},{
	date: '2020-05-18',
	name: 'zhangyiyang4',
	address: '/ddd/dd4444'
}




]`
}
