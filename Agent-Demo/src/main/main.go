package main

import (
	"Agent-Demo/src/main/kafka"
	"Agent-Demo/src/main/routers"
	"Agent-Demo/src/main/server"
	"Agent-Demo/src/main/service"
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	if len(os.Args) > 1 {
		for _, v := range os.Args {
			fmt.Println(v)
		}
	} else {
		fmt.Println("没有传入任何参数")
	}
	wg.Add(1)
	go func() {
		routers.Router.Init().Run(":8080")
	}()
	service.InitTimeWheelOnStart()
	routers.InitRouter()
	wg.Add(1)
	//接收cmd指令
	go func() {
		server.InitCollectOnStart()
	}()
	kafka.InitKafka()
	wg.Wait()
}
