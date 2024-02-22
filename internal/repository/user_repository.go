package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/PNYwise/user-service/internal/domain"
	"github.com/jackc/pgx/v5"
)

type userRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) domain.IUserRepository {
	return &userRepository{db}
}

// Create implements domain.IUserRepository.
func (u *userRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users(email,username,timezone,created_at)
		VALUES($1,$2,$3,$4)
		RETURNING (id,uuid)`
	err := u.db.QueryRow(
		ctx,
		query,
		user.Email, user.Username, user.Timezone, time.Now(),
	).Scan(&user.Id, &user.Uuid, &user.Username, &user.Timezone, &user.CreatedAt)
	if err != nil {
		fmt.Printf("Error executing query: %v", err)
		return err
	}
	return nil
}

// GetByUuid implements domain.IUserRepository.
func (p *userRepository) GetByUuid(ctx context.Context, Uuid string) (*domain.User, error) {
	query := `SELECT u.id, u.uuid, u.username, u.timezone, u.created_at FROM users u WHERE u.uuid = $1 LIMIT 1`
	var user domain.User
	err := p.db.QueryRow(ctx, query, Uuid).Scan(&user.Id, &user.Uuid, &user.Username, &user.Timezone, &user.CreatedAt)
	if err != nil {
		fmt.Printf("Error executing query: %v", err)
		return nil, err
	}
	return &user, nil
}
