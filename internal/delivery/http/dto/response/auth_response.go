// 7. AUTH DTOs - RESPONSE
// File: internal/delivery/http/dto/response/auth_response.go

package response

import (
	"github.com/google/uuid"
	"github.com/haninhammoud01/go-academic-service/internal/domain/entity"
)

type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
	IsActive bool      `json:"is_active"`
}

func ToUserResponse(user *entity.User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		IsActive: user.IsActive,
	}
}
