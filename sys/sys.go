package sys

import (
	"bufio"
	"fmt"
	"github.com/wailsapp/wails"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
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
	return nil
}

func (s *Stats) GetOuts(k string) *CPUUsage {
	return &CPUUsage{
		Average: k,
	}
}
func (s *Stats) GetCom(projectname string) {
	//var selectedProject Confs
	//for _, conf := range projectConfs {
	//	if conf.Name == projectname {
	//		selectedProject = conf
	//		break
	//	}
	//}

	s.CmdAndChangeDirToShow("", "", nil)
}
func (s *Stats) CmdAndChangeDirToShow(dir string, commandName string, params []string) error {
	//cmd := exec.Command("cmd.exe", "/c", "cd D:\\1-WorkSpace\\0_SvnProject\\bdcdj && d: && dir")
	ePath, _ := os.Executable()
	fmt.Println(ePath)
	runningPath := path.Dir(ePath)
	command := ` ` + runningPath + `/resource/install.sh ` + "/Users/ghoul/1-svnWork/bdcdj" + ` .`
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
