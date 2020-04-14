package sys

import (
	"archive/zip"
	"bufio"
	"copy/docx"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// thisCopy .
type ThisCopy struct {
	log     *wails.CustomLogger
	runtime *wails.Runtime
}

// TimesStruct .
type TimesStruct struct {
	Str string `json:"str"`
}

// WailsInit .
func (s *ThisCopy) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("ThisCopy")
	s.runtime = runtime
	runtime.Events.Emit("cpu_usage", s.GetOuts("start"))
	runtime.Events.Emit("builds_pl", s.GetOuts("start"))
	return nil
}

func (s *ThisCopy) GetOuts(k string) *TimesStruct {
	return &TimesStruct{
		Str: k,
	}
}
func (s *ThisCopy) GetCom(projectname string, infos string, isbuild bool) {
	s.runtime.Events.Emit("cpu_usage", s.GetOuts("[COPY-INFO] Beginning ."))
	var selectedProject Confs
	for _, conf := range ProjectConfs {
		if conf.Name == projectname {
			selectedProject = conf
			break
		}
	}
	if selectedProject.Dir_path != "" {
		if isbuild {
			s.CmdAndChangeDirToShow(selectedProject.Dir_path, false)
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
	dateNow := time.Now().Format("20060102150405")
	_, fileDirPath, _, fileType, _ := s.copyfiles(selectedProject, ss, dateNow)

	//fmt.Println(dateNow)
	//fmt.Println(fileDirPath)
	//fmt.Println(iNum)
	//fmt.Println(fileType)
	//复制jar文件
	if fileType == "jar" {
		s.copyJars(selectedProject, dateNow, fileDirPath)
	}

	ZipDir(selectedProject.Out_path+PathSeparator+dateNow+PathSeparator+fileDirPath, selectedProject.Out_path+PathSeparator+dateNow+PathSeparator+fileDirPath+".zip")
	//复制docx并替换
	r, err := docx.ReadDocxFile(RunningPath + PathSeparator + "resource" + PathSeparator + "template.docx")
	if err != nil {
		panic(err)
	}
	docx1 := r.Editable()
	docx1.Replace("zyy_bxr", MyConfig.Username, -1)
	docx1.Replace("zyysj_bxr", time.Now().Format("2006-01-02"), -1)
	docx1.Replace("logs", SubLogs, -1)
	docx1.Replace("Zyy_xqbh", RequsNum, -1)
	docx1.Replace("Zyy_subSvnPath", selectedProject.Sub_path+"/"+dateNow, -1)
	docx1.WriteToFile(selectedProject.Out_path + PathSeparator + dateNow + PathSeparator + "更新说明文档-" + dateNow + ".docx")
	r.Close()

	s.runtime.Events.Emit("cpu_usage", s.GetOuts("[COPY-INFO] All Done ."))
	s.runtime.Events.Emit("cpu_usage", s.GetOuts("[COPY-INFO] file folder : "+selectedProject.Out_path+PathSeparator+dateNow))

}

/**
1 时间戳文件夹名称
2 war或jar文件名
3 文件数量
4 文件类型
5 错误
*/
//通过获取日志将文件拷贝到对应位置
func (s *ThisCopy) copyfiles(projectConf Confs, checkedInfos []SvnInfo, dateNow string) (string, string, int64, string, error) {
	//获取项目路径
	baseUrl := projectConf.Dir_path
	baseUrl = baseUrl + PathSeparator + "target"
	fmt.Println(baseUrl)
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
	if fileType == "war" {
		dirExist, fileDir, _ := PathExists(baseUrl + PathSeparator + fileDirPath)
		if !dirExist {
			fmt.Println("文件夹不存在")
			//return
		}
		if !fileDir.IsDir() {
			fmt.Println("找到的不是文件夹")
			//return
		}
	}

	for _, info := range checkedInfos {
		stringTemp := strings.Split(info.Path, PathSeparator)
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
				_fileurl += PathSeparator + s2
			}
		}
		_exten := lastStr[k+1:]
		//fmt.Println(_exten)
		if _exten == "java" || _exten == "groovy" {
			_exten = "class"
		}
		//fmt.Println(_exten)
		var jar_web = false
		if fileType == "jar" {
			_fileurl = strings.Replace(_fileurl, "src"+PathSeparator+"main"+PathSeparator+"java", "classes", -1)
			_fileurl = strings.Replace(_fileurl, "src"+PathSeparator+"main"+PathSeparator+"resources", "classes", -1)
			if strings.Contains(_fileurl, "src"+PathSeparator+"main"+PathSeparator+"webapp") {
				_fileurl = strings.Replace(_fileurl, "src"+PathSeparator+"main"+PathSeparator+"webapp", "webapp", -1)
				jar_web = true
			}
		} else {
			_fileurl = strings.Replace(_fileurl, "src"+PathSeparator+"main"+PathSeparator+"java", "WEB-INF"+PathSeparator+"classes", -1)
			_fileurl = strings.Replace(_fileurl, "src"+PathSeparator+"main"+PathSeparator+"resources", "WEB-INF"+PathSeparator+"classes", -1)
		}
		if _web != "" {
			_fileurl = strings.Replace(_fileurl, _web+PathSeparator, "", -1)
		}

		uFrom := baseUrl + PathSeparator + fileDirPath + PathSeparator + _fileurl
		if fileType == "jar" {
			uFrom = baseUrl + PathSeparator + _fileurl
		}
		//fmt.Println(uFrom)
		uTo := projectConf.Out_path + PathSeparator + dateNow + PathSeparator + fileDirPath + PathSeparator + _fileurl
		if fileType == "jar" {
			if jar_web {
				_fileurl = strings.Replace(_fileurl, "webapp"+PathSeparator, "", -1)
				uTo = projectConf.Out_path + PathSeparator + dateNow + PathSeparator + fileDirPath + PathSeparator + "META-INF" + PathSeparator + "resources" + PathSeparator + _fileurl
			} else {
				uTo = projectConf.Out_path + PathSeparator + dateNow + PathSeparator + fileDirPath + PathSeparator + "BOOT-INF" + PathSeparator + _fileurl
			}
		}
		os.MkdirAll(strings.Replace(uTo, lastStr, "", -1), 0777)
		//fmt.Println(uTo)
		pFrom := strings.Replace(uFrom, PathSeparator+lastStr, "", -1)
		//fmt.Println(pFrom)
		_, err := os.Stat(pFrom)
		if err == nil {
			if _exten == "class" {
				fileNameTemp := lastStr[:strings.LastIndex(lastStr, ".")]
				filesTemp, _ := ioutil.ReadDir(pFrom)
				for _, fileInfo := range filesTemp {
					if strings.HasPrefix(fileInfo.Name(), fileNameTemp) && strings.HasSuffix(fileInfo.Name(), "."+_exten) {
						//fmt.Println("mmm:" + fileInfo.Name())
						//fmt.Println("copy to : " + strings.Replace(uTo, PathSeparator+lastStr, "", -1) + PathSeparator + fileInfo.Name())
						//fmt.Println("copy ffrom : ", pFrom+PathSeparator+fileInfo.Name())
						CopyFile(strings.Replace(uTo, PathSeparator+lastStr, "", -1)+PathSeparator+fileInfo.Name(), pFrom+PathSeparator+fileInfo.Name())
						iNum = iNum + 1
					}
				}
			} else {
				CopyFile(uTo, uFrom)
				iNum = iNum + 1
			}
		}
	}
	return dateNow, fileDirPath, iNum, fileType, nil
}

//编译并将信息输出到runtime参数中
func (s *ThisCopy) CmdAndChangeDirToShow(dir string, ispl bool) error {
	//cmd := exec.Command("cmd.exe", "/c", "cd D:\\1-WorkSpace\\0_SvnProject\\bdcdj && d: && dir")
	var cmd *exec.Cmd
	sysType := runtime.GOOS

	var err error
	if sysType == "windows" {
		cmd =  exec.Command("cmd.exe", "/c", "cd "+dir+" && "+dir[0:2]+" && mvn clean && mvn install")
	}else{
		ePath, _ := os.Executable()
		fmt.Println(ePath)
		runningPath := path.Dir(ePath)
		command := ` ` + runningPath + `/resource/install.sh ` + dir + ` .`
		cmd = exec.Command("/bin/bash", "-c", command)
	}

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
			if ispl {
				s.runtime.Events.Emit("builds_pl", s.GetOuts(strings.Replace(line, "\n", "", -1)))
			} else {
				s.runtime.Events.Emit("cpu_usage", s.GetOuts(strings.Replace(line, "\n", "", -1)))
			}

		}

	}
	err = cmd.Wait()

	return err
}

//复制jar包
func (s *ThisCopy) copyJars(projectConf Confs, dateNow string, fileDirPath string) {
	uTo := projectConf.Out_path + PathSeparator + dateNow + PathSeparator + fileDirPath + PathSeparator + "BOOT-INF" + PathSeparator + "lib"
	//创建lib文件夹
	exist, _, _ := PathExists(uTo)
	if !exist {
		os.MkdirAll(uTo, 0777)
	}
	libDir := projectConf.Dir_path + PathSeparator + "target" + PathSeparator + "lib"
	filesTemp, _ := ioutil.ReadDir(libDir)
	jars := MyConfig.Jarnames
	for _, jar := range jars {
		for _, fileInfo := range filesTemp {
			if strings.HasPrefix(fileInfo.Name(), jar) {
				CopyFile(uTo+PathSeparator+fileInfo.Name(), libDir+PathSeparator+fileInfo.Name())
			}
		}
	}
}

//判断文件夹（文件）是否存在
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

//单纯复制文件
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

//压缩文件夹
func ZipDir(srcFile string, destZip string) error {
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()
	archive := zip.NewWriter(zipfile)
	defer archive.Close()
	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = strings.TrimPrefix(path, filepath.Dir(srcFile)+"/")
		// header.Name = path
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})

	return err
}
