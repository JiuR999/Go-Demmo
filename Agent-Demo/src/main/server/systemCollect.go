package server

import (
	"Agent-Demo/src/main/common"
	"Agent-Demo/src/main/kafka"
	"Agent-Demo/src/main/models"
	"Agent-Demo/src/main/server/systemPlugins"
	"Agent-Demo/src/main/service"
	"encoding/json"
	"fmt"
	"time"
)

func ExecuteDisk(cmd models.SystemCyncCmd) {
	//10s上报一次磁盘信息
	service.AddReportCronTask(cmd.Unique, time.Duration(cmd.Tick)*time.Second, func() {
		partitions, err2 := systemPlugins.Partitions()
		fmt.Printf("每%v s获取磁盘分区信息", cmd.Tick)
		if err2 == nil {
			
			marshal, _ := json.Marshal(partitions)
			kafka.Send2Topic("agent", string(marshal))
		}
	})
}

func ExecuteNetwork(cmd models.SystemCyncCmd) {
	//10s上报一次磁盘信息
	service.AddReportCronTask(cmd.Unique, time.Duration(cmd.Tick)*time.Second, func() {
		fmt.Printf("每%v s获取网络信息", cmd.Tick)
		r, err := systemPlugins.GetNetWorkInfo()
		if err == nil {
			res := common.NewSuccessResp(r)
			marshal, _ := json.Marshal(res)
			kafka.Send2Topic("agent", string(marshal))
		}

	})
}
