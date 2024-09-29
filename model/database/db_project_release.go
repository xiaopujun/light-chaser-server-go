package database

import "github.com/light-chaser/server/model/database/base"

type DbProjectRelease struct {
	base.ORMBaseModel
	Name           string `gorm:"name"`
	State          int    `gorm:"state"`
	Encrypt        bool   `gorm:"encrypt"`
	Password       string `gorm:"password"`
	DataJson       string `gorm:"data_json"`
	ProjectId      int64  `gorm:"project_id"`
	ExpirationDate int    `gorm:"expiration_date"`
}

func (DbProjectRelease) TableName() string {
	return "project_release"
}
