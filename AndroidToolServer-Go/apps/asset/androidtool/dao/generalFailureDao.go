package dao

import (
	"AndroidToolServer-Go/common/abstract"
	"AndroidToolServer-Go/model"
)

var generalDao = new(generalFailureDao)

type generalFailureDao struct {
	abstract.Dao
}

func init() {
	generalDao.Init()
}

func (d generalFailureDao) Page(page int, size int) (failures []model.DgsGeneralFailure, err error) {
	var datas []model.DgsGeneralFailure
	generalDao.Gm.Limit(size).Offset((page - 1) * size).Find(&datas)
	/*db := generalDao.Gm.Find(&datas)
	if db.Error != nil {
		return nil, db.Error
	}*/
	return datas, nil
}

func init() {
	generalDao.Init()
	generalDao.Model = model.DgsGeneralFailure{}
}

func GetGeneralFailureDao() *generalFailureDao {
	return generalDao
}
