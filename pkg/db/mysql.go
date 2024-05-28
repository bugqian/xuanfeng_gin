package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"
	"time"
	"xuanfeng_gin/pkg/log"
)

type MysqlConf struct {
	Connection      string
	MaxIdleConns    int
	ConnMaxIdleTime int
	ConnMaxLifeTime int
	MaxOpenConns    int
}

func NewMysql(conf *MysqlConf) (*gorm.DB, func(), error) {
	gormZapLog := zapgorm2.New(log.L)
	gormZapLog.SetAsDefault()
	gormZapLog.LogLevel = logger.Info
	db, err := gorm.Open(mysql.Open(conf.Connection), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // 单数表名

	}, Logger: gormZapLog})

	if err != nil {
		return nil, nil, err
	}
	dbInstance, _ := db.DB()
	dbInstance.SetMaxIdleConns(conf.MaxIdleConns)
	dbInstance.SetConnMaxIdleTime(time.Second * time.Duration(conf.ConnMaxIdleTime))
	dbInstance.SetConnMaxLifetime(time.Second * time.Duration(conf.ConnMaxLifeTime))
	if conf.MaxOpenConns > 0 {
		dbInstance.SetMaxOpenConns(conf.MaxOpenConns)
	}
	deferFunc := func() {
		dbInstance.Close()
	}
	return db, deferFunc, nil
}
