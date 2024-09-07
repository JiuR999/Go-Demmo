package main

import (
	"AndroidToolServer-Go/common/utils"
	"AndroidToolServer-Go/model"
	"AndroidToolServer-Go/roof/env"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dadabaseInfo struct {
	DBName     string
	DBUserName string
	DBPassword string
	DBPort     string
	DBHost     string
}

func main() {

	var dream model.ZbDream
	MysqlConfig := createMysqlConfig()
	mysql, _ := createDB(mysql.Open(MysqlConfig))
	var dreams []model.ZbDream
	mysql.Find(&dreams)

	//dsn := "user=postgres password=123456 dbname=dgth_test host=192.168.200.130 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := utils.CreatePostgreSQLDB()
	if err != nil {
		panic("failed to connect database")
	}
	MigrateTable(&dream, db, dreams)
}

func createDB(dialector gorm.Dialector, options ...gorm.Option) (db *gorm.DB, err error) {
	return gorm.Open(dialector, options...)
}

func createMysqlConfig() string {
	mysqlInfo := dadabaseInfo{
		DBName:     env.Config.DBName,
		DBUserName: env.Config.DBUserName,
		DBPort:     env.Config.DBPort,
		DBHost:     env.Config.DBHost,
		DBPassword: env.Config.DBPassword,
	}
	MysqlConfig := mysqlInfo.DBName + ":" + mysqlInfo.DBPassword + "@tcp(" + mysqlInfo.DBHost + ":" + mysqlInfo.DBPort + ")/" + mysqlInfo.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return MysqlConfig
}

const BATCH_SIZE = 100

/*
*
table 需要迁移的表结构
db 新的数据库
datas 迁移数据
*/
func MigrateTable(table interface{}, db *gorm.DB, datas interface{}) {
	db.AutoMigrate(table)
	result := db.CreateInBatches(datas, BATCH_SIZE)
	if result.Error != nil {
		fmt.Println("插入失败", result.Error)
	}
}
