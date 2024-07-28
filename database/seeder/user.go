package seeder

import (
	"errors"

	"github.com/corlys/adminlte/core/entity"

	"github.com/pquerna/otp/totp"
	"gorm.io/gorm"
)

func Userseeder(db *gorm.DB) error {

	var dummyUsers = []entity.User{
		{
			Email:    "opunk55@gmail.com",
			FullName: "Mehmud",
			Password: "password",
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

		key, errGenOtp := totp.Generate(totp.GenerateOpts{
			Issuer:      "app",
			AccountName: data.Email,
		})
		if errGenOtp != nil {
			return errGenOtp
		}

		totpSecret := key.Secret()
		data.TotpSecret = &totpSecret

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
