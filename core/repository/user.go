package repository

import (
	"errors"

	"github.com/corlys/adminlte/core/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	GetUserByEmail(email string) (entity.User, error)
	CreateNewUser(user entity.User) (entity.User, error)
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