package config

import (
	"log"
	"time"

	"exchangeapp/EXCHANGEAPP_BACKEND/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := Appconfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to initialize database, got error: %v", err)
	}

	sqlDB, err := db.DB()

	//设置连接池中空闲最大连接数量
	sqlDB.SetMaxIdleConns(Appconfig.Database.MaxIdleConns)
	//设置打开数据库时的最大连接数量
	sqlDB.SetMaxOpenConns(Appconfig.Database.MaxOpenConns)
	//每个连接使用一小时之后会关闭
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatalf("Failed to configure database, got an error: %v", err)
	}

	//将数据库实例赋值给全局变量
	global.Db = db

}
