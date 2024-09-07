package service

import (
	"AndroidToolServer-Go/apps/asset/androidtool/models/dto"
	"AndroidToolServer-Go/model"
	"AndroidToolServer-Go/roof/db"
	"fmt"
	"strings"
)

var zdDreamServiceSingleton = new(zdDreamService)

type zdDreamService struct{}

const SELECT_STRING = "select id,title,message,biglx,smalllx from zb_dream where "

func (s *zdDreamService) List(dreamDto dto.DreamDTO) any {
	var datas []model.ZbDream
	// 动态条件构建
	//var conditions = make(map[string]interface{})
	//conditions["biglx"] = dreamDto.Biglx
	var conditions []interface{}
	var sb strings.Builder
	if dreamDto.Biglx != "" {
		conditions = append(conditions, dreamDto.Biglx)
		sb.WriteString("biglx=?")
	}
	if dreamDto.Smalllx != "" {
		conditions = append(conditions, dreamDto.Smalllx)
		if sb.Len() > 1 {
			sb.WriteString(" AND ")
		}
		sb.WriteString("smalllx=?")
	}
	fmt.Println(sb.String())
	fmt.Println(conditions...)
	//var sql = SELECT_STRING + sb.String()
	db.PostgreSQLOrm.DB().Where(sb.String(), conditions...).Find(&datas)
	return datas
}

func GetZdDreamService() *zdDreamService {
	return zdDreamServiceSingleton
}
