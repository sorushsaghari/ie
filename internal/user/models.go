package user

import (
	"crypto/rand"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sorushsaghari/ie/internal/platforms/database"
	"log"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Family   string
	Username string
	Email    string `gorm:"unique;not null"`
	Password string
	Avatar   string
}

type Auth struct {
	gorm.Model
	UserID  uint
	Token   string
	TimeOut *time.Time
}

func Store(auth *Auth) error {
	if err := database.DB().Where("user_id=?", auth.UserID).First(&Auth{}).RecordNotFound(); err == false {
		fmt.Println(auth)
		db := database.DB().Debug().Model(&Auth{}).Where("user_id=?", auth.UserID).Update(&auth).GetErrors()
		if len(db) != 0 {
			log.Println(db)
			return db[0]
		}
		return nil
	}

	db := database.DB().Create(&auth).GetErrors()
	if len(db) != 0 {
		log.Println(db)
		return db[0]
	}
	return nil
}

func UserByToken(token string) (*User, *time.Time, error) {
	user := User{}
	auth := Auth{Token: token}
	err := database.DB().Debug().First(&auth).Error
	if err != nil {
		return nil, nil, err
	}
	err = database.DB().Debug().Where("id=?", auth.UserID).First(&user).Error
	if err != nil {
		return nil, nil, err
	}
	return &user, auth.TimeOut, nil
}
func Create(user *User) error {
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
