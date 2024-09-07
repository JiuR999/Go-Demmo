package controller

import (
	"AndroidToolServer-Go/apps/asset/androidtool/models"
	"AndroidToolServer-Go/apps/asset/androidtool/service"
	"AndroidToolServer-Go/common"
	cm "AndroidToolServer-Go/common/models"
	"AndroidToolServer-Go/model"
	"AndroidToolServer-Go/roof/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var GeneralFailureController *generalFailureController

type generalFailureController struct{}

func init() {
	GeneralFailureController = &generalFailureController{}
}

// @Summary		查询所有通用故障
// @title			Swagger API
// @version		1.0
// @Tags			Asset-GeneralFailure
// @description	查询所有通用故障
// @Produce		json
//
// @Success		200	{object}	nil
//
// @router			/failure/list [Get]
func (g *generalFailureController) List(context *gin.Context) {
	var failures []model.DgsGeneralFailure
	db.PostgreSQLOrm.DB().Find(&failures)
	context.JSON(http.StatusOK, models.Result{
		Code: 1,
		Data: failures,
	})
}

// @Summary		分页查询通用故障
// @title			Swagger API
// @version		1.0
// @Tags			Asset-GeneralFailure
// @description	分页查询通用故障
// @Produce		json
//
// @Param			pageSize	query		int	false	"分页大小"
// @Param			page		query		int	false	"页码"
//
// @Success		200			{object}	nil
//
// @router			/failure/page [Get]
func (g *generalFailureController) Page(context *gin.Context) {

	page, _ := strconv.Atoi(context.Query(common.QUERY_KEY_PAGE))
	pageSize, _ := strconv.Atoi(context.Query(common.QUERY_KEY_PAGESIZE))
	cm.InitPageIfAbsent(&page, &pageSize)
	failures, err := service.GetGeneralFailureService().Page(page, pageSize)
	if err != nil {
		context.JSON(http.StatusOK, models.Result{
			Code: 0,
			Data: err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, models.Result{
			Code: 1,
			Data: failures,
		})
	}
}

// @Summary		根据ID查询故障
// @title			Swagger API
// @version		1.0
// @Tags			Asset-GeneralFailure
// @description	根据ID查询故障
// @Produce		json
//
// @Param			id	path		int	false	"故障ID"
//
// @Success		200	{object}	nil
//
// @router			/failure/{id} [Get]
func (g generalFailureController) GetByID(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	failure := service.GetGeneralFailureService().GetById(id)
	context.JSON(http.StatusOK, models.Result{
		Code: 1,
		Data: failure,
	})
}
