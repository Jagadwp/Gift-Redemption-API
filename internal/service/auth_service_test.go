package service

import (
	"testing"

	"github.com/gift-redemption/internal/config"
	"github.com/gift-redemption/internal/dto"
	"github.com/gift-redemption/internal/model"
	"github.com/gift-redemption/internal/pkg/apperror"
	"github.com/gift-redemption/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_Login_Success(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	cfg := &config.Config{
		JWT: config.JWTConfig{
			Secret:      "test-secret",
			ExpiryHours: 24,
		},
	}
	authService := NewAuthService(mockUserRepo, cfg)

	user := &model.User{
		ID:    1,
		Name:  "Test User",
		Email: "test@example.com",
		Role:  model.RoleUser,
	}
	_ = user.HashPassword("password123")

	mockUserRepo.On("FindByEmail", "test@example.com").Return(user, nil)

	req := dto.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	result, err := authService.Login(req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Token)
	assert.Equal(t, user.ID, result.User.ID)
	assert.Equal(t, user.Email, result.User.Email)
	mockUserRepo.AssertExpectations(t)
}

func TestAuthService_Login_WrongPassword(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	cfg := &config.Config{
		JWT: config.JWTConfig{
			Secret:      "test-secret",
			ExpiryHours: 24,
		},
	}
	authService := NewAuthService(mockUserRepo, cfg)

	user := &model.User{
		ID:    1,
		Email: "test@example.com",
	}
	_ = user.HashPassword("correctpassword")

	mockUserRepo.On("FindByEmail", "test@example.com").Return(user, nil)

	req := dto.LoginRequest{
		Email:    "test@example.com",
		Password: "wrongpassword",
	}

	result, err := authService.Login(req)

	assert.Error(t, err)
	assert.Equal(t, apperror.ErrNotFound, err)
	assert.Nil(t, result)
	mockUserRepo.AssertExpectations(t)
}

func TestAuthService_Login_UserNotFound(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	cfg := &config.Config{
		JWT: config.JWTConfig{
			Secret:      "test-secret",
			ExpiryHours: 24,
		},
	}
	authService := NewAuthService(mockUserRepo, cfg)

	mockUserRepo.On("FindByEmail", "nonexistent@example.com").Return(nil, apperror.ErrNotFound)

	req := dto.LoginRequest{
		Email:    "nonexistent@example.com",
		Password: "password123",
	}

	result, err := authService.Login(req)

	assert.Error(t, err)
	assert.Equal(t, apperror.ErrNotFound, err)
	assert.Nil(t, result)
	mockUserRepo.AssertExpectations(t)
}
