package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	host := "192.168.200.130"
	command := exec.Command("ping", host)
	output, err := command.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	if strings.Contains(string(output), "来自 "+host+" 的回复") {
		fmt.Println("连接成功")
	} else {
		fmt.Println("连接失败")
	}
}
