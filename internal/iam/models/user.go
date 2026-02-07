package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// User represents a user in our system.
type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	LastLogin    *time.Time
}

// UserRepository handles database operations for users.
type UserRepository struct {
	db DBTX
}

// NewUserRepository creates a new user repository.
func NewUserRepository(db DBTX) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser adds a new user to the database.
func (r *UserRepository) CreateUser(email, passwordHash string) (*User, error) {
	user := &User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
	}

	query := `
        INSERT INTO iam.users (id, email, password_hash, created_at)
        VALUES ($1, $2, $3, $4)
    `

	_, err := r.db.Exec(context.Background(), query, user.ID, user.Email, user.PasswordHash, user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByEmail retrieves a user by their email address.
func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	query := `SELECT id, email, password_hash, created_at, last_login FROM iam.users WHERE email = $1`

	var (
		user      User
		lastLogin sql.NullTime
	)

	err := r.db.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&lastLogin,
	)
	if err != nil {
		return nil, err
	}

	if lastLogin.Valid {
		user.LastLogin = &lastLogin.Time
	}

	return &user, nil
}

// GetUserByID retrieves a user by their ID.
func (r *UserRepository) GetUserByID(id uuid.UUID) (*User, error) {
	query := `SELECT id, email, password_hash, created_at, last_login FROM iam.users WHERE id = $1`

	var (
		user      User
		lastLogin sql.NullTime
	)

	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&lastLogin,
	)
	if err != nil {
		return nil, err
	}

	if lastLogin.Valid {
		user.LastLogin = &lastLogin.Time
	}

	return &user, nil
}
