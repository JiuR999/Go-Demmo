package models

import (
	"time"
)

const TableNameApArticle = "ap_article"

// ApArticle 文章信息表，存储已发布的文章
type ApArticle struct {
	ID          int64   `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	Title       *string `gorm:"column:title;type:varchar(50);comment:标题" json:"title"`                        // 标题
	AuthorID    *int32  `gorm:"column:author_id;type:int unsigned;comment:文章作者的ID" json:"authorId,string"`    // 文章作者的ID
	AuthorName  *string `gorm:"column:author_name;type:varchar(20);comment:作者昵称" json:"authorName"`           // 作者昵称
	ChannelID   *int32  `gorm:"column:channel_id;type:int unsigned;comment:文章所属频道ID" json:"channelId,string"` // 文章所属频道ID
	ChannelName *string `gorm:"column:channel_name;type:varchar(10);comment:频道名称" json:"channelName"`         // 频道名称
	/*
		文章布局
		            0 无图文章
		            1 单图文章
		            2 多图文章
	*/
	Layout *int32 `gorm:"column:layout;type:tinyint unsigned;comment:文章布局\n            0 无图文章\n            1 单图文章\n            2 多图文章" json:"layout"`
	/*
		文章标记
		            0 普通文章
		            1 热点文章
		            2 置顶文章
		            3 精品文章
		            4 大V 文章
	*/
	Flag *int32 `gorm:"column:flag;type:tinyint unsigned;comment:文章标记\n            0 普通文章\n            1 热点文章\n            2 置顶文章\n            3 精品文章\n            4 大V 文章" json:"flag"`
	/*
		文章图片
		            多张逗号分隔
	*/
	Images      *string    `gorm:"column:images;type:varchar(1000);comment:文章图片\n            多张逗号分隔" json:"images"`
	Labels      *string    `gorm:"column:labels;type:varchar(500);comment:文章标签最多3个 逗号分隔" json:"labels"`      // 文章标签最多3个 逗号分隔
	Likes       *int32     `gorm:"column:likes;type:int unsigned;comment:点赞数量" json:"likes"`                 // 点赞数量
	Collection  *int32     `gorm:"column:collection;type:int unsigned;comment:收藏数量" json:"collection"`       // 收藏数量
	Comment     *int32     `gorm:"column:comment;type:int unsigned;comment:评论数量" json:"comment"`             // 评论数量
	Views       *int32     `gorm:"column:views;type:int unsigned;comment:阅读数量" json:"views"`                 // 阅读数量
	ProvinceID  *int32     `gorm:"column:province_id;type:int unsigned;comment:省市" json:"provinceId,string"` // 省市
	CityID      *int32     `gorm:"column:city_id;type:int unsigned;comment:市区" json:"cityId,string"`         // 市区
	CountyID    *int32     `gorm:"column:county_id;type:int unsigned;comment:区县" json:"countyId,string"`     // 区县
	CreatedTime *time.Time `gorm:"column:created_time;type:datetime;comment:创建时间" json:"createdTime"`        // 创建时间
	PublishTime *time.Time `gorm:"column:publish_time;type:datetime;comment:发布时间" json:"publishTime"`        // 发布时间
	SyncStatus  *bool      `gorm:"column:sync_status;type:tinyint(1);comment:同步状态" json:"syncStatus"`        // 同步状态
	Origin      *int32     `gorm:"column:origin;type:tinyint unsigned;comment:来源" json:"origin"`             // 来源
	StaticURL   *string    `gorm:"column:static_url;type:varchar(150)" json:"staticUrl"`
}

// TableName ApArticle's table name
func (*ApArticle) TableName() string {
	return TableNameApArticle
}
