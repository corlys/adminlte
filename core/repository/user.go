package repository

import (
	"errors"
	"fmt"

	"github.com/corlys/adminlte/core/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	GetUserByEmail(email string) (entity.User, error)
	CreateNewUser(user entity.User) (entity.User, error)
	UpsertTotpSecret(user entity.User, secret string) error
	GetTotpSecret(email string) (string, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.db.Debug().Where("email = $1", email).Take(&user).Error
	if err != nil && (!errors.Is(err, gorm.ErrRecordNotFound)) {
		return entity.User{}, err
	}
	return user, nil
}
func (r *userRepository) CreateNewUser(user entity.User) (entity.User, error) {
	if err := r.db.Debug().Create(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}
func (r *userRepository) UpsertTotpSecret(user entity.User, secret string) error {
	err := r.db.Debug().Model(&entity.User{}).Where("email = ?", user.Email).Update("totp_secret", secret).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (r *userRepository) GetTotpSecret(email string) (string, error) {
	var user entity.User
	err := r.db.Debug().Where("email = $1", email).Take(&user).Error
	if err != nil && (!errors.Is(err, gorm.ErrRecordNotFound)) {
		return "", err
	}
	return *user.TotpSecret, nil
}
