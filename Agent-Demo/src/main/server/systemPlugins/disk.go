package systemPlugins

import (
	"Agent-Demo/src/main/common"
	"Agent-Demo/src/main/models"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/disk"
	"net/http"
)

/*
*
获取磁盘分区信息
*/
func Partitions() (models.RespModel, error) {
	var res models.RespModel
	stats, err := disk.Partitions(true)
	if err != nil {
		return res, err
	}

	return common.NewSuccessResp(stats), nil
}

/*
*
获取磁盘流量信息
*/
func GetDiskFlowInfo(ctx *gin.Context) {
	counters, _ := disk.IOCounters()
	ctx.JSON(http.StatusOK, common.NewSuccessResp(counters))
}

/*
*
获取磁盘分区使用信息
*/
func GetDiskUsageInfo(ctx *gin.Context) {
	partitions, err := disk.Partitions(true)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
	}
	diskUsageStats := make([]disk.UsageStat, len(partitions))

	for i, part := range partitions {
		usage, err := disk.Usage(part.Mountpoint)
		if err != nil {
			ctx.JSON(http.StatusOK, err.Error())
		}
		diskUsageStats[i] = *usage
	}
	ctx.JSON(http.StatusOK, common.NewSuccessResp(diskUsageStats))
}
