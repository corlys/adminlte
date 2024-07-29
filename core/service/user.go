package service

import (
	"fmt"
	"reflect"

	"github.com/corlys/adminlte/common/util"
	"github.com/corlys/adminlte/core/entity"
	"github.com/corlys/adminlte/core/helper/dto"
	errs "github.com/corlys/adminlte/core/helper/errors"
	"github.com/corlys/adminlte/core/repository"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	VerifyLogin(email string, password string) bool
	GetUserByEmail(email string) (dto.UserResponse, error)
	RegisterUser(userRequest dto.UserRegisterRequest) (dto.UserResponse, error)
	GenerateTotp(email string) (*otp.Key, error)
	ValidateTotp(email string, code string) bool
	GetUserTotp(email string) string
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (s *userService) VerifyLogin(email string, password string) bool {
	userCheck, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		fmt.Println("wrong user")
		return false
	}
	passwordCheck, err := util.PasswordCompare(userCheck.Password, []byte(password))
	if err != nil {
		fmt.Println("wrong password")
		return false
	}

	if userCheck.Email == email && passwordCheck {
		return true
	}
	return false
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
func (s *userService) RegisterUser(userRequest dto.UserRegisterRequest) (dto.UserResponse, error) {
	userCheck, err := s.userRepository.GetUserByEmail(userRequest.Email)
	if err != nil {
		return dto.UserResponse{}, err
	}
	if !(reflect.DeepEqual(userCheck, entity.User{})) {
		return dto.UserResponse{}, errs.ErrEmailAlreadyExists
	}
	user := entity.User{
		FullName: userRequest.FullName,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
	fmt.Println("data entity", user)
	res, err := s.userRepository.CreateNewUser(user)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return dto.UserResponse{
		ID:      fmt.Sprint(res.ID),
		Name:    res.FullName,
		Email:   res.Email,
		Picture: *res.GravatarUrl,
	}, nil
}
func (s *userService) GenerateTotp(email string) (*otp.Key, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "adminlte-go",
		AccountName: email,
	})
	if err != nil {
		return nil, err
	}
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	fmt.Println(key.Secret())
	fmt.Println(key.URL())
	err = s.userRepository.UpsertTotpSecret(user, key.URL())
	if err != nil {
		return nil, err
	}
	return key, nil
}
func (s *userService) GetUserTotp(email string) string {
	url, err := s.userRepository.GetTotpSecret(email)
	if err != nil || url == "" {
		fmt.Println(err, url)
		return ""
	}
	return url
}
func (s *userService) ValidateTotp(email string, code string) bool {
	secret, err := s.userRepository.GetTotpSecret(email)
	if err != nil || secret == "" {
		fmt.Println(err, secret)
		return false
	}
	key, err := otp.NewKeyFromURL(secret)
	if err != nil {
		fmt.Println(err, secret, key)
		return false
	}
	return totp.Validate(code, key.Secret())
}
