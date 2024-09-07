package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Movie struct {
	Top           int    `json:"top"`
	MovieName     string `json:"movieName"`
	ReleaseInfo   string `json:"releaseInfo"`
	SumBoxDesc    string `json:"sumBoxDesc"`
	BoxRate       string `json:"boxRate"`
	ShowCount     int    `json:"showCount"`
	ShowCountRate string `json:"showCountRate"`
	AvgShowView   string `json:"avgShowView"`
	AvgSeatView   string `json:"avgSeatView"`
}

func main() {
	resp, err := http.Get("https://api.pearktrue.cn/api/lsjt/?type=json")
	if err == nil {
		res, _ := io.ReadAll(resp.Body)
		var data map[string]interface{}
		json.Unmarshal(res, &data)
		inters := data["data"].([]interface{})
		for _, v := range inters {
			fmt.Println(v.(string))
		}
	}
}
