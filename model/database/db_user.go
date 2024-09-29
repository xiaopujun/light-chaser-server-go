package database

import "github.com/light-chaser/server/model/database/base"

type DbUser struct {
	base.ORMBaseModel
	Name     string `gorm:"name"`
	UserName string `gorm:"user_name"`
	Email    string
	Phone    string
	Avatar   string
	Roles    string
}

func (DbUser) TableName() string {
	return "user"
}
