package service

import (
	"context"
	"errors"

	"github.com/PNYwise/user-service/internal/domain"
)

type userService struct {
	userRepo domain.IUserRepository
}

func NewUserService(userRepo domain.IUserRepository) domain.IUserService {
	return &userService{userRepo}
}

// Create implements domain.IUserService.
func (u *userService) Create(ctx context.Context, request *domain.CreateUserRequest) (*domain.User, error) {
	user := &domain.User{
		Email:    request.Email,
		Username: request.Username,
		Timezone: request.Timezone,
	}
	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, errors.New("internal server error")
	}
	return user, nil
}

// GetByUuid implements domain.IUserService.
func (p *userService) GetByUuid(ctx context.Context, uuid string) (*domain.User, error) {
	post, err := p.userRepo.GetByUuid(ctx, uuid)
	if err != nil {
		return nil, errors.New("internal server error")
	}
	return post, nil
}
