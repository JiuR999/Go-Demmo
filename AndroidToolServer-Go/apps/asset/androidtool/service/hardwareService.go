package service

import (
	"AndroidToolServer-Go/apps/asset/androidtool/models"
	"AndroidToolServer-Go/common"
	"AndroidToolServer-Go/roof/db"
	"fmt"
)

var hardwareServiceSingleton = new(hardwareService)

type hardwareService struct{}

func (s hardwareService) Page(page, pageSize int) []models.HardWare {
	var hardwares []models.HardWare
	if pageSize == 0 {
		pageSize = common.PAGE_SIZE
	}
	offset := (page - 1) * pageSize
	totals := db.MysqlOrm.DB().Limit(pageSize).Offset(offset).Find(&hardwares).RowsAffected
	fmt.Println("共查询得到%d条记录", totals)
	return hardwares
}

func GetHardwareService() *hardwareService {
	return hardwareServiceSingleton
}
