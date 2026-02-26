package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/iam/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/iam/domain/iam"
	"github.com/cooperlutz/go-full/pkg/securitee"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type ErrUserNotFound struct{}

func (e ErrUserNotFound) Error() string {
	return "user not found"
}

type ErrInvalidCredentials struct{}

func (e ErrInvalidCredentials) Error() string {
	return "invalid credentials"
}

type ErrInvalidToken struct{}

func (e ErrInvalidToken) Error() string {
	return "invalid token"
}

type ErrExpiredToken struct{}

func (e ErrExpiredToken) Error() string {
	return "token has expired"
}

type ErrEmailInUse struct{}

func (e ErrEmailInUse) Error() string {
	return "email already in use"
}

type IIamQueries interface {
	FindUserByEmail(ctx context.Context, email string) (outbound.IamUser, error)
	FindUserByID(ctx context.Context, id string) (outbound.IamUser, error)
}

// IamService provides authentication functionality.
type IamService struct {
	jwtSecret       []byte
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration

	iamRepository          iam.UserRepository
	refreshTokenRepository iam.RefreshTokenRepository
	Queries                IIamQueries
}

// NewIamService creates a new authentication service.
func NewIamService(
	iamRepository iam.UserRepository,
	refreshTokenRepository iam.RefreshTokenRepository,
	iamQueryInterface IIamQueries,
	jwtSecret string,
	accessTokenTTL, refreshTokenTTL time.Duration,
) *IamService {
	return &IamService{
		iamRepository:          iamRepository,
		refreshTokenRepository: refreshTokenRepository,
		Queries:                iamQueryInterface,
		jwtSecret:              []byte(jwtSecret),
		accessTokenTTL:         accessTokenTTL,
		refreshTokenTTL:        refreshTokenTTL,
	}
}

// Register creates a new user with the provided credentials.
func (s *IamService) Register(ctx context.Context, email, password string) (outbound.IamUser, error) {
	ctx, span := telemetree.AddSpan(ctx, "iam.service.register")
	defer span.End()

	_, err := s.Queries.FindUserByEmail(ctx, email)
	if err == nil {
		telemetree.RecordError(ctx, ErrEmailInUse{}, "email already in use")

		return outbound.IamUser{}, ErrEmailInUse{}
	}

	// Only proceed if the error was "user not found"
	if !errors.Is(err, sql.ErrNoRows) {
		telemetree.RecordError(ctx, err, "failed to check if email is already in use")

		return outbound.IamUser{}, err
	}

	// Hash the password
	hashedPassword, err := securitee.HashPassword(password)
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to hash password")

		return outbound.IamUser{}, err
	}

	user := iam.NewUser(email, hashedPassword)

	// Create the user
	err = s.iamRepository.CreateUser(ctx, user)
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to create user")

		return outbound.IamUser{}, err
	}

	createdUser, err := s.Queries.FindUserByEmail(ctx, email)
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to retrieve created user")

		return outbound.IamUser{}, ErrUserNotFound{}
	}

	return createdUser, nil
}

// ValidateToken verifies a JWT token and returns the claims.
func (s *IamService) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken{}
		}

		return s.jwtSecret, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken{}
		}

		return nil, ErrInvalidToken{}
	}

	// Extract and validate claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken{}
}

// Login authenticates a user and returns both access and refresh tokens.
func (s *IamService) Login(ctx context.Context, email, password string) (accessToken, refreshTokenToken string, err error) {
	ctx, span := telemetree.AddSpan(ctx, "iam.service.login")
	defer span.End()

	existingUser, err := s.Queries.FindUserByEmail(ctx, email)
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to find user by email during login")

		return "", "", ErrUserNotFound{}
	}

	// Verify the password
	if err := securitee.VerifyPassword(existingUser.PasswordHash, password); err != nil {
		telemetree.RecordError(ctx, err, "invalid password during login")

		return "", "", ErrInvalidCredentials{}
	}

	user, err := s.iamRepository.GetUser(ctx, existingUser.ID.String())
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to get user during login")

		return "", "", err
	}

	// Generate an access token
	accessToken, err = user.NewAccessToken(s.jwtSecret, s.accessTokenTTL)
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to generate access token during login")

		return "", "", err
	}

	refreshToken, err := s.createRefreshToken(ctx, user.ID)
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to create refresh token during login")

		return "", "", err
	}

	return accessToken, refreshToken.Token, nil
}

// RefreshAccessToken creates a new access token using a refresh token.
func (s *IamService) RefreshAccessToken(ctx context.Context, refreshTokenString string) (string, error) {
	ctx, span := telemetree.AddSpan(ctx, "iam.service.refresh_access_token")
	defer span.End()

	// Retrieve the refresh token
	token, err := s.refreshTokenRepository.GetRefreshToken(ctx, refreshTokenString)
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to get refresh token")

		return "", ErrInvalidToken{}
	}

	// Check if the token is valid
	if token.Revoked {
		telemetree.RecordError(ctx, ErrInvalidToken{}, "refresh token revoked")

		return "", ErrInvalidToken{}
	}

	// Check if the token has expired
	if time.Now().After(token.ExpiresAt) {
		telemetree.RecordError(ctx, ErrExpiredToken{}, "refresh token expired")

		return "", ErrExpiredToken{}
	}

	// Get the user
	user, err := s.iamRepository.GetUser(ctx, token.UserID.String())
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to get user for refresh token")

		return "", err
	}

	// Generate a new access token
	accessToken, err := user.NewAccessToken(s.jwtSecret, s.accessTokenTTL)
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to generate access token")

		return "", err
	}

	return accessToken, nil
}

func (s *IamService) createRefreshToken(ctx context.Context, userId uuid.UUID) (*iam.RefreshToken, error) {
	ctx, span := telemetree.AddSpan(ctx, "iam.service.create_refresh_token")
	defer span.End()

	refreshToken := iam.NewRefreshToken(userId, s.refreshTokenTTL)

	// Create a refresh token
	err := s.refreshTokenRepository.CreateRefreshToken(ctx, refreshToken)
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to create refresh token")

		return nil, err
	}

	return refreshToken, nil
}
