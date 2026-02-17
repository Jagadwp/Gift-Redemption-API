package service

import (
	"fmt"
	"time"

	"github.com/gift-redemption/internal/config"
	"github.com/gift-redemption/internal/dto"
	"github.com/gift-redemption/internal/pkg/apperror"
	"github.com/gift-redemption/internal/repository"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Login(req dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	userRepo repository.UserRepository
	cfg      *config.Config
}

func NewAuthService(userRepo repository.UserRepository, cfg *config.Config) AuthService {
	return &authService{userRepo, cfg}
}

func (s *authService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, apperror.ErrNotFound
	}

	if !user.CheckPassword(req.Password) {
		return nil, apperror.ErrNotFound
	}

	token, err := s.generateToken(user.ID, string(user.Role))
	if err != nil {
		return nil, fmt.Errorf("generate token: %w", err)
	}

	return &dto.LoginResponse{
		Token: token,
		User: dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      string(user.Role),
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
		},
	}, nil
}

func (s *authService) generateToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Duration(s.cfg.JWT.ExpiryHours) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.JWT.Secret))
}
