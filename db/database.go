package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Name  string
	Price float64
}

func InitDataBase() *gorm.DB {
	//连接数据库
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/test_go"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("连接数据库失败")
	}
	//创建数据库连接池
	dbPool, err := db.DB()
	if err != nil {
		fmt.Println(err)
		panic("获取数据库连接池失败")
	}
	dbPool.SetMaxIdleConns(10)           // 设置空闲连接池中连接的最大数量
	dbPool.SetMaxOpenConns(100)          // 设置打开数据库连接的最大数量
	dbPool.SetConnMaxLifetime(time.Hour) // 设置连接可复用的最大时间
	return db
}
