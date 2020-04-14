package sys

import (
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

//全局参数
//地址分割符
var PathSeparator = string(os.PathSeparator)

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
func (s *ThisCopy) GetProjectName() string {
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
	//fmt.Println(vreStr)
	return vreStr
}

//获取查询姓名信息
func (s *ThisCopy) GetEName() string {
	return MyConfig.Searchname
}

//获取中文姓名信息
func (s *ThisCopy) GetCName() string {
	return MyConfig.Username
}

//批量输出地址
func (s *ThisCopy) GetPlOuPath() string {
	return MyConfig.PlOutPath
}

//获取所有项目配置信息
func (s *ThisCopy) GetProjectConfs() []Confs {
	return MyConfig.Conf
}

//jar文件名
func (s *ThisCopy) GetJarNames() []string {
	return MyConfig.Jarnames
}

//获取已经提交的日志信息
func (s *ThisCopy) GetSubmitedLogInfo(projectname string, kssj string, jssj string, czr string) string {
	fmt.Println(projectname)
	var selectedProject Confs
	for _, conf := range ProjectConfs {
		if conf.Name == projectname {
			selectedProject = conf
			break
		}
	}
	if selectedProject.Name == "" {
		li := []SvnInfo{}
		svnInfoTemp := SvnInfo{}
		svnInfoTemp.Name = ""
		svnInfoTemp.Path = "项目信息异常。无法获取日志"
		svnInfoTemp.Version = ""
		svnInfoTemp.Time = ""
		li = append(li, svnInfoTemp)
		lang, err := json.Marshal(li)
		if err == nil {
		}
		str := string(lang)
		return str
	}
	var logs string
	sysType := runtime.GOOS
	var output []byte
	var err error
	if sysType == "windows" {
		//fmt.Println("WIN")
		var enc mahonia.Decoder
		enc = mahonia.NewDecoder("gbk")
		cmd := exec.Command("cmd.exe", "/c", "cd "+selectedProject.Dir_path+" && "+selectedProject.Dir_path[0:2]+`&&svn log -r {`+strings.Replace(kssj, " ", "T", -1)+`}:{`+strings.Replace(jssj, " ", "T", -1)+`} -v`)
		output, err = cmd.Output()
		if err != nil {
			//Execute Shell: failed with error:exit status
			fmt.Println(err)
			fmt.Println("日志获取失败了")
		}
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
		//fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		li := []SvnInfo{}
		svnInfoTemp := SvnInfo{}
		svnInfoTemp.Name = ""
		svnInfoTemp.Path = "SVN连接超时，日志获取失败！"
		svnInfoTemp.Version = ""
		svnInfoTemp.Time = ""
		li = append(li, svnInfoTemp)
		lang, err := json.Marshal(li)
		if err == nil {
		}
		str := string(lang)
		return str
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
	svnPaths := strings.Split(selectedProject.Svn_path, "/")
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
				svnInfoTemp.Path = strings.Trim(s2[4:], " ")
				fmt.Println("dd:[" + svnInfoTemp.Path + "]")
				var pathTemp = ""
				for _, path := range svnPaths {
					fmt.Println("xx:[" + path + "]")
					if path != "" && strings.HasPrefix(svnInfoTemp.Path, "/"+path) {
						pathTemp = "/" + path
						fmt.Println("pathTemp:[" + pathTemp + "]")
						break
					}
				}
				fmt.Println("pathTemp:[" + pathTemp + "]")
				fmt.Println(strings.Index(selectedProject.Svn_path, pathTemp))

				startStr := selectedProject.Svn_path[strings.Index(selectedProject.Svn_path, pathTemp):]
				fmt.Println(startStr)
				if strings.HasPrefix(svnInfoTemp.Path, startStr) {
					li = append(li, svnInfoTemp)
				}
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
	lang, err := json.Marshal(li)
	if err == nil {
	}
	str := string(lang)
	return str
}

//获取所有的项目
func (s *ThisCopy) GetAllProject(projectName string, kssj string, jssj string, czr string, buildOrNot bool) {
	fmt.Println(kssj)
	fmt.Println(jssj)
	//全部项目的日志信息
	var proCheckedSvns []ProCheckedSvn
	projectNames := strings.Split(projectName, "^")
	for _, name := range projectNames {
		if name != "" {
			fmt.Println(name)
			proCheckedSvn := ProCheckedSvn{}
			for _, conf := range ProjectConfs {
				if conf.Name == name {
					proCheckedSvn.Project = conf
					break
				}
			}

			sysType := runtime.GOOS
			var output []byte
			var logs string
			var err error
			if sysType == "windows" {
				//fmt.Println("WIN")
				var enc mahonia.Decoder
				enc = mahonia.NewDecoder("gbk")
				cmd := exec.Command("cmd.exe", "/c", "cd "+proCheckedSvn.Project.Dir_path+" && "+proCheckedSvn.Project.Dir_path[0:2]+`&&svn log -r {`+strings.Replace(kssj, " ", "T", -1)+`}:{`+strings.Replace(jssj, " ", "T", -1)+`} -v`)
				output, err = cmd.Output()
				if err != nil {
					//Execute Shell: failed with error:exit status
					fmt.Println(err)
					fmt.Println("日志获取失败了")
					return
				}
				logs = enc.ConvertString(string(output))
			} else {
				command := `` + RunningPath + `/resource/log.sh ` + proCheckedSvn.Project.Dir_path + ` ` + strings.Replace(kssj, " ", "T", -1) + ` ` + strings.Replace(jssj, " ", "T", -1) + ` .`
				cmd := exec.Command("/bin/bash", "-c", command)
				output, err = cmd.Output()
				if err != nil {
					//Execute Shell: failed with error:exit status
					fmt.Println(err)
					fmt.Println("日志获取失败了")
					return
				}
				logs = string(output)
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
			//多项目一起提交时区分为本项目的文件
			svnPaths := strings.Split(proCheckedSvn.Project.Svn_path, "/")
			var svnInfoTemp SvnInfo
			for _, str := range searchLog {
				svnInfoTemp = SvnInfo{}
				if str == "" {
					continue
				}
				res1 := strings.Split(str, "\n")
				isSvnFiles := false
				isSvnLogLine := -1
				for i, s2 := range res1 {
					s.runtime.Events.Emit("builds_pl", s.GetOuts("[COPY-INFO] "+s2))
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
						svnInfoTemp.Path = strings.Trim(s2[4:], " ")
						fmt.Println("dd:[" + svnInfoTemp.Path + "]")
						var pathTemp = ""
						for _, path := range svnPaths {
							fmt.Println("xx:[" + path + "]")
							if path != "" && strings.HasPrefix(svnInfoTemp.Path, "/"+path) {
								pathTemp = "/" + path
								fmt.Println("pathTemp:[" + pathTemp + "]")
								break
							}
						}
						fmt.Println("pathTemp:[" + pathTemp + "]")
						fmt.Println(strings.Index(proCheckedSvn.Project.Svn_path, pathTemp))

						startStr := proCheckedSvn.Project.Svn_path[strings.Index(proCheckedSvn.Project.Svn_path, pathTemp):]
						fmt.Println(startStr)
						if strings.HasPrefix(svnInfoTemp.Path, startStr) {
							proCheckedSvn.Svns = append(proCheckedSvn.Svns, svnInfoTemp)
						}
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
			s.runtime.Events.Emit("builds_pl", s.GetOuts("[COPY-INFO] =========================================="))
			s.runtime.Events.Emit("builds_pl", s.GetOuts("[COPY-INFO] "+proCheckedSvn.Project.Name+" 日志获取完成。"))
			s.runtime.Events.Emit("builds_pl", s.GetOuts("[COPY-INFO] =========================================="))
			proCheckedSvns = append(proCheckedSvns, proCheckedSvn)
		}
	}
	dateNow := time.Now().Format("20060102150405")
	//输出文件夹位置
	plOuPath := MyConfig.PlOutPath
	//循环 编译和获取文件
	for _, svn := range proCheckedSvns {
		//无日志情况
		if len(svn.Svns) == 0 {
			continue
		}
		//保证批量输出到同一个文件夹中
		if plOuPath == "" {
			plOuPath = svn.Project.Out_path
		}
		svn.Project.Out_path = plOuPath
		if buildOrNot {
			s.CmdAndChangeDirToShow(svn.Project.Dir_path, true)
		}
		_, fileDirPath, _, fileType, _ := s.copyfiles(svn.Project, svn.Svns, dateNow)
		//复制jar文件
		if fileType == "jar" {
			s.copyJars(svn.Project, dateNow, fileDirPath)
		}
		ZipDir(svn.Project.Out_path+PathSeparator+dateNow+PathSeparator+fileDirPath, svn.Project.Out_path+PathSeparator+dateNow+PathSeparator+fileDirPath+".zip")
	}
	s.runtime.Events.Emit("builds_pl", s.GetOuts("[COPY-INFO] 全部完成"))
	s.runtime.Events.Emit("builds_pl", s.GetOuts("[COPY-INFO] file folder : "+plOuPath+PathSeparator+dateNow))
}

//更新系统配置
func (s *ThisCopy) UpSysConfig(cname string, ename string, plOuPath string, textarea string) string {
	viper.Set("Username", cname)
	viper.Set("Searchname", ename)
	viper.Set("PlOutPath", plOuPath)
	k := strings.Split(textarea, "\n")
	viper.Set("Jarnames", k)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
		return "操作失败"
	} else {
		return "操作成功"
	}

}

//更新项目配置
func (s *ThisCopy) SaveOrUpdatePorjectConfig(project string) string {
	str := []byte(project)
	conf := Confs{}
	json.Unmarshal(str, &conf)
	existConfs := MyConfig.Conf
	//是否重名
	var doubleName = 0
	//是否更新
	var isModify = 0
	for _, existConf := range existConfs {
		if existConf.Name == conf.Name && existConf.Uid != conf.Uid {
			doubleName += 1
		}
		if existConf.Uid == conf.Uid {
			isModify += 1
		}
	}
	if doubleName > 0 {
		return "项目名称重复，请修改项目名称。"
	}
	if isModify >= 1 {
		for i, existConf := range existConfs {
			if existConf.Name == conf.Name {
				existConfs[i] = conf
			}
		}
	} else {
		existConfs = append(existConfs, conf)
	}
	viper.Set("conf", existConfs)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
		return "操作失败"
	} else {
		return "操作成功"
	}
}
func (s *ThisCopy) DelProjectConfig(project string) string {
	str := []byte(project)
	conf := Confs{}
	json.Unmarshal(str, &conf)
	existConfs := MyConfig.Conf
	for i, j := 0, len(existConfs); i < j; i = i + 1 {
		if existConfs[i].Uid == conf.Uid {
			existConfs = append(existConfs[:i], existConfs[i+1:]...)
			viper.Set("conf", existConfs)
			err := viper.WriteConfig()
			if err != nil {
				fmt.Println(err)
				return "操作失败"
			} else {
				return "操作成功"
			}
		}
	}
	return "操作成功"
}
