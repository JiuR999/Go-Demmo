package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	name, _ := os.Hostname()
	fmt.Println(name)

	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}
