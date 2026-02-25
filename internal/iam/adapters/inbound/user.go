package inbound

import (
	"encoding/json"
	"net/http"

	"github.com/cooperlutz/go-full/internal/iam/service"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
	"github.com/cooperlutz/go-full/pkg/securitee"
)

func NewIamUserApiController(iamSvc *service.IamService) http.Handler {
	iamUserRouter := hteeteepee.NewRouter("iam.user")
	handler := NewUserHandler(iamSvc)
	iamUserRouter.HandleFunc("/profile", handler.Profile)

	return iamUserRouter
}

// UserHandler contains HTTP handlers for user-related endpoints.
type UserHandler struct {
	userRepo service.IIamQueries
}

// NewUserHandler creates a new user handler.
func NewUserHandler(iamSvc *service.IamService) *UserHandler {
	return &UserHandler{
		userRepo: iamSvc.Queries,
	}
}

// UserResponse represents the user data returned to clients.
type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// Profile returns the authenticated user's profile.
func (h *UserHandler) Profile(w http.ResponseWriter, r *http.Request) {
	// Get user ID from request context (set by auth middleware)
	userID, ok := securitee.GetUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)

		return
	}

	// Get user from database
	user, err := h.userRepo.FindUserByID(r.Context(), userID.String())
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)

		return
	}

	// Return user profile (excluding sensitive data)
	response := UserResponse{
		ID:    user.ID.String(),
		Email: user.Email,
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}
}
