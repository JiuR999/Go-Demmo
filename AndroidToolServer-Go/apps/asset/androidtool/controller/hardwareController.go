package controller

import (
	"AndroidToolServer-Go/apps/asset/androidtool/models"
	"AndroidToolServer-Go/apps/asset/androidtool/service"
	"AndroidToolServer-Go/common"
	cm "AndroidToolServer-Go/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var HardwareController *hardwareController

type hardwareController struct{}

func init() {
	HardwareController = &hardwareController{}
}

/*
	func (hardwareController *hardwareController) List(context *gin.Context) {
		var hardwares []models.HardWare
		hardwares = service.GetHardwareService().Page()
		context.JSON(http.StatusOK, models.Result{
			Code: 1,
			Data: hardwares})
	}
*/
/**
分页查询
*/
func (hardwareController *hardwareController) Page(context *gin.Context) {
	var hardwares []models.HardWare

	page, _ := strconv.Atoi(context.Query(common.QUERY_KEY_PAGE))
	pageSize, _ := strconv.Atoi(context.Query(common.QUERY_KEY_PAGESIZE))
	cm.InitPageIfAbsent(&page, &pageSize)
	hardwares = service.GetHardwareService().Page(page, pageSize)
	context.JSON(http.StatusOK, models.Result{
		Code: 1,
		Data: hardwares})
}
