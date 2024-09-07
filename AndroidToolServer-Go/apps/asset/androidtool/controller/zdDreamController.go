package controller

import (
	"AndroidToolServer-Go/apps/asset/androidtool/models"
	"AndroidToolServer-Go/apps/asset/androidtool/models/dto"
	"AndroidToolServer-Go/apps/asset/androidtool/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var ZdDreamController *zdDreamController

const (
	BIGLX   = "biglx"
	SMALLLX = "smalllx"
	TITLE   = "title"
)

type zdDreamController struct {
}

// @Summary		根据条件解梦
// @title			Swagger API
// @version		1.0
// @Tags			周公解梦
// @description	根据条件解梦
// @Produce		json
//
// @Param			biglx	query		string	false	"梦的大分类"
// @Param			smalllx	query		string	false	"梦的细分类"
// @Param			title	query		string	false	"梦的主题"
//
// @Success		200		{object}	nil
//
// @router			/dream/list [Get]
func (zc zdDreamController) List(context *gin.Context) {
	var dreamDTO dto.DreamDTO
	if context.Query("biglx") != "" {
		dreamDTO.Biglx = context.Query(BIGLX)
	}
	if context.Query("smalllx") != "" {
		dreamDTO.Smalllx = context.Query(SMALLLX)
	}
	if context.Query("title") != "" {
		dreamDTO.Title = context.Query(TITLE)
	}

	context.JSON(http.StatusOK, models.Result{
		Code: 1,
		Data: service.GetZdDreamService().List(dreamDTO),
	})
}

func init() {
	ZdDreamController = &zdDreamController{}
}
