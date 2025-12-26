package core

import (
	"database/sql"
	"fmt"
	"modulename/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {

	if global.Config.Mysql.Host == "" {
		global.Log.Warnf("未配置mysql，取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var myssqlLogger logger.Interface
	if global.Config.System.ENV == "debug" {
		//开发环境 显示所有的sql
		myssqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		myssqlLogger = logger.Default.LogMode(logger.Error)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: myssqlLogger,
	})
	if err != nil {
		global.Log.Error(fmt.Println("[%s] mysql连接失败", dsn))
		panic(err)
	}

	sqlDB := sql.DB{}
	sqlDB.SetMaxIdleConns(10)  //最大空闲连接数
	sqlDB.SetMaxOpenConns(100) //最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	return db
}
