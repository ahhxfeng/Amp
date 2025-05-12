/*
* db conn about this project
 */

package internal

import (
	"fmt"
	"log/slog"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ahhxfeng/Amp/configs"
)

func Conn() {
	// dsn := "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"

}

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
		return fmt.Errorf("Connect to mysql failed:%v", err)
	}

	// 连接数据库实例
	sqlDB, err := db.DB()

	if err != nil {
		return fmt.Errorf("获取实例化对象失败：%v", err)
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
