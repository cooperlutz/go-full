package inbound

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/cooperlutz/go-full/internal/iam/service"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

func NewIamAuthApiController(iamSvc *service.IamService) http.Handler {
	iamRouter := hteeteepee.NewRouter("iam.adapter.inbound.auth")
	authHandler := NewAuthHandler(iamSvc)
	iamRouter.HandleFunc("/register", authHandler.Register)
	iamRouter.HandleFunc("/login", authHandler.Login)
	iamRouter.HandleFunc("/refresh", authHandler.RefreshToken)
	iamRouter.HandleFunc("/logout", authHandler.Logout)

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
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)

		return
	}

	user, err := h.authService.Register(r.Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrEmailInUse{}) {
			http.Error(w, "Email already in use", http.StatusConflict)

			return
		}

		http.Error(w, "Error creating user", http.StatusInternalServerError)

		return
	}

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

// Login handles user login by setting httpOnly cookies for access and refresh tokens.
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}

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

	setTokenCookies(w, accessToken, refreshToken, h.authService.AccessTokenTTL(), h.authService.RefreshTokenTTL())

	w.WriteHeader(http.StatusNoContent)
}

// RefreshToken handles access token refresh using the refresh_token cookie.
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshCookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "Missing refresh token", http.StatusUnauthorized)

		return
	}

	accessToken, err := h.authService.RefreshAccessToken(r.Context(), refreshCookie.Value)
	if err != nil {
		if errors.Is(err, service.ErrInvalidToken{}) || errors.Is(err, service.ErrExpiredToken{}) {
			clearTokenCookies(w)
			http.Error(w, "Invalid or expired refresh token", http.StatusUnauthorized)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}

		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   int(h.authService.AccessTokenTTL().Seconds()),
	})

	w.WriteHeader(http.StatusNoContent)
}

// Logout clears authentication cookies.
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	clearTokenCookies(w)
	w.WriteHeader(http.StatusNoContent)
}

// setTokenCookies sets the access and refresh token cookies.
func setTokenCookies(w http.ResponseWriter, accessToken, refreshToken string, accessTTL, refreshTTL time.Duration) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   int(accessTTL.Seconds()),
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/auth",
		MaxAge:   int(refreshTTL.Seconds()),
	})
}

// clearTokenCookies clears the access and refresh token cookies.
func clearTokenCookies(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   -1,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/auth",
		MaxAge:   -1,
	})
}
