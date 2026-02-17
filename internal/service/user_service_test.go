package service

import (
	"testing"

	"github.com/gift-redemption/internal/dto"
	"github.com/gift-redemption/internal/model"
	"github.com/gift-redemption/internal/pkg/apperror"
	"github.com/gift-redemption/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_GetByID_Success(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	userService := NewUserService(mockUserRepo)

	user := &model.User{
		ID:    1,
		Name:  "Test User",
		Email: "test@example.com",
		Role:  model.RoleUser,
	}

	mockUserRepo.On("FindByID", uint(1)).Return(user, nil)

	result, err := userService.GetByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.Name, result.Name)
	assert.Equal(t, user.Email, result.Email)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_GetByID_NotFound(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	userService := NewUserService(mockUserRepo)

	mockUserRepo.On("FindByID", uint(999)).Return(nil, apperror.ErrNotFound)

	result, err := userService.GetByID(999)

	assert.Error(t, err)
	assert.Equal(t, apperror.ErrNotFound, err)
	assert.Nil(t, result)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_Create_Success(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	userService := NewUserService(mockUserRepo)

	req := dto.CreateUserRequest{
		Name:     "New User",
		Email:    "new@example.com",
		Password: "password123",
		Role:     "user",
	}

	mockUserRepo.On("Create", mock.MatchedBy(func(u *model.User) bool {
		return u.Name == req.Name && u.Email == req.Email && u.Role == model.RoleUser
	})).Return(nil)

	result, err := userService.Create(req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, req.Name, result.Name)
	assert.Equal(t, req.Email, result.Email)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_Create_DuplicateEmail(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	userService := NewUserService(mockUserRepo)

	req := dto.CreateUserRequest{
		Name:     "Duplicate User",
		Email:    "duplicate@example.com",
		Password: "password123",
		Role:     "user",
	}

	mockUserRepo.On("Create", mock.MatchedBy(func(u *model.User) bool {
		return u.Email == req.Email
	})).Return(apperror.ErrDuplicateEntry)

	result, err := userService.Create(req)

	assert.Error(t, err)
	assert.Equal(t, apperror.ErrDuplicateEntry, err)
	assert.Nil(t, result)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_Update_Success(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	userService := NewUserService(mockUserRepo)

	existingUser := &model.User{
		ID:    1,
		Name:  "Old Name",
		Email: "old@example.com",
		Role:  model.RoleUser,
	}

	req := dto.UpdateUserRequest{
		Name:  "New Name",
		Email: "new@example.com",
		Role:  "admin",
	}

	mockUserRepo.On("FindByID", uint(1)).Return(existingUser, nil)
	mockUserRepo.On("Update", existingUser).Return(nil)

	result, err := userService.Update(1, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, req.Name, result.Name)
	assert.Equal(t, req.Email, result.Email)
	assert.Equal(t, "admin", result.Role)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_Delete_Success(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	userService := NewUserService(mockUserRepo)

	mockUserRepo.On("Delete", uint(1)).Return(nil)

	err := userService.Delete(1)

	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_Delete_NotFound(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	userService := NewUserService(mockUserRepo)

	mockUserRepo.On("Delete", uint(999)).Return(apperror.ErrNotFound)

	err := userService.Delete(999)

	assert.Error(t, err)
	assert.Equal(t, apperror.ErrNotFound, err)
	mockUserRepo.AssertExpectations(t)
}
