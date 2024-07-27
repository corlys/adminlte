package service

import (
	"fmt"

	"github.com/corlys/adminlte/core/helper/dto"
	"github.com/corlys/adminlte/core/repository"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	VerifyLogin(email string, password string) bool
	GetUserByEmail(email string) (dto.UserResponse, error)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (s *userService) VerifyLogin(email string, password string) bool {
	return true
}
func (s *userService) GetUserByEmail(email string) (dto.UserResponse, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return dto.UserResponse{}, err
	}
	userId := fmt.Sprint(user.ID)
	userDto := dto.UserResponse{
		ID:    userId,
		Email: user.Email,
		Name:  user.FullName,
	}
	if user.GravatarUrl != nil {
		userDto.Picture = *user.GravatarUrl
	}
	return userDto, nil
}
