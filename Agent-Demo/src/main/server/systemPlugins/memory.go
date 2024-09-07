package systemPlugins

import (
	"Agent-Demo/src/main/models"
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
)

const B = 1024

func ExecuteMemory(req models.CollectReq) {
	switch req.ParamNum {
	//总内存
	case "200":
		TotalMemory()
	//使用量
	case "202":
		UsedMemory()
	//空闲内存
	case "203":
		FreeMemory()
	}
}
func TotalMemory() int {
	memory, err := mem.VirtualMemory()
	if err == nil {
		fmt.Println("总内存:", memory.Total/B/B/B, "GB")
		return int(memory.Total)
	}
	return 0
}

func UsedMemory() int {
	memory, err := mem.VirtualMemory()
	if err == nil {
		fmt.Println("已经使用:", memory.Used/B/B/B, "GB")
		return int(memory.Used)
	}
	return 0
}

func FreeMemory() int {
	memory, err := mem.VirtualMemory()
	if err == nil {
		fmt.Println("空闲内存:", memory.Free)
		return int(memory.Free)
	}
	return 0
}

func UsedPersent() {
	memory, err := mem.VirtualMemory()
	if err == nil {
		fmt.Println("内存使用率:", memory.UsedPercent, "%")

	}
}
