package iam

import (
	"time"

	"github.com/cooperlutz/go-full/pkg/utilitee"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	LastLogin    *time.Time
}

func NewUser(email, passwordHash string) *User {
	return &User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    utilitee.RightNow(),
		LastLogin:    nil,
	}
}

// NewAccessToken creates a new JWT access token.
func (u *User) NewAccessToken(jwtSecret []byte, accessTokenTTL time.Duration) (string, error) {
	expirationTime := utilitee.RightNow().Add(accessTokenTTL)

	// Create the JWT claims
	claims := jwt.MapClaims{
		"sub":   u.ID.String(),              // subject (user ID)
		"email": u.Email,                    // custom claim
		"exp":   expirationTime.Unix(),      // expiration time
		"iat":   utilitee.RightNow().Unix(), // issued at time
	}

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
