package routers

import (
	ctl "AndroidToolServer-Go/apps/asset/androidtool/controller"
	"github.com/gin-gonic/gin"
)

func articleApi(group *gin.RouterGroup) {
	//articleApi := r.Group("/article")
	group.GET("/:id", ctl.ArticleController.GetById)
	group.GET("", ctl.ArticleController.GetByChannelName)
	group.POST("", ctl.ArticleController.Insert)
	group.PUT("/:id", ctl.ArticleController.Update)
	group.DELETE("/:id", ctl.ArticleController.DeleteById)
}

func hardwareApi(group *gin.RouterGroup) {
	group.GET("/page", ctl.HardwareController.Page)
	//group.GET("", ac.HardwareController.Insert)
}

func generalFaluireApi(group *gin.RouterGroup) {
	group.GET("/list", ctl.GeneralFailureController.List)
	group.GET("/page", ctl.GeneralFailureController.Page)
	group.GET("/:id", ctl.GeneralFailureController.GetByID)
}

func downloadApi(group *gin.RouterGroup) {
	group.GET("", ctl.DownloadController.DownApk)
}

func zdDreamApi(group *gin.RouterGroup) {
	group.GET("/list", ctl.ZdDreamController.List)
}
