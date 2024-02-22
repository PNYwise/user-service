package domain

import (
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

type IUserService interface {
	Create(user *User) (*User, error)
	GetByUuid(uuid string) (*User, error)
}

type IUserRepository interface {
	Create(user *User) error
	GetByUuid(Uuid string) (*User, error)
}
