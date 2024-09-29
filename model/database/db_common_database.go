package database

import "github.com/light-chaser/server/model/database/base"

type DbCommonDatabase struct {
	base.ORMBase
	Name         string `gorm:"name"`
	Type         string `gorm:"type"`
	Ip           string `gorm:"ip"`
	Port         string `gorm:"port"`
	DatabaseName string `gorm:"database_name"`
	UserName     string `gorm:"user_name"`
	Password     string `gorm:"password"`
	Url          string `gorm:"url"`
	Description  string `gorm:"description"`
}

func (DbCommonDatabase) TableName() string {
	return "common_database"
}
