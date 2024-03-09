package repository

import (
	"context"
	"fmt"
	"log"
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
		RETURNING id, uuid, followers, following, created_at`
	err := u.db.QueryRow(
		ctx,
		query,
		user.Email, user.Username, user.Timezone, time.Now(),
	).Scan(&user.Id, &user.Uuid, &user.Followers, &user.Following, &user.CreatedAt)
	if err != nil {
		fmt.Printf("Error executing query: %v", err)
		return err
	}
	return nil
}

// GetByUuid implements domain.IUserRepository.
func (u *userRepository) GetByUuid(ctx context.Context, Uuid string) (*domain.User, error) {
	query := `
		SELECT u.id, u.uuid, u.email, u.username, u.timezone, u.created_at 
		FROM users u 
		WHERE u.uuid = $1 
		LIMIT 1`
	var user domain.User
	err := u.db.QueryRow(ctx, query, Uuid).Scan(&user.Id, &user.Uuid, &user.Email, &user.Username, &user.Timezone, &user.CreatedAt)
	if err != nil {
		fmt.Printf("Error executing query: %v", err)
		return nil, err
	}
	return &user, nil
}

// ExistByUsername implements domain.IUserRepository.
func (u *userRepository) ExistByUsername(ctx context.Context, username string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)"
	var exist bool
	row, err := u.db.Query(context.Background(), query, username)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
		return false, err
	}
	for row.Next() {
		if err := row.Scan(&exist); err != nil {
			log.Fatalf("Error Scaning query: %v", err)
			return false, err
		}
	}
	return exist, nil
}

// TODO: exist by email, username
