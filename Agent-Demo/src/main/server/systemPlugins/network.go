package systemPlugins

import (
	"Agent-Demo/src/main/common"
	"fmt"
	"github.com/gin-gonic/gin"
	gnet "github.com/shirou/gopsutil/v3/net"
	"net"
	"net/http"
)

/*
*
获取IP地址列表
*/
func GetNetWorkInfo() ([]string, error) {
	var addrestr []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {

		return addrestr, err
	}
	for _, address := range addrs {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				addrestr = append(addrestr, ipNet.IP.String())
			}
		}
	}
	return addrestr, nil
}

/*
*
获取网卡流量信息
*/
func GetIOCounter() {
	counters, _ := gnet.IOCounters(false)
	fmt.Println(len(counters))
	for _, v := range counters {
		fmt.Println(v)
	}
}

func GetNetInterfaInfo(ctx *gin.Context) {
	interfaces, err := gnet.Interfaces()
	if err != nil {
		ctx.JSON(http.StatusOK, common.NewErrorResp(err.Error()))
	}
	ctx.JSON(http.StatusOK, common.NewSuccessResp(interfaces))
}
