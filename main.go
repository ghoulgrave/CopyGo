package main

import (
	"fmt"
	"github.com/leaanthony/mewn"
	"github.com/spf13/viper"
	"github.com/wailsapp/wails"
	"os"
	"path"
)

func basic() string {
	return "Hello World!"
}

//我的配置文件
var MyConfig Config

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	ePath, _ := os.Executable()
	fmt.Println(ePath)
	conpath := path.Dir(ePath)

	viper.SetConfigName("copy")             // 设置配置文件名 (不带后缀)
	viper.AddConfigPath(conpath + "/conf/") // 第一个搜索路径
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig() // 读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Unmarshal(&MyConfig) // 将配置信息绑定到结构体上

	//更新配置文件
	//fmt.Println(MyConfig)
	v := MyConfig.Conf
	k := Confs{Sub_path: "http://ssss", Name: "ttttt", Dir_path: "/user/dfds/sdfsfd", Out_path: "/sss.dd/fffs/affd"}
	v = append(v, k)
	viper.Set("conf", v)

	viper.Unmarshal(&MyConfig)
	//fmt.Println(MyConfig)

	viper.WriteConfig()
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
	app.Bind(sub)
	app.Bind(logselect)
	app.Run()
}

func sub(xx string) string {
	println(xx)
	return "ok"
}

func logselect() string {

	return `[{
      value: '选项1',
      label: '黄金糕22'
    }, {
      value: '选项2',
      label: '双皮奶33'
    }, {
      value: '选项3',
      label: '蚵仔煎44'
    }, {
      value: '选项4',
      label: '龙须面55'
    }, {
      value: '选项5',
      label: '北京烤鸭66'
    }]`
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

type SvnInfo struct {
	Name    string `json:"name"`
	Time    string `json:"time"`
	Version string `json:"version"`
	Path    string `json:"path"`
	SubLogs string `json:"sublogs"`
}

type MyData struct {
	Code  int64     `json:"code"`
	Msg   string    `json:"msg"`
	Count int64     `json:"count"`
	Data  []SvnInfo `json:"data"`
	Other int64     `json:"-"` // 直接忽略字段
}

type Search struct {
	ConfigNum string
	Kssj      string
	Jssj      string
	Name      string
	Build     string
}

type Config struct {
	Username string
	Conf     []Confs
}

type Confs struct {
	Sub_path string
	Dir_path string
	Svn_path string
	Out_path string
	Name     string
}
