package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Init() error {
	DB, err := gorm.Open("mysql", "user:password@localhost/dbname:port?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}

	db = DB
	return nil
}

func DB() *gorm.DB { return db }

func Close() {
	db.Close()
}