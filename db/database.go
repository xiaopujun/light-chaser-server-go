package db

import (
	"fmt"
	"github.com/light-chaser/server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Product struct {
	gorm.Model
	Name  string
	Price float64
}

func InitDatabase() *gorm.DB {
	dbConfig := global.LC_ENV.Database
	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" + dbConfig.Schema + "?charset=utf8mb4&parseTime=True&loc=Local"
	//连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to the database")
	}
	//创建数据库连接池
	dbPool, err := db.DB()
	if err != nil {
		fmt.Println(err)
		panic("Failed to get database connection pool")
	}
	dbPool.SetMaxIdleConns(10)           // 设置空闲连接池中连接的最大数量
	dbPool.SetMaxOpenConns(100)          // 设置打开数据库连接的最大数量
	dbPool.SetConnMaxLifetime(time.Hour) // 设置连接可复用的最大时间
	return db
}
