package inbound

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cooperlutz/go-full/internal/iam/service"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

func NewIamAuthApiController(iamSvc *service.IamService) http.Handler {
	iamRouter := hteeteepee.NewRouter("iam.adapter.inbound.auth")
	authHandler := NewAuthHandler(iamSvc)
	iamRouter.HandleFunc("/register", authHandler.Register)
	iamRouter.HandleFunc("/login", authHandler.Login)
	iamRouter.HandleFunc("/refresh", authHandler.RefreshToken)

	return iamRouter
}

// AuthHandler contains HTTP handlers for authentication.
type AuthHandler struct {
	authService *service.IamService
}

// NewAuthHandler creates a new auth handler.
func NewAuthHandler(authService *service.IamService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// RegisterRequest represents the registration payload.
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterResponse contains the user data after successful registration.
type RegisterResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// Register handles user registration.
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}

	// Validate input
	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)

		return
	}

	// Call the auth service to register the user
	user, err := h.authService.Register(r.Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrEmailInUse{}) {
			http.Error(w, "Email already in use", http.StatusConflict)

			return
		}

		http.Error(w, "Error creating user", http.StatusInternalServerError)

		return
	}

	// Return the created user (without sensitive data)
	response := RegisterResponse{
		ID:    user.ID.String(),
		Email: user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}
}

// LoginRequest represents the login payload.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Update the LoginResponse to include refresh token
// LoginResponse contains the JWT token and refresh token after successful login.
type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// Update the Login function.
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}

	// Attempt to login with refresh token generation
	accessToken, refreshToken, err := h.authService.Login(
		r.Context(),
		req.Email,
		req.Password,
	)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials{}) {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}

		return
	}

	// Return the tokens
	response := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}
}

// RefreshRequest represents the refresh token payload.
type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

// RefreshResponse contains the new access token.
type RefreshResponse struct {
	Token string `json:"token"`
}

// RefreshToken handles access token refresh.
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var req RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}

	// Attempt to refresh the token
	token, err := h.authService.RefreshAccessToken(r.Context(), req.RefreshToken)
	if err != nil {
		if errors.Is(err, service.ErrInvalidToken{}) || errors.Is(err, service.ErrExpiredToken{}) {
			http.Error(w, "Invalid or expired refresh token", http.StatusUnauthorized)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}

		return
	}

	// Return the new access token
	response := RefreshResponse{Token: token}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}
}
