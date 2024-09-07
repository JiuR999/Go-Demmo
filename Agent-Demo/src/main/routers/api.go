package routers

import (
	"Agent-Demo/src/main/server/systemPlugins"
	"github.com/gin-gonic/gin"
)

func networkApi(group *gin.RouterGroup) {
	group.GET("/interface", systemPlugins.GetNetInterfaInfo)
}

func cpuApi(group *gin.RouterGroup) {
	group.GET("/info", systemPlugins.GetCpuInfo)
	group.GET("/struct", systemPlugins.GetCpuStruct)
	group.GET("/persent", systemPlugins.GetCpuPercent)
}

func diskApi(group *gin.RouterGroup) {
	group.GET("/flow", systemPlugins.GetDiskFlowInfo)
	group.GET("/usage", systemPlugins.GetDiskUsageInfo)
}
