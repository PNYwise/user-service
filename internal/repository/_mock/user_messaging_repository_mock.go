package _mock

import (
	"github.com/PNYwise/user-service/internal/domain"
	"github.com/stretchr/testify/mock"
)

type UserMessagingRepositoryMock struct {
	mock.Mock
}

func (u *UserMessagingRepositoryMock) PublishMessage(user *domain.User) error {
	args := u.Called(user)
	return args.Error(0)
}
