package iam

import (
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/utilitee"
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
	expiresAt := utilitee.RightNow().Add(refreshTokenTTL)

	token := &RefreshToken{
		ID:        tokenID,
		UserID:    userID,
		Token:     tokenID.String(), // Use the UUID as the token
		ExpiresAt: expiresAt,
		CreatedAt: utilitee.RightNow(),
		Revoked:   false,
	}

	return token
}
