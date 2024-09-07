package systemPlugins

import (
	"Agent-Demo/src/main/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"net/http"
	"time"
)

func GetCpuInfo(ct *gin.Context) {
	info, _ := cpu.Info()
	ct.JSON(http.StatusOK, common.NewSuccessResp(info))
}

/*
*
获取CPU详细信息
*/
func GetCpuStruct(ctx *gin.Context) {
	info, _ := cpu.Info()
	ctx.JSON(http.StatusOK, common.NewSuccessResp(info))
}

func GetCpuPercent(ctx *gin.Context) {
	tick := 10
	//false 总负载使用率
	percent, _ := cpu.Percent(time.Duration(tick)*time.Second, false)
	s := fmt.Sprintf("%.2f%%", percent[0])
	fmt.Println(s)
	ctx.JSON(http.StatusOK, common.NewSuccessResp(percent))
}
