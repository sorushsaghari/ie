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

func Find(userId uint) ([] *Note, error){
	notes := make([]*Note,1)
	errs := database.DB().Where("user_id=?", userId).Find(&notes).Error
	if errs != nil {
		return nil, errs
	}
	return notes, nil
}
func One(id uint) (*Note, error) {
	result := Note{}
	err := database.DB().Where("id=?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}