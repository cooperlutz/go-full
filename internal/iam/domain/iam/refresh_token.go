package iam

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
	Revoked   bool
}

func NewRefreshToken(userID uuid.UUID, refreshTokenTTL time.Duration) *RefreshToken {
	tokenID := uuid.New()
	expiresAt := time.Now().Add(refreshTokenTTL)

	token := &RefreshToken{
		ID:        tokenID,
		UserID:    userID,
		Token:     tokenID.String(), // Use the UUID as the token
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
		Revoked:   false,
	}

	return token
}
