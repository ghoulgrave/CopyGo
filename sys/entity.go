package sys

//svn 信息列表
type SvnInfo struct {
	Name    string `json:"name"`
	Time    string `json:"time"`
	Version string `json:"version"`
	Path    string `json:"path"`
	SubLogs string `json:"sublogs"`
}

//项目与选择的日志关系
type ProCheckedSvn struct {
	Svns    []SvnInfo `json:"svn"`
	Project Confs     `json:"project"`
}

type MyData struct {
	Code  int64     `json:"code"`
	Msg   string    `json:"msg"`
	Count int64     `json:"count"`
	Data  []SvnInfo `json:"data"`
	Other int64     `json:"-"` // 直接忽略字段
}

//查询条件
type Search struct {
	ConfigNum string
	Kssj      string
	Jssj      string
	Name      string
	Build     string
}

//系统展示配置
type Config struct {
	Username   string
	Searchname string
	PlOutPath  string
	Jarnames   []string
	Conf       []Confs
}

//项目配置
type Confs struct {
	Sub_path string `json:"subpath"`
	Dir_path string `json:"dirpath"`
	Svn_path string `json:"svnpath"`
	Out_path string `json:"outpath"`
	Name     string `json:"name"`
	Uid      string `json:"uid"`
}
