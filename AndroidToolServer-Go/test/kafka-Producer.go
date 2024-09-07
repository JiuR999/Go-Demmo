package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"192.168.200.130:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}

	defer client.Close()

	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		send2Topic(client, createMsg("agent", time.Now().String()))
	}
}

func createMsg(topic, value string) *sarama.ProducerMessage {
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(value)
	return msg
}
func send2Topic(client sarama.SyncProducer, msg *sarama.ProducerMessage) {
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("消费者发送消息,%v,pid:%v offset:%v\n", msg.Value, pid, offset)
}
