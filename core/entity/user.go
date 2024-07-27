package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email       string  `json:"email" gorm:"unique;not null"`
	FullName    string  `json:"fullname" gorm:"not null"`
	Password    string  `json:"password" gorm:"not null"`
	GravatarUrl *string `json:"gravatarUrl"`
}
