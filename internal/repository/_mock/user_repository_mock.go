package _mock

import (
	"context"

	"github.com/PNYwise/user-service/internal/domain"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (u *UserRepositoryMock) Create(ctx context.Context, user *domain.User) error {
	args := u.Called(ctx, user)
	return args.Error(0)
}

func (u *UserRepositoryMock) GetByUuid(ctx context.Context, Uuid string) (*domain.User, error) {
	args := u.Called(ctx, Uuid)
	return args.Get(0).(*domain.User), args.Error(1)
}
