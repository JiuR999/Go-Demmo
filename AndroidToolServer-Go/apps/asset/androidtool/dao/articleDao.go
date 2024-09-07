package dao

import (
	"AndroidToolServer-Go/apps/asset/androidtool/models"
	"AndroidToolServer-Go/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dao = new(articleDao)

type articleDao struct {
	DB *gorm.DB
}

func GetArticleDao() *articleDao {
	return dao
}

func init() {
	username := "root"           //账号
	password := "zx200387"       //密码
	host := "localhost"          //数据库地址
	port := "3306"               //端口
	Dnname := "leadnews_article" //数据库名
	timeout := "10s"             //连接超时，10s
	//root:root@tcp(127.0.0.1:3306)/test？
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=%s&parseTime=true", username, password, host, port, Dnname, timeout)
	//连接mysql，获得DB类型实例，用于后面数据库的读写操作
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("连接数据库失败，error=" + err.Error())
	}
	dao = &articleDao{
		DB: db,
	}

	//连接成功
	fmt.Println("连接数据库" + Dnname)
	if err := db.AutoMigrate(&models.Article{}); err != nil {
		fmt.Println("创建表失败")
	}
}

func (articleDao *articleDao) Insert(a *models.Article) *gorm.DB {
	return dao.DB.Create(a)
}

func (articleDao *articleDao) GetByChannelName(channelName string) ([]models.ApArticle, error) {
	var articles []models.ApArticle
	result := dao.DB.Where("channel_name=?", channelName).Find(&articles)
	if result.Error != nil {
		return nil, result.Error
	}
	return articles, nil
}

func (ad *articleDao) GetById(id int) model.ApAuthor {
	var article models.ApArticle
	dao.DB.Where("id=?", id).Find(&article)
	var author model.ApAuthor
	dao.DB.Where("id=?", article.AuthorID).Find(&author)
	return author
}
func (articleDao *articleDao) DeleteById(id int) *gorm.DB {
	return dao.DB.Delete(&models.Article{Id: id})
}
