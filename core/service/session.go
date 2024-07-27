package service

import (
	"github.com/corlys/adminlte/core/helper/dto"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type sessionService struct{}

type SessionService interface {
	GetUserSession(ctx *gin.Context, userDto dto.UserResponse) dto.UserResponse
	SetUserSession(ctx *gin.Context, userDto dto.UserResponse)
}

func NewSessionService() SessionService {
	return &sessionService{}
}

func (ss sessionService) GetUserSession(ctx *gin.Context, userDto dto.UserResponse) dto.UserResponse {
	session := sessions.Default(ctx)
	userInterface := session.Get("user:" + userDto.Email)
	if userInterface != nil {
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
	session.Set("user:"+userDto.Email, userDto)
	session.Save()
}
