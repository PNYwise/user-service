package service

import (
	"context"
	"testing"

	"github.com/PNYwise/user-service/internal/domain"
	"github.com/PNYwise/user-service/internal/repository/_mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// init mock
var mockUserRepo = new(_mock.UserRepositoryMock)
var mockUserMessagingRepo = new(_mock.UserMessagingRepositoryMock)

// ctx
var ctx = context.Background()

// call user service
var calledUserService = NewUserService(mockUserRepo, mockUserMessagingRepo)

var email = "test@gmail.com"
var username = "test.test"
var timeZone = "Asia/Jakarta"

func TestCreate(t *testing.T) {
	request := &domain.CreateUserRequest{
		Email:    email,
		Username: username,
		Timezone: timeZone,
	}
	user := &domain.User{
		Email:     email,
		Username:  username,
		Timezone:  timeZone,
		Followers: 0,
		Following: 0,
	}

	// Expect the Create method to be called with the correct argument
	mockUserRepo.On("Create", ctx, user).Return(nil)
	mockUserRepo.On("ExistByUsername", ctx, username).Return(false, nil)
	mockUserMessagingRepo.On("PublishMessage", user).Return(nil)
	createUser, err := calledUserService.Create(ctx, request)

	// Assert that the mock repository's Create method was called with the correct argument
	mockUserRepo.AssertExpectations(t)
	mockUserMessagingRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, request.Email, createUser.Email)
	assert.Equal(t, request.Timezone, createUser.Timezone)
	assert.Equal(t, request.Username, createUser.Username)
}

func TestGetByUuid(t *testing.T) {

	fakeUuid := uuid.New().String()

	user := &domain.User{
		Id:        1,
		Uuid:      fakeUuid,
		Email:     email,
		Username:  username,
		Timezone:  timeZone,
		Followers: 0,
		Following: 0,
	}
	mockUserRepo.On("GetByUuid", ctx, fakeUuid).Return(user, nil)
	getUserByUuid, err := calledUserService.GetByUuid(ctx, fakeUuid)

	// Assert that the mock repository's Create method was called with the correct argument
	mockUserRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, getUserByUuid.Email)
	assert.Equal(t, user.Timezone, getUserByUuid.Timezone)
	assert.Equal(t, user.Username, getUserByUuid.Username)
	assert.Equal(t, user.Id, getUserByUuid.Id)
	assert.Equal(t, user.Uuid, getUserByUuid.Uuid)
	assert.Equal(t, user.Uuid, getUserByUuid.Uuid)

}
