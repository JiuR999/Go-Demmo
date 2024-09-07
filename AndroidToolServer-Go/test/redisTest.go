package main

import (
	"AndroidToolServer-Go/model"
	"AndroidToolServer-Go/roof/db"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	ctx := context.Background()
	rcli := redis.NewClient(&redis.Options{
		Addr:     "192.168.200.130:6379",
		Password: "",
		DB:       1,
	})

	fmt.Println("Redis开始读取数据......")
	start := time.Now()
	res, _ := rcli.Get(ctx, "Test1").Result()
	fmt.Println(time.Since(start))
	if res == "" {
		fmt.Println("无缓存，读取数据库")
		var data model.DgsGeneralFailure
		start = time.Now()
		db.PostgreSQLOrm.DB().Take(&data)
		fmt.Println(time.Since(start))
		jsonFailure, _ := json.Marshal(data)
		err := rcli.Set(ctx, "Test1", string(jsonFailure), 0).Err()
		if err != nil {
			fmt.Println("插入数据失败:" + err.Error())
		}
	}
	var i model.DgsGeneralFailure

	json.Unmarshal([]byte(res), &i)
	rcli.Publish(ctx, "TestPub", "payload")
	psub := rcli.Subscribe(ctx, "TestPub")
	defer psub.Close()

	for {
		msg, _ := psub.ReceiveMessage(ctx)
		fmt.Println(msg.Channel, msg.Payload)
	}
	//fmt.Println(i.AlarmTitle)
	//testPubSub(rcli)
}

func testPubSub(cli *redis.Client) {
	ctx := context.Background()
	psub := cli.Subscribe(ctx, "TestPub")
	defer psub.Close()

	for {
		msg, _ := psub.ReceiveMessage(ctx)
		fmt.Println(msg.Channel, msg.Payload)
	}
}
