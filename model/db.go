package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"time"
)

type Config struct {
	UserName string
	Password string
	Host     string
	DbName   string
}

var DB *gorm.DB

func ConnDB(c Config) (err error) {

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", c.UserName, c.Password, c.Host, c.DbName)
	log.Debugf("url:%s", url)
	DB, err = gorm.Open("mysql", url)

	// 启用Logger，显示详细日志
	DB.LogMode(true)
	DB.DB().SetConnMaxLifetime(time.Second * 20)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	return err
}

func CloseDB() error {
	return DB.Close()
}
