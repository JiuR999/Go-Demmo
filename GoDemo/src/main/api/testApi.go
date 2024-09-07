package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	res := make(chan string)
	start := time.Now()
	go func() {
		response, err := http.Get("https://api.pearktrue.cn/api/hitokoto/")
		if err != nil {
			return
		}
		r, err := io.ReadAll(response.Body)
		if err != nil {
			return
		}
		//fmt.Println(string(r))
		res <- string(r)
	}()
	fmt.Println(<-res)
	fmt.Println("耗时", time.Since(start))
}
