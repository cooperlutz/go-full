package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/cooperlutz/go-full/internal/iam/adapters/outbound"
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

// IamService provides authentication functionality.
type IamService struct {
	iamRepository   outbound.Querier
	jwtSecret       []byte
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

// NewIamService creates a new authentication service.
func NewIamService(iamRepository outbound.Querier, jwtSecret string, accessTokenTTL, refreshTokenTTL time.Duration) *IamService {
	return &IamService{
		iamRepository:   iamRepository,
		jwtSecret:       []byte(jwtSecret),
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

func (s *IamService) GetRefreshTokenTTL() time.Duration {
	return s.refreshTokenTTL
}

// Register creates a new user with the provided credentials.
func (s *IamService) Register(ctx context.Context, email, password string) (outbound.IamUser, error) {
	ctx, span := telemetree.AddSpan(ctx, "iam.service.register")
	defer span.End()

	_, err := s.iamRepository.GetUserByEmail(ctx, outbound.GetUserByEmailParams{Email: email})
	if err == nil {
		return outbound.IamUser{}, ErrEmailInUse{}
	}

	// Only proceed if the error was "user not found"
	if !errors.Is(err, sql.ErrNoRows) {
		return outbound.IamUser{}, err
	}

	// Hash the password
	hashedPassword, err := securitee.HashPassword(password)
	if err != nil {
		return outbound.IamUser{}, err
	}

	userId := uuid.New()

	// Create the user
	user, err := s.iamRepository.CreateUser(ctx, outbound.CreateUserParams{
		ID:           pgtype.UUID{Bytes: userId, Valid: true},
		Email:        email,
		PasswordHash: hashedPassword,
		LastLogin:    pgtype.Timestamp{Time: time.Time{}, Valid: false},
	})
	if err != nil {
		return outbound.IamUser{}, err
	}

	return user, nil
}

// generateAccessToken creates a new JWT access token.
func (s *IamService) generateAccessToken(user outbound.IamUser) (string, error) {
	expirationTime := time.Now().Add(s.accessTokenTTL)

	// Create the JWT claims
	claims := jwt.MapClaims{
		"sub":   user.ID.String(),      // subject (user ID)
		"email": user.Email,            // custom claim
		"exp":   expirationTime.Unix(), // expiration time
		"iat":   time.Now().Unix(),     // issued at time
	}

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret key
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
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

// LoginWithRefresh authenticates a user and returns both access and refresh tokens.
func (s *IamService) LoginWithRefresh(ctx context.Context, email, password string, refreshTokenTTL time.Duration) (accessToken, refreshToken string, err error) {
	ctx, span := telemetree.AddSpan(ctx, "iam.service.loginwithrefresh")
	defer span.End()

	user, err := s.iamRepository.GetUserByEmail(ctx, outbound.GetUserByEmailParams{Email: email})
	if err != nil {
		return "", "", ErrInvalidCredentials{}
	}

	// Verify the password
	if err := securitee.VerifyPassword(user.PasswordHash, password); err != nil {
		return "", "", ErrInvalidCredentials{}
	}

	// Generate an access token
	accessToken, err = s.generateAccessToken(user)
	if err != nil {
		return "", "", err
	}

	expiresAt := time.Now().Add(refreshTokenTTL)
	refreshTokenId := uuid.New()

	// Create a refresh token
	token, err := s.iamRepository.CreateRefreshToken(ctx, outbound.CreateRefreshTokenParams{
		ID:        pgtype.UUID{Bytes: refreshTokenId, Valid: true},
		UserID:    user.ID,
		Token:     accessToken,
		ExpiresAt: pgtype.Timestamp{Time: expiresAt, Valid: true},
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		Revoked:   false,
	})
	if err != nil {
		return "", "", err
	}

	return token.Token, refreshTokenId.String(), nil
}

// RefreshAccessToken creates a new access token using a refresh token.
func (s *IamService) RefreshAccessToken(ctx context.Context, refreshTokenString string) (string, error) {
	ctx, span := telemetree.AddSpan(ctx, "iam.service.refreshaccesstoken")
	defer span.End()

	// Retrieve the refresh token
	token, err := s.iamRepository.GetRefreshToken(ctx, outbound.GetRefreshTokenParams{
		ID: pgtype.UUID{Bytes: uuid.MustParse(refreshTokenString), Valid: true},
	})
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
	if time.Now().After(token.ExpiresAt.Time) {
		telemetree.RecordError(ctx, ErrExpiredToken{}, "refresh token expired")

		return "", ErrExpiredToken{}
	}

	// Get the user
	user, err := s.iamRepository.GetUserByID(ctx, outbound.GetUserByIDParams{ID: token.UserID})
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to get user for refresh token")

		return "", err
	}

	// Generate a new access token
	accessToken, err := s.generateAccessToken(user)
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to generate access token")

		return "", err
	}

	return accessToken, nil
}
