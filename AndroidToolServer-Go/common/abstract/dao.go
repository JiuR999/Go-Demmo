package abstract

import (
	database "AndroidToolServer-Go/roof/db"
	"gorm.io/gorm"
)

type Dao struct {
	Gm    *gorm.DB
	Model any
}

func (dao *Dao) Init() {
	dao.Gm = database.PostgreSQLOrm.DB()
}

func (dao *Dao) GetById(id int64, record any) {
	dao.Gm.Model(dao.Model).Where("id=?", id).Take(record)
}
