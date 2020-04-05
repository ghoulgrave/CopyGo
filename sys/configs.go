package sys

import (
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"os/exec"
	"runtime"
	"strings"
)

//全局参数
//本项目的物理地址
var RunningPath string

//我的配置文件
var MyConfig Config

//项目配置信息（全部）
var ProjectConfs []Confs

//========================
//全局变量，但是每次查询之后都应该变换
//需求编号
var RequsNum string

//提交日志
var SubLogs string

//获取项目信息
func (s *Stats) GetProjectName() string {
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

//已经提交的日志信息
func (s *Stats) GetSubmitedLogInfo(projectname string, kssj string, jssj string, czr string) string {
	var selectedProject Confs
	for _, conf := range ProjectConfs {
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
		command := `` + RunningPath + `/resource/log.sh ` + selectedProject.Dir_path + ` ` + strings.Replace(kssj, " ", "T", -1) + ` ` + strings.Replace(jssj, " ", "T", -1) + ` .`
		cmd := exec.Command("/bin/bash", "-c", command)
		output, err = cmd.Output()
		if err != nil {
			//Execute Shell: failed with error:exit status
			fmt.Println(err)
			fmt.Println("日志获取失败了")
		}
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

//获取所有的项目
func (s *Stats) GetAllProject(projectName string) {

	projectNames := strings.Split(projectName, "^")
	for _, name := range projectNames {
		if name != "" {
			fmt.Println(name)
		}
	}
	s.runtime.Events.Emit("builds_pl", s.GetOuts("xxxxx"))
}
