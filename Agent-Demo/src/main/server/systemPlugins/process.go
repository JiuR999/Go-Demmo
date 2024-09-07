package systemPlugins

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
)

/*
*
获取系统进程总数
*/
func GetPid() []int32 {
	pids, _ := process.Pids()
	return pids
}

/*
*
获取进程名称 状态
*/
func GetProcessDetail() {
	processes, _ := process.Processes()
	for _, p := range processes {
		name, _ := p.Name()
		isRunning, _ := p.IsRunning()
		exe, _ := p.Exe()
		fmt.Printf("进程名字：%v\t 进程状态：%v %v\t ", name, isRunning, exe)
	}
}
