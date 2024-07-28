package service

import (
	"fmt"

	"github.com/corlys/adminlte/core/helper/dto"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type sessionService struct{}

type SessionService interface {
	GetUserSession(ctx *gin.Context) dto.UserResponse
	SetUserSession(ctx *gin.Context, userDto dto.UserResponse)
	DeleteUserSession(ctx *gin.Context)
	SetTotpKeySession(ctx *gin.Context, key string)
	GetTotpKeySession(ctx *gin.Context) string
	DeleteTotpKeySession(ctx *gin.Context)
}

func NewSessionService() SessionService {
	return &sessionService{}
}

func (ss sessionService) GetUserSession(ctx *gin.Context) dto.UserResponse {
	session := sessions.Default(ctx)
	userInterface := session.Get("user")
	fmt.Println("userInterface ", userInterface)
	if userInterface == nil {
		return dto.UserResponse{}
	}
	user, ok := userInterface.(dto.UserResponse)
	if !ok {
		return dto.UserResponse{}
	} else {
		return user
	}
}
func (ss sessionService) SetUserSession(ctx *gin.Context, userDto dto.UserResponse) {
	session := sessions.Default(ctx)
	session.Set("user", userDto)
	err := session.Save()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Session is saved ", userDto)
}
func (ss sessionService) DeleteUserSession(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete("user")
	session.Save()
}
func (ss sessionService) SetTotpKeySession(ctx *gin.Context, key string) {
	session := sessions.Default(ctx)
	session.Set("totp-key", key)
	session.Save()
}
func (ss sessionService) GetTotpKeySession(ctx *gin.Context) string {
	session := sessions.Default(ctx)
	key := session.Get("totp-key")
	if key == nil {
		return ""
	}
	secret, ok := key.(string)
	if !ok {
		return ""
	} else {
		return secret
	}
}
func (ss sessionService) DeleteTotpKeySession(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete("totp-key")
	session.Save()
}
