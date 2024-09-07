package db

import (
	"AndroidToolServer-Go/common/utils"
	"AndroidToolServer-Go/roof/env"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var MysqlOrm *orm
var PostgreSQLOrm *orm

func init() {
	MysqlOrm = NewMySQLOrm()
	postgresql, _ := utils.CreatePostgreSQLDB()
	PostgreSQLOrm = &orm{
		engine: postgresql,
	}
}

func NewMySQLOrm() *orm {
	o := new(orm)
	o.loadMySQLConfig()
	logrus.Info("连接MySQL数据库dgth")
	return o
}

func NewPostGreSQLOrm() *orm {
	o := new(orm)
	o.loadPostGreSQLConfig()
	logrus.Info("连接MySQL数据库dgth")
	return o
}

type orm struct {
	engine *gorm.DB
}

func (db *orm) loadPostGreSQLConfig() {

}
func (db *orm) loadMySQLConfig() {
	var err error
	dbName := env.Config.DBName
	dbUser := env.Config.DBUserName
	dbPassword := env.Config.DBPassword
	//dbHost := env.Config.DBHost
	dbPort := env.Config.DBPort
	MysqlConfig := dbUser + ":" + dbPassword + "@tcp(" + "localhost" + ":" + dbPort + ")/" + string(dbName) + "?charset=utf8mb4&parseTime=True&loc=Local"
	/*gormLogger := logger.New(
		logrus.New(),
		logger.Config{
			SlowThreshold:             500 * time.Millisecond,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
			LogLevel:                  logger.Info,
		},
	)*/
	gormConfig := &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	}

	db.engine, err = gorm.Open(mysql.Open(MysqlConfig), gormConfig)

	if err != nil {
		log.Fatal("Open database error:" + err.Error())
	}

	sqlDB, _ := db.engine.DB()
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(1000)
	sqlDB.SetConnMaxLifetime(60 * time.Second)
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
func (db *orm) DB() *gorm.DB {
	return db.engine
}
