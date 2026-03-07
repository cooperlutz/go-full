package iam

import "context"

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, userId string) (*User, error)
	UpdateUser(
		ctx context.Context,
		userId string,
		updateFn func(*User) (*User, error),
	) error
}

type RefreshTokenRepository interface {
	CreateRefreshToken(ctx context.Context, token *RefreshToken) error
	GetRefreshToken(ctx context.Context, id string) (*RefreshToken, error)
	UpdateRefreshToken(
		ctx context.Context,
		tokenId string,
		updateFn func(*RefreshToken) (*RefreshToken, error),
	) error
}
