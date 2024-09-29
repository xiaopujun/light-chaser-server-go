package database

import "github.com/light-chaser/server/model/database/base"

type DbRemoteComponent struct {
	base.ORMBase
	ComponentName string `gorm:"component_name"`
	Url           string `gorm:"url"`
}

func (DbRemoteComponent) TableName() string {
	return "remote_component"
}
