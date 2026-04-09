package iam

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/utilitee"
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

// NewAccessToken creates a new JWT access token signed with an RSA private key (RS256).
func (u *User) NewAccessToken(privateKey *rsa.PrivateKey, accessTokenTTL time.Duration) (string, error) {
	expirationTime := utilitee.RightNow().Add(accessTokenTTL)

	claims := jwt.MapClaims{
		"sub":   u.ID.String(),              // subject (user ID)
		"email": u.Email,                    // custom claim
		"exp":   expirationTime.Unix(),      // expiration time
		"iat":   utilitee.RightNow().Unix(), // issued at time
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
