package domain

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	Id        uint64
	Uuid      string
	Email     string
	Username  string
	Timezone  string
	Followers uint64
	Following uint64
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

type CreateUserRequest struct {
	Email    string
	Username string
	Timezone string
}

type IUserService interface {
	Create(ctx context.Context, request *CreateUserRequest) (*User, error)
	GetByUuid(ctx context.Context, uuid string) (*User, error)
}

type IUserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByUuid(ctx context.Context, Uuid string) (*User, error)
}

type IUserMessagingRepository interface {
	PublishMessage(user *User) error
}
