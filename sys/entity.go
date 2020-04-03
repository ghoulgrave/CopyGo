package sys

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
