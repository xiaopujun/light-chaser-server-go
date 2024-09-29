package database

import "github.com/light-chaser/server/model/database/base"

// DbProject 数据库project 表映射结构体
type DbProject struct {
	base.ORMBaseModel
	Name     string `gorm:"name" json:"name,omitempty"`
	Des      string `gorm:"des" json:"des,omitempty"`
	DataJson string `gorm:"data_json" json:"dataJson,omitempty"`
	Cover    string `gorm:"cover" json:"cover,omitempty"`
	UserId   int64  `gorm:"user_id" json:"userId,omitempty"`
}

// TableName 自定义映射数据库表名
func (DbProject) TableName() string {
	return "project"
}
