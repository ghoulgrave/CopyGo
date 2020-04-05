package sys

import (
	"bufio"
	"copy/docx"
	"copy/logger"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

// Stats .
type Stats struct {
	log     *wails.CustomLogger
	runtime *wails.Runtime
}

// CPUUsage .
type CPUUsage struct {
	Average string `json:"avg"`
}

// WailsInit .
func (s *Stats) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Stats")
	s.runtime = runtime
	runtime.Events.Emit("cpu_usage", s.GetOuts("start"))
	runtime.Events.Emit("builds_pl", s.GetOuts("start"))
	return nil
}

func (s *Stats) GetOuts(k string) *CPUUsage {
	return &CPUUsage{
		Average: k,
	}
}
func (s *Stats) GetCom(projectname string, infos string, isbuild bool) {

	var selectedProject Confs
	for _, conf := range ProjectConfs {
		if conf.Name == projectname {
			selectedProject = conf
			break
		}
	}
	if selectedProject.Dir_path != "" {
		if isbuild {
			s.CmdAndChangeDirToShow(selectedProject.Dir_path, "", nil)
		}
	} else {
		s.log.Info("[COPY-INFO] 项目：【" + projectname + "】 文件夹地址没有配置")
		return
	}

	//fmt.Println(infos)
	var ss []SvnInfo
	json.Unmarshal([]byte(infos), &ss)
	//fmt.Println("xxxxxx")
	//fmt.Println(ss[1].Path)

	//k := "[{\"name\":\"zhangyiyang\",\"time\":\"2020-04-03 09:21:31 +0800 (Fri, 03 Apr 2020)\",\"version\":\"r254129\",\"path\":\" /bdcdj/branches/bdcdj_dbqy/src/main/resources/conf/bdcdj-mybatis/BdcZm.xml\",\"sublogs\":\"\"},{\"name\":\"chenchunxue\",\"time\":\"2020-04-03 10:04:05 +0800 (Fri, 03 Apr 2020)\",\"version\":\"r254139\",\"path\":\" /bdcdj/branches/bdcdj_dbqy/src/main/java/cn/gtmap/bdcdj/utils/Constants.java\",\"sublogs\":\"\"}]"
	//var ss []SvnInfo
	//json.Unmarshal([]byte(k), &ss)
	////fmt.Println("xxxxxx")
	////fmt.Println(len(ss))
	dateNow, _ := s.Copyfiles(selectedProject, ss)

	fmt.Println(dateNow)
	//fmt.Println(fileDirPath)
	//fmt.Println(iNum)
	//复制jar文件

	//复制docx并替换
	pathSeparator := string(os.PathSeparator)
	r, err := docx.ReadDocxFile(RunningPath + pathSeparator + "resource" + pathSeparator + "template.docx")
	if err != nil {
		panic(err)
	}
	docx1 := r.Editable()
	docx1.Replace("zyy_bxr", MyConfig.Username, -1)
	docx1.Replace("zyysj_bxr", time.Now().Format("2006-01-02"), -1)
	docx1.Replace("logs", SubLogs, -1)
	docx1.Replace("Zyy_xqbh", RequsNum, -1)
	docx1.Replace("Zyy_subSvnPath", selectedProject.Sub_path+"/"+dateNow, -1)
	docx1.WriteToFile(selectedProject.Out_path + pathSeparator + dateNow + pathSeparator + "更新说明文档-" + dateNow + ".docx")
	r.Close()

	logger.Info("fffff")

}
func (s *Stats) Copyfiles(projectConf Confs, checkedInfos []SvnInfo) (dateNow string, err error) {
	//获取项目路径
	pathSeparator := string(os.PathSeparator)
	baseUrl := projectConf.Dir_path
	baseUrl = baseUrl + pathSeparator + "target"
	//fmt.Println(baseUrl)
	var file os.FileInfo
	var fileType string
	var iNum = int64(0)
	files, _ := ioutil.ReadDir(baseUrl)
	if len(files) > 0 {
		for i := range files {
			//fmt.Println(files[i].Name())
			if strings.HasSuffix(files[i].Name(), ".jar") && !strings.Contains(files[i].Name(), "sources") {
				file = files[i]
				fileType = "jar"
				break
			}
			if strings.HasSuffix(files[i].Name(), ".war") {
				file = files[i]
				fileType = "war"
				break
			}
		}
	} else {
		fmt.Println("无文件")
		//return
	}
	if file == nil {
		fmt.Println("没找到jar 或 war 文件")
		//return
	}

	fileDirPath := strings.ReplaceAll(file.Name(), "."+fileType, "")
	dirExist, fileDir, _ := PathExists(baseUrl + pathSeparator + fileDirPath)
	if !dirExist {
		fmt.Println("文件夹不存在")
		//return
	}
	if !fileDir.IsDir() {
		fmt.Println("找到的不是文件夹")
		//return
	}
	dateNow = time.Now().Format("20060102150405")
	for _, info := range checkedInfos {
		stringTemp := strings.Split(info.Path, pathSeparator)
		var _fileurl string
		var flag = true
		var _web string
		//没有扩展名的文件放弃
		lastStr := stringTemp[len(stringTemp)-1]
		k := strings.LastIndex(lastStr, ".")
		if k < 0 {
			continue
		}

		for _, s2 := range stringTemp {
			if flag {
				if "web" == strings.ToLower(s2) || "webapp" == strings.ToLower(s2) {
					_web = s2
					_fileurl += s2
					flag = !flag
				} else if "src" == s2 {
					_fileurl += s2
					flag = !flag
				} else {
				}
			} else {
				_fileurl += pathSeparator + s2
			}
		}
		_exten := lastStr[k+1:]
		//fmt.Println(_exten)
		if _exten == "java" || _exten == "groovy" {
			_exten = "class"
		}
		//fmt.Println(_exten)
		if fileType == "jar" {
			_fileurl = strings.Replace(_fileurl, "src"+pathSeparator+"main"+pathSeparator+"java", "classes", -1)
			_fileurl = strings.Replace(_fileurl, "src"+pathSeparator+"main"+pathSeparator+"resources", "classes", -1)
		} else {
			_fileurl = strings.Replace(_fileurl, "src"+pathSeparator+"main"+pathSeparator+"java", "WEB-INF"+pathSeparator+"classes", -1)
			_fileurl = strings.Replace(_fileurl, "src"+pathSeparator+"main"+pathSeparator+"resources", "WEB-INF"+pathSeparator+"classes", -1)
		}
		if _web != "" {
			_fileurl = strings.Replace(_fileurl, _web+pathSeparator, "", -1)
		}

		uFrom := baseUrl + pathSeparator + fileDirPath + pathSeparator + _fileurl
		if fileType == "jar" {
			uFrom = baseUrl + "\\" + _fileurl
		}
		//fmt.Println(uFrom)
		uTo := projectConf.Out_path + pathSeparator + dateNow + pathSeparator + fileDirPath + pathSeparator + _fileurl
		if fileType == "jar" {
			uTo = projectConf.Out_path + pathSeparator + dateNow + pathSeparator + fileDirPath + pathSeparator + "BOOT-INF" + pathSeparator + _fileurl
		}
		os.MkdirAll(strings.Replace(uTo, lastStr, "", -1), 0777)
		//fmt.Println(uTo)
		pFrom := strings.Replace(uFrom, pathSeparator+lastStr, "", -1)
		//fmt.Println(pFrom)
		_, err := os.Stat(pFrom)
		if err == nil {
			if _exten == "class" {
				fileNameTemp := lastStr[:strings.LastIndex(lastStr, ".")]
				filesTemp, _ := ioutil.ReadDir(pFrom)
				for _, fileInfo := range filesTemp {
					if strings.HasPrefix(fileInfo.Name(), fileNameTemp) && strings.HasSuffix(fileInfo.Name(), "."+_exten) {
						//fmt.Println("mmm:" + fileInfo.Name())
						//fmt.Println("copy to : " + strings.Replace(uTo, pathSeparator+lastStr, "", -1) + pathSeparator + fileInfo.Name())
						//fmt.Println("copy ffrom : ", pFrom+pathSeparator+fileInfo.Name())
						CopyFile(strings.Replace(uTo, pathSeparator+lastStr, "", -1)+pathSeparator+fileInfo.Name(), pFrom+pathSeparator+fileInfo.Name())
						iNum = iNum + 1
					}
				}
			} else {
				CopyFile(uTo, uFrom)
				iNum = iNum + 1
			}
		}
	}
	return dateNow, nil
}

func (s *Stats) CmdAndChangeDirToShow(dir string, commandName string, params []string) error {
	//cmd := exec.Command("cmd.exe", "/c", "cd D:\\1-WorkSpace\\0_SvnProject\\bdcdj && d: && dir")
	ePath, _ := os.Executable()
	fmt.Println(ePath)
	runningPath := path.Dir(ePath)
	command := ` ` + runningPath + `/resource/install.sh ` + dir + ` .`
	cmd := exec.Command("/bin/bash", "-c", command)

	//cmd := exec.Command("/bin/bash", "-c", "ls")
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

		if line != "" {
			fmt.Println(line)
			s.runtime.Events.Emit("cpu_usage", s.GetOuts(strings.Replace(line, "\n", "", -1)))
		}

	}
	err = cmd.Wait()
	s.runtime.Events.Emit("cpu_usage", s.GetOuts("[COPY-INFO] All Done ."))
	return err
}

func PathExists(path string) (bool, os.FileInfo, error) {
	file, err := os.Stat(path)
	if err == nil {
		return true, file, nil
	}
	if os.IsNotExist(err) {
		return false, nil, nil
	}
	return false, nil, err
}
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
