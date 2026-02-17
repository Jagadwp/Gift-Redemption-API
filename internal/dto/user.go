package dto

import (
	"time"

	"github.com/gift-redemption/internal/model"
)

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"omitempty,oneof=admin user"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Role  string `json:"role" binding:"required,oneof=admin user"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

func ToUserResponse(u model.User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      string(u.Role),
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
	}
}
