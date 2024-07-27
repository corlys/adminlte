package seeder

import (
	"errors"

	"github.com/corlys/adminlte/core/entity"
	"gorm.io/gorm"
)

func Userseeder(db *gorm.DB) error {
	var dummyUsers = []entity.User{
		{
			Email:    "opunk55@gmail.com",
			FullName: "Mehmud",
			Password: "123456",
		},
	}

	hasTable := db.Migrator().HasTable(&entity.User{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.User{}); err != nil {
			return err
		}
	}

	for _, data := range dummyUsers {
		var user entity.User
		err := db.Where(&entity.User{Email: data.Email}).First(&user).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&user, "email = ?", data.Email).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
