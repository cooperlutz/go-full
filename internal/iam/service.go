// auth/service.go
package iam

import (
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/cooperlutz/go-full/internal/iam/models"
	"github.com/cooperlutz/go-full/pkg/securitee"
)

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
	userRepo         *models.UserRepository
	refreshTokenRepo *models.RefreshTokenRepository
	jwtSecret        []byte
	accessTokenTTL   time.Duration
}

// NewIamService creates a new authentication service.
func NewIamService(userRepo *models.UserRepository, refreshTokenRepo *models.RefreshTokenRepository, jwtSecret string, accessTokenTTL time.Duration) *IamService {
	return &IamService{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
		jwtSecret:        []byte(jwtSecret),
		accessTokenTTL:   accessTokenTTL,
	}
}

// Register creates a new user with the provided credentials.
func (s *IamService) Register(email, password string) (*models.User, error) {
	// Check if user already exists
	_, err := s.userRepo.GetUserByEmail(email)
	if err == nil {
		return nil, ErrEmailInUse{}
	}

	// Only proceed if the error was "user not found"
	if !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	// Hash the password
	hashedPassword, err := securitee.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create the user
	user, err := s.userRepo.CreateUser(email, hashedPassword)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// generateAccessToken creates a new JWT access token.
func (s *IamService) generateAccessToken(user *models.User) (string, error) {
	// Set the expiration time
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
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
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
func (s *IamService) LoginWithRefresh(email, password string, refreshTokenTTL time.Duration) (accessToken, refreshToken string, err error) {
	// Get the user from the database
	user, err := s.userRepo.GetUserByEmail(email)
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

	// Create a refresh token
	token, err := s.refreshTokenRepo.CreateRefreshToken(user.ID, refreshTokenTTL)
	if err != nil {
		return "", "", err
	}

	return accessToken, token.Token, nil
}

// RefreshAccessToken creates a new access token using a refresh token.
func (s *IamService) RefreshAccessToken(refreshTokenString string) (string, error) {
	// Retrieve the refresh token
	token, err := s.refreshTokenRepo.GetRefreshToken(refreshTokenString)
	if err != nil {
		return "", ErrInvalidToken{}
	}

	// Check if the token is valid
	if token.Revoked {
		return "", ErrInvalidToken{}
	}

	// Check if the token has expired
	if time.Now().After(token.ExpiresAt) {
		return "", ErrExpiredToken{}
	}

	// Get the user
	user, err := s.userRepo.GetUserByID(token.UserID)
	if err != nil {
		return "", err
	}

	// Generate a new access token
	accessToken, err := s.generateAccessToken(user)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
