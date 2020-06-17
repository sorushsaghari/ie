package user

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string
	Family string
	Username string
	Email string
	Password string
	Avatar string
}

func One()  {

}