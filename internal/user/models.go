package user

import (
	"github.com/jinzhu/gorm"
	"github.com/sorushsaghari/ie/internal/platforms/database"
	"log"
)

type User struct {
	gorm.Model
	Name string
	Family string
	Username string
	Email string `gorm:"unique;not null"`
	Password string
	Avatar string
}

func Create(user *User)	error {
	db := database.DB().Create(&user).GetErrors()
	if len(db) != 0 {
		log.Println(db)
		return db[0]
	}
	return nil
}