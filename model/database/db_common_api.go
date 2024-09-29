package database

import (
	"github.com/light-chaser/server/model/database/base"
)

type DbCommonApi struct {
	base.ORMBaseModel
	Name    string
	Url     string
	Method  string
	Params  string
	Headers string
}

func (DbCommonApi) TableName() string {
	return "common_api"
}
