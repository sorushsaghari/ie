package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Init(password string, user string, port int, dbname string, host string) error {
	DB, err := gorm.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?connect_timeout=10&sslmode=disable", user, password, host, port, dbname))
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