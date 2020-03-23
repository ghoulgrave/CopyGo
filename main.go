package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/leaanthony/mewn"
	"github.com/spf13/viper"
	"github.com/wailsapp/wails"
	"io"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)


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
	var logs string
	sysType := runtime.GOOS
	var output []byte
	var err error
	var command string
	if sysType == "windows" {
		//fmt.Println("WIN")
		var enc mahonia.Decoder
		enc = mahonia.NewDecoder("gbk")
		cmd := exec.Command("cmd.exe", "/c", "cd "+selectedProject.Dir_path+" && "+selectedProject.Dir_path[0:2]+`&&svn log -r {`+strings.Replace(kssj, " ", "T", -1)+`}:{`+strings.Replace(jssj, " ", "T", -1)+`} -v`)
		output, err = cmd.Output()
		logs = enc.ConvertString(string(output))
	} else {
		command := `` + runningPath + `/resource/log.sh ` + selectedProject.Dir_path + ` ` + strings.Replace(kssj, " ", "T", -1) + ` ` + strings.Replace(jssj, " ", "T", -1) + ` .`
		cmd := exec.Command("/bin/bash", "-c", command)
		output, err = cmd.Output()
		logs = string(output)
	}
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
	}
	//找到的所有日志
	baseLog := strings.Split(logs, "------------------------------------------------------------------------")
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
	for _, s := range searchLog {
		svnInfoTemp = SvnInfo{}
		if s == "" {
			continue
		}
		res1 := strings.Split(s, "\n")
		isSvnFiles := false
		isSvnLogLine := -1
		for i, s2 := range res1 {
			if strings.Contains(s2, czr) && strings.Contains(s2, "|") {
				uAt := strings.Split(res1[i], "|")
				svnInfoTemp.Name = strings.TrimSpace(uAt[1])
				svnInfoTemp.Time = strings.TrimSpace(uAt[2])
				svnInfoTemp.Version = strings.TrimSpace(uAt[0])
			}
			if i != 0 && i != len(res1)-1 && s2 == "" {
				isSvnLogLine = i
			}
			if len(s2) == 1 || s2 == "" {
				isSvnFiles = false
				continue
			}
			if strings.Contains(s2, "Changed paths:") {
				isSvnFiles = true
				continue
			}
			if isSvnFiles {
				svnInfoTemp.Path = s2[4:]
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
	//fmt.Println(str)
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
return ""
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

func CmdAndChangeDirToShow(dir string, commandName string, params []string) error {
	//cmd := exec.Command("cmd.exe", "/c", "cd D:\\1-WorkSpace\\0_SvnProject\\bdcdj && d: && dir")

	cmd := exec.Command(commandName, params...)
	fmt.Println("CmdAndChangeDirToFile", dir, cmd.Args)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("cmd.StdoutPipe: ", err)
		return err
	}
	cmd.Stderr = os.Stderr
	//cmd.Dir = dir
	err = cmd.Start()
	if err != nil {
		return err
	}
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}
	err = cmd.Wait()
	return err
}
