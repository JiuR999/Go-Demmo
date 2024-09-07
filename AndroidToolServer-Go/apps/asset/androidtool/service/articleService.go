package service

import (
	"AndroidToolServer-Go/apps/asset/androidtool/dao"
	"AndroidToolServer-Go/apps/asset/androidtool/models"
	"AndroidToolServer-Go/model"
	"gorm.io/gorm"
)

var articleServiceSingleton = new(articleService)

type articleService struct{}

func GetArticleService() *articleService {
	return articleServiceSingleton
}

func (articleService *articleService) GetByChannelName(channelName string) ([]models.ApArticle, error) {
	return dao.GetArticleDao().GetByChannelName(channelName)
}

func (articleService *articleService) Insert(a *models.Article) *gorm.DB {
	return dao.GetArticleDao().Insert(a)
}

func (articleService *articleService) Update(a *models.Article) *gorm.DB {
	return dao.GetArticleDao().DB.Updates(a)
}

func (articleService *articleService) DeleteById(id int) *gorm.DB {
	return dao.GetArticleDao().DeleteById(id)
}

func (articleService *articleService) GetById(id int) model.ApAuthor {
	return dao.GetArticleDao().GetById(id)
}
