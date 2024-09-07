package main

import (
	"AndroidToolServer-Go/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"io"
	"log"
	"strconv"
)

type Document struct {
	Id        int
	Data      any
	IndexName string
}

func main() {
	config := elasticsearch.Config{
		Addresses: []string{
			"http://192.168.200.130:9200",
		},
	}

	client, err := elasticsearch.NewClient(config)
	if err != nil {
		log.Fatal("初始化es客户端失败")
	}

	/*failures, _ := dao.GetGeneralFailureDao().Page(1, 10)
	indexName := "alarm"
	for i, v := range failures {
		id := *v.ID
		document := Document{
			Id:        int(id),
			Data:      v,
			IndexName: indexName,
		}

		addDocument(client, document)
		fmt.Println("成功插入第", i, "条")
	}*/

	getDocument(client, "2")

}
func addDocument(client *elasticsearch.Client, document Document) {
	body := &bytes.Buffer{}
	_ = json.NewEncoder(body).Encode(document.Data)
	_, err := client.Create("alarm", strconv.Itoa(document.Id),
		body)
	if err != nil {
		fmt.Println("插入文档失败:", err)
	}

}

// getDocument 获取文档
func getDocument(client *elasticsearch.Client, id string) {
	resp, err := client.Get("alarm", id)
	if err != nil {
		fmt.Printf("get document by id failed, err:%v\n", err)
		return
	}
	res, err := io.ReadAll(resp.Body)
	var data model.DgsGeneralFailure
	json.Unmarshal(res, &data)
	fmt.Println(data)
}

func str2Ptr(str string) *string {
	return &str
}
