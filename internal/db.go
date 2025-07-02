// Package internal provide the db conn
// use driver mysql
// use gorm
// Copyright (c) 2025 TF Author. All Rights Reserved.
package internal

import (
	"fmt"
	"log/slog"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ahhxfeng/Amp/configs"
)

func InitMysql() error {
	cfg := configs.Conf.Database

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=Ture&loc=Local",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
	// get the database conn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("connect to mysql failed:%v", err)
	}

	// 连接数据库实例
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get the object instance: %v", err)
	}

	// set the conn config
	sqlDB.SetMaxIdleConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxConnLifeTime) * time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Duration(cfg.MaxConnIdleTime) * time.Minute)

	// test conn
	err = sqlDB.Ping()
	if err != nil {
		return fmt.Errorf("ping failed:%v", err)
	}

	slog.Info("Mysql init success ")

	return nil
}
