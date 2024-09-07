package utils

import (
	"AndroidToolServer-Go/roof/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreatePostgreSQLDB() (db *gorm.DB, err error) {
	dsn := "user=postgres password=123456 dbname=dgth_test host=192.168.200.130 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func CreateMySQLDB() (db *gorm.DB, err error) {
	dbName := "dgth"
	dbUser := env.Config.DBUserName
	dbPassword := env.Config.DBPassword
	dbHost := env.Config.DBHost
	dbPort := env.Config.DBPort
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
