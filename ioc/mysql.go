package ioc

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zhuguangfeng123/go-chat/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"log"
)

// InitMysql 初始话mysql连接
func InitMysql() *gorm.DB {
	type Config struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Database string `yaml:"database"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}
	var cfg Config
	err := viper.UnmarshalKey("mysql", &cfg)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.Open(
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)), &gorm.Config{
		Logger: glogger.New(log.New(log.Writer(), "\r\n", log.LstdFlags), glogger.Config{
			// 慢查询
			SlowThreshold: 0,
			LogLevel:      glogger.Info,
		}),
	})

	if err != nil {
		panic(err)
	}

	if err := model.InitTables(db); err != nil {
		panic(err)
	}

	return db
}
