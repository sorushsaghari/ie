package user

import (
	"crypto/rand"
	"fmt"
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

type Auth struct {
	gorm.Model
	UserID uint
	Token string
}
func Store(auth *Auth) error {
	db := database.DB().Create(&auth).GetErrors()
	if len(db) != 0 {
		log.Println(db)
		return db[0]
	}
	return nil
}
func Create(user *User)	error {
	db := database.DB().Create(&user).GetErrors()
	if len(db) != 0 {
		log.Println(db)
		return db[0]
	}
	return nil
}

func One(user *User) (*User, error) {
	var result User
	err := database.DB().Where(&user).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func TokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}