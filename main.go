package main

import (
	"copy/sys"
	"fmt"
	"github.com/leaanthony/mewn"
	"github.com/spf13/viper"
	"github.com/wailsapp/wails"
	"os"
	"path"
)

//========================
func main() {

	//var k []string
	//k = append(k,"/c" )
	//k = append(k,"cd D:\\1-WorkSpace\\0_SvnProject\\bdcdj && d: && mvn clean && mvn install" )
	//
	//
	//CmdAndChangeDirToShow("dir string", "cmd.exe", k)

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	ePath, _ := os.Executable()
	fmt.Println(ePath)
	sys.RunningPath = path.Dir(ePath)

	viper.SetConfigName("copy")                     // 设置配置文件名 (不带后缀)
	viper.AddConfigPath(sys.RunningPath + "/conf/") // 第一个搜索路径
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig() // 读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Unmarshal(&sys.MyConfig) // 将配置信息绑定到结构体上

	////更新配置文件
	////fmt.Println(MyConfig)
	//v := MyConfig.Conf
	//k := Confs{Sub_path: "http://ssss", Name: "ttttt", Dir_path: "/user/dfds/sdfsfd", Out_path: "/sss.dd/fffs/affd"}
	//v = append(v, k)
	//viper.Set("conf", v)
	//
	//viper.Unmarshal(&MyConfig)
	////fmt.Println(MyConfig)
	//viper.WriteConfig()

	sys.ProjectConfs = sys.MyConfig.Conf

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1440,
		Height:    900,
		Title:     "copy",
		JS:        js,
		CSS:       css,
		Colour:    "#000000",
		Resizable: true,
	})

	stats := &sys.Stats{}
	app.Bind(stats)
	app.Run()
}
