package controller

import (
	"AndroidToolServer-Go/apps/asset/androidtool/models"
	"AndroidToolServer-Go/apps/asset/androidtool/service"
	"AndroidToolServer-Go/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var ArticleController *articleController

type articleController struct {
}

func init() {
	ArticleController = &articleController{}
}

func (articleController *articleController) GetById(c *gin.Context) {
	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)
	var author model.ApAuthor
	author = service.GetArticleService().GetById(id)
	c.JSON(http.StatusOK, models.Success(author))
}

func (articleController *articleController) GetByChannelName(context *gin.Context) {
	channelName := context.Param("channelName")
	fmt.Println("查询文章，ChannelName=", channelName)
	articles, err := service.GetArticleService().GetByChannelName(channelName)
	if err != nil {
		fmt.Println("查询失败！")
	}
	context.JSON(http.StatusOK, models.Success(articles))
}

func (articleController *articleController) Insert(context *gin.Context) {
	var article models.Article

	if err := context.BindJSON(&article); err == nil {
		ret := service.GetArticleService().Insert(&article)
		if ret.RowsAffected > 0 {
			//修改成功
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "修改成功"})
		}
	} else {
		context.JSON(http.StatusOK, gin.H{"err": err.Error()})
	}
}

func (articleController *articleController) Update(context *gin.Context) {
	var article models.Article
	id := context.Param("id")
	article.Id, _ = strconv.Atoi(id)
	if err := context.BindJSON(&article); err == nil {
		service.GetArticleService().Update(&article)
	}
}

func (articleController *articleController) DeleteById(context *gin.Context) {
	var article models.Article
	id := context.Param("id")
	article.Id, _ = strconv.Atoi(id)
	ret := service.GetArticleService().DeleteById(article.Id)
	if ret.RowsAffected > 0 {
		context.JSON(http.StatusOK, models.Success("删除成功"))
	} else {
		context.JSON(http.StatusOK, models.Error("删除成功"))
	}
}
