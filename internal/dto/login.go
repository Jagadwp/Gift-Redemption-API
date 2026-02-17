package dto

import "github.com/gift-redemption/internal/model"

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

func ToLoginResponse(token string, user model.User) LoginResponse {
	return LoginResponse{
		Token: token,
		User:  ToUserResponse(user),
	}
}
