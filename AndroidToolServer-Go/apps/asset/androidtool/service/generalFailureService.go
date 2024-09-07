package service

import (
	"AndroidToolServer-Go/apps/asset/androidtool/dao"
	"AndroidToolServer-Go/model"
)

type generalFailureService struct{}

var generalFailureServiceSingleton = new(generalFailureService)

func GetGeneralFailureService() *generalFailureService {
	return generalFailureServiceSingleton
}

func (s generalFailureService) Page(page int, size int) (failures []model.DgsGeneralFailure, err error) {
	return dao.GetGeneralFailureDao().Page(page, size)
}

func (s generalFailureService) GetById(id int) (record model.DgsGeneralFailure) {
	failure := model.DgsGeneralFailure{}
	dao.GetGeneralFailureDao().GetById(int64(id), &failure)
	return failure
}
