package server

import (
	"Agent-Demo/src/main/common"
	"Agent-Demo/src/main/models"
	"Agent-Demo/src/main/server/systemPlugins"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"strings"
)

func InitCollectOnStart() {
	consumer, err := sarama.NewConsumer([]string{"192.168.200.130:9092"}, nil)
	if err != nil {
		fmt.Println("Failed to start consumer: %s", err)
		return
	}
	partition, err := consumer.ConsumePartition(common.Config.RECEIVE_CMD_TOPIC, 0, sarama.OffsetNewest) //获得该topic所有的分区
	if err != nil {
		fmt.Println("Failed to get the list of partition:, ", err)
		return
	}
	defer func() {
		partition.Close()
	}()

	for {
		select {
		case msg := <-partition.Messages():
			fmt.Println("收到命令消息:::", string(msg.Value))
			var systemSyncCmd models.SystemCyncCmd
			json.Unmarshal(msg.Value, &systemSyncCmd)
			fmt.Println(systemSyncCmd.Tick)
			if isMatched(msg) {
				fmt.Println("是本机的命令 执行采集", systemSyncCmd.ParamNum, "的命令")
				if systemSyncCmd.ParamNum >= 100 && systemSyncCmd.ParamNum < 200 {
					ExecuteDisk(systemSyncCmd)
				}
				if systemSyncCmd.ParamNum >= 200 && systemSyncCmd.ParamNum < 300 {
					ExecuteNetwork(systemSyncCmd)
				}
			}
		}
	}
}

func isMatched(msg *sarama.ConsumerMessage) bool {
	var systemSyncCmd models.SystemCyncCmd
	json.Unmarshal(msg.Value, &systemSyncCmd)

	if systemSyncCmd.Unique != "-1" && strings.TrimSpace(systemSyncCmd.Unique) != "" &&
		systemSyncCmd.Unique == "123456" {
		//匹配成功
		return true
	} else if strings.TrimSpace(systemSyncCmd.IP) != "" {
		address, _ := systemPlugins.GetNetWorkInfo()
		for _, add := range address {
			if add == systemSyncCmd.IP {
				return true
			}
		}
	}
	return false
}
