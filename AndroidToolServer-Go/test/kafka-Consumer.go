package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"sync"
)

// kafka consumer
var wg sync.WaitGroup

func main() {
	consumer, err := sarama.NewConsumer([]string{"192.168.200.130:9092"}, nil)
	defer consumer.Close()
	if err != nil {
		fmt.Println("Failed to start consumer: %s", err)
		return
	}
	_, err = consumer.Partitions("web_log") //获得该topic所有的分区
	if err != nil {
		fmt.Println("Failed to get the list of partition:, ", err)
		return
	}

	Subscription(consumer, "agent")

	wg.Wait()

}

func Subscription(consumer sarama.Consumer, topic string) {
	//获得该topic所有的分区
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println("Failed to get the list of partition:, ", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		wg.Add(1)
		go func(sarama.PartitionConsumer) { //为每个分区开一个go协程去取值
			for msg := range pc.Messages() { //阻塞直到有值发送过来，然后再继续等待
				value := string(msg.Value)
				fmt.Println("Receive Message: ", value)
			}
			defer pc.AsyncClose()
			wg.Done()
		}(pc)
	}
}

func start1() {
	consumer, err := sarama.NewConsumer([]string{"192.168.200.130:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("web_log") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	wg.Add(1)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
			wg.Done()
		}(pc)
	}

	wg.Wait()
}
