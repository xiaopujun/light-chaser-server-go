package database

import "github.com/light-chaser/server/model/database/base"

type DbLocalAuth struct {
	base.ORMBase
	Password string `gorm:"password"`
	UserId   int64  `gorm:"user_id"`
}

func (DbLocalAuth) TableName() string {
	return "local_auth"
}
