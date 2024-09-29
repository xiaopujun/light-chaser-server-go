package base

import "time"

type ORMBaseModel struct {
	Id         int64     `gorm:"column:id;primary_key"`
	Deleted    int       `gorm:"column:deleted"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}
