package models

type Article struct {
	Id          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string `gorm:"column:title;type:character varying(20)" json:"title"`
	ChannelName string `gorm:"column:channel_name;type:character varying(10)" json:"channelName"`
}
