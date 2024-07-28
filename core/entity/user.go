package entity

import (
	"github.com/corlys/adminlte/common/util"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string  `json:"email" gorm:"unique;not null"`
	FullName    string  `json:"fullname" gorm:"not null"`
	Password    string  `json:"password" gorm:"not null"`
	TotpSecret  *string `json:"totpSecret"`
	GravatarUrl *string `json:"gravatarUrl"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	var err error
	u.Password, err = util.PasswordHash(u.Password)
	if err != nil {
		return err
	}
	gravatarUrl := util.GetGravatarURL(u.Email)
	u.GravatarUrl = &gravatarUrl
	return nil
}

func (u *User) BeforeUpdate(db *gorm.DB) error {
	if u.Password != "" {
		var err error
		u.Password, err = util.PasswordHash(u.Password)
		if err != nil {
			return err
		}
	}
	gravatarUrl := util.GetGravatarURL(u.Email)
	u.GravatarUrl = &gravatarUrl
	return nil
}

