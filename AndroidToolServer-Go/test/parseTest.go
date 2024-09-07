package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	hex := "0x00000073"
	decimal, _ := strconv.ParseInt(hex[2:], 16, 64)
	fmt.Println(decimal)
	//GetPhone("http://www.zhaohaowang.com")
}

// 爬手机号
func GetPhone(url string) {
	pageStr := GetPageStr(url)
	rePhone := `1[3456789]\d\s?\d{4}\s?\d{4}`
	re := regexp.MustCompile(rePhone)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

// 抽取根据url获取内容
func GetPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	// 字节转字符串
	pageStr = string(pageBytes)
	return pageStr
}

// 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
