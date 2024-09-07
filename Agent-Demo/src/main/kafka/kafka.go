package kafka

import (
	"Agent-Demo/src/main/models"
	"Agent-Demo/src/main/server/systemPlugins"
	"Agent-Demo/src/main/service"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

var client sarama.SyncProducer

func InitKafka() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	// 连接kafka
	producer, err := sarama.NewSyncProducer([]string{"192.168.200.130:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	client = producer
}

func main() {

	//60s上传一次 网卡信息
	service.AddReportCronTask("1", 60*time.Second, func() {
		address, err := systemPlugins.GetNetWorkInfo()
		if err == nil {
			r := &models.RespModel{
				Val:  address,
				Time: time.Now().String(),
			}
			marshal, err := json.Marshal(r)
			if err == nil {
				Send2Topic("agent", string(marshal))
			}
		}
	})

	select {}
}

func ExecuteNetWork(client sarama.SyncProducer) {
	address, err := systemPlugins.GetNetWorkInfo()
	if err == nil {
		r := &models.RespModel{
			Val:  address,
			Time: time.Now().String(),
		}
		marshal, err := json.Marshal(r)
		if err == nil {
			Send2Topic("agent", string(marshal))
		}
	}
}
func createMsg(topic, value string) *sarama.ProducerMessage {
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(value)
	return msg
}
func Send2Topic(topic, value string) {
	msg := createMsg(topic, value)
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("消费者发送消息,%v,pid:%v offset:%v\n", msg.Value, pid, offset)
}
