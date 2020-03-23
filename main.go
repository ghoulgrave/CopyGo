package main

import (
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/leaanthony/mewn"
	"github.com/spf13/viper"
	"github.com/wailsapp/wails"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

func basic() string {
	return "Hello World!"
}

//本项目的物理地址
var runningPath string

//我的配置文件
var MyConfig Config

//项目配置信息（全部）
var projectConfs []Confs

//========================
//全局变量，但是每次查询之后都应该变换
//需求编号
var RequsNum string

//提交日志
var SubLogs string

//========================
func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	ePath, _ := os.Executable()
	fmt.Println(ePath)
	runningPath = path.Dir(ePath)

	viper.SetConfigName("copy")                 // 设置配置文件名 (不带后缀)
	viper.AddConfigPath(runningPath + "/conf/") // 第一个搜索路径
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig() // 读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Unmarshal(&MyConfig) // 将配置信息绑定到结构体上

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

	projectConfs = MyConfig.Conf

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1440,
		Height:    900,
		Title:     "copy",
		JS:        js,
		CSS:       css,
		Colour:    "#f5fffa",
		Resizable: true,
	})

	app.Bind(basic)
	app.Bind(getPaths)
	app.Bind(sub)
	app.Bind(projectName)
	app.Run()
}

func sub(projectname string, kssj string, jssj string, czr string) string {

	var selectedProject Confs
	for _, conf := range projectConfs {
		if conf.Name == projectname {
			selectedProject = conf
			break
		}
	}
	sysType := runtime.GOOS
	var output []byte
	var err error
	var command string
	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("gbk")
	if sysType == "windows" {
		fmt.Println("WIN")
		fmt.Println(selectedProject.Dir_path)
		fmt.Println(selectedProject.Dir_path[0:2])

		//command := `` + runningPath + `/resource/log.bat ` + selectedProject.Dir_path +` ` +selectedProject.Dir_path[0:2] +` `+ ` ""` + kssj + `" "` + jssj + `" .`
		//cmd := exec.Command("cmd.exe",  command)
		//output, err = cmd.Output()
		cmd := exec.Command("cmd.exe", "/c", "cd "+selectedProject.Dir_path+" && "+selectedProject.Dir_path[0:2]+`&&svn log -r {`+strings.Replace(kssj, " ", "T", -1)+`}:{`+strings.Replace(jssj, " ", "T", -1)+`} -v`)
		output, err = cmd.Output()

	} else {
		command := `` + runningPath + `/resource/log.sh ` + selectedProject.Dir_path + ` ` + strings.Replace(kssj, " ", "T", -1) + ` ` + strings.Replace(jssj, " ", "T", -1) + ` .`
		cmd := exec.Command("/bin/bash", "-c", command)
		output, err = cmd.Output()

	}
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
	}
	//获取返回日志
	var kk = string(output)
	kk = enc.ConvertString(kk)
	//fmt.Println("UTF-8 to GBK:", enc.ConvertString(kk))
	fmt.Println(kk)

	//fmt.Println(kk)
	//找到的所有日志
	baseLog := strings.Split(kk, "------------------------------------------------------------------------")
	//筛选之后的日志
	var searchLog []string
	for _, s := range baseLog {
		if strings.Contains(s, czr) {
			searchLog = append(searchLog, s)
		}
	}
	RequsNum = ""
	SubLogs = ""
	var svnInfoTemp SvnInfo
	li := []SvnInfo{}
	fmt.Println(len(searchLog))
	for _, s := range searchLog {
		svnInfoTemp = SvnInfo{}
		fmt.Println("==:" + s + ":==")
		if s == "" {
			fmt.Println("不操作")
			continue
		}
		fmt.Println("   ")
		res1 := strings.Split(s, "\n")
		isSvnFiles := false
		isSvnLogLine := -1
		for i, s2 := range res1 {
			fmt.Println("1==================1:")
			sss :="开始"+s2+"完"
			fmt.Println(sss)
			fmt.Println(len(s2))
			fmt.Println("2===============2")

			if strings.Contains(s2, czr) && strings.Contains(s2, "|") {
				uAt := strings.Split(res1[i], "|")
				svnInfoTemp.Name = strings.TrimSpace(uAt[1])
				svnInfoTemp.Time = strings.TrimSpace(uAt[2])
				svnInfoTemp.Version = strings.TrimSpace(uAt[0])
			}
			if i != 0 && i != len(res1)-1 && s2 == "" {
				isSvnLogLine = i
			}
			if len(s2) == 1 || s2=="" {
				fmt.Println("不是文件了")
				isSvnFiles = false
				continue
			}
			if strings.Contains(s2, "Changed paths:") {
				fmt.Println("可以是文件了")
				isSvnFiles = true
				continue
			}
			if isSvnFiles {

				svnInfoTemp.Path = s2[4:]

				fmt.Println("33=======↓33")
				fmt.Println(svnInfoTemp)
				li = append(li, svnInfoTemp)
			}

			if i == isSvnLogLine+1 {
				temp := strings.Split(s2, ":")
				if len(temp) == 1 {
					temp = strings.Split(s2, "：")
				}
				if len(temp) > 1 {
					RequsNum += temp[1] + " "
				}
			}
			if i > isSvnLogLine+1 && isSvnLogLine > 0 {
				SubLogs += s2 + "\r\n"
			}
		}
	}
	//strVre := "["
	//for _, info := range li {
	//
	//}
	lang, err := json.Marshal(li)
	if err == nil {
	}
	str := string(lang)
	fmt.Println(str)
	return str
}

func projectName() string {

	myConfs := MyConfig.Conf
	vreStr := "["
	for i, conf := range myConfs {
		if i == len(myConfs)-1 {
			vreStr += `{value:'` + conf.Name + `',label:'` + conf.Name + `'}`
		} else {
			vreStr += `{value:'` + conf.Name + `',label:'` + conf.Name + `'},`
		}
	}
	vreStr += "]"
	fmt.Println(vreStr)
	return vreStr
}

func getPaths() string {

	str := "[{\"name\":\"zyy\",\"time\":\"2020-03-20 15:51:55 +0800 (五, 20  3 2020)\",\"version\":\"r251597\",\"path\":\" /bdcdj/branches/bdcdj_dbqy/src/main/resources/META-INF/conf/bdcdj/application.properties\",\"sublogs\":\"\"},{\"name\":\"zyy\",\"time\":\"2020-03-21 14:06:08 +0800 (六, 21  3 2020)\",\"version\":\"r251734\",\"path\":\" /bdcdj/branches/bdcdj_dbqy/src/main/java/cn/gtmap/bdcdj/service/impl/JzReadHtxxServiceImpl.java\",\"sublogs\":\"\"}]"
	//fmt.Println(str)
	return str
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
