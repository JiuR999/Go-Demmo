package models

import (
	"time"
)

const TableNameHardWare = "hard_ware"

// HardWare mapped from table <hard_ware>
type HardWare struct {
	ID             int32     `gorm:"column:id;type:int;primaryKey" json:"id"`
	Name           string    `gorm:"column:name;type:varchar(16)" json:"name"`
	Code           string    `gorm:"column:code;type:varchar(16)" json:"code"`
	CreateTime     time.Time `gorm:"column:create_time;type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime     time.Time `gorm:"column:update_time;type:datetime" json:"updateTime"`
	MonitorStatus  *int32    `gorm:"column:monitor_status;type:int" json:"monitorStatus"`
	Location       *string   `gorm:"column:location;type:varchar(15)" json:"location"`
	HardWareType   *string   `gorm:"column:hard_ware_type;type:varchar(10)" json:"hardWareType"`
	U              *int32    `gorm:"column:u;type:int" json:"u"`
	UserID         *string   `gorm:"column:user_id;type:mediumtext" json:"userId,string"`
	UserName       *string   `gorm:"column:user_name;type:varchar(10)" json:"userName"`
	UnitDepartID   *string   `gorm:"column:unit_depart_id;type:mediumtext" json:"unitDepartId,string"`
	UnitDepartName *string   `gorm:"column:unit_depart_name;type:varchar(10)" json:"unitDepartName"`
	UnitGroupID    *string   `gorm:"column:unit_group_id;type:varchar(10)" json:"unitGroupId,string"`
	UnitGroupName  *string   `gorm:"column:unit_group_name;type:varchar(10)" json:"unitGroupName"`
}

// TableName HardWare's table name
func (*HardWare) TableName() string {
	return TableNameHardWare
}
