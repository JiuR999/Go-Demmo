package models

type RespModel struct {
	//1----成功
	//2----失败
	Code int
	Val  any
	Time string
}

type SystemCyncCmd struct {
	Unique   string `json:"id"`
	IP       string `json:"ip"`
	Tick     int    `json:"tick"`
	ParamNum int    `json:"paramNum"`
}

/*
{
"id":"12",
"ip":"192.168.200.130",
"Tick":10
}
*/
type CollectReq struct {
	ParamNum  string `json:"paramNum"`
	Frequency int    `json:"frequency"`
}

type DiskModel struct {
	Device     string   `json:"device"`
	Mountpoint string   `json:"mountpoint"`
	Fstype     string   `json:"fstype"`
	Opts       []string `json:"opts"`
}
