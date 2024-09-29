package database

import "github.com/light-chaser/server/model/database/base"

type DbFile struct {
	base.ORMBaseModel
	Name      string `gorm:"name"`
	Url       string `gorm:"url"`
	Type      int    `gorm:"type"`
	ProjectId int64  `gorm:"project_id"`
	Hash      string `gorm:"hash"`
}

func (DbFile) TableName() string {
	return "file"
}
