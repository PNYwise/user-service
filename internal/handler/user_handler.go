package handler

import (
	"context"

	"github.com/PNYwise/user-service/internal/domain"
	user_service "github.com/PNYwise/user-service/proto"
)

type userHandler struct {
	userService domain.IUserService
	user_service.UnimplementedUserServer
}

func NewUserHandler(userService domain.IUserService) *userHandler {
	return &userHandler{userService: userService}
}

func (u *userHandler) CreateUser(ctx context.Context, request *user_service.UserDetail) (*user_service.UserDetail, error) {
	userRequest := &domain.CreateUserRequest{
		Email:    request.GetEmail(),
		Username: request.GetUsername(),
		Timezone: request.GetTimezone(),
	}

	user, err := u.userService.Create(ctx, userRequest)
	if err != nil {
		return nil, err
	}
	return &user_service.UserDetail{
		Uuid:      user.Uuid,
		Email:     user.Email,
		Username:  user.Username,
		Timezone:  user.Timezone,
		Followers: user.Followers,
		Following: user.Following,
	}, nil
}

func (u *userHandler) GetUserByUuid(ctx context.Context, request *user_service.Uuid) (*user_service.UserDetail, error) {
	user, err := u.userService.GetByUuid(ctx, request.GetUuid())
	if err != nil {
		return nil, err
	}
	return &user_service.UserDetail{
		Uuid:      user.Uuid,
		Email:     user.Email,
		Username:  user.Username,
		Timezone:  user.Timezone,
		Followers: user.Followers,
		Following: user.Following,
	}, nil
}
