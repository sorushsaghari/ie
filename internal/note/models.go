package note

import (
	"github.com/jinzhu/gorm"
	"github.com/sorushsaghari/ie/internal/platforms/database"
	"log"
)

type Note struct {
	gorm.Model
	Text   string
	UserID uint
	Topic  string
}

func Create(note Note) error {
	db := database.DB().Create(&note).GetErrors()
	if len(db) != 0 {
		log.Println(db)
		return db[0]
	}
	return nil

}
