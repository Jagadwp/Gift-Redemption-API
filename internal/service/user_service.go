package service

import (
	"github.com/gift-redemption/internal/dto"
	"github.com/gift-redemption/internal/model"
	"github.com/gift-redemption/internal/repository"
)

type UserService interface {
	GetAll() ([]dto.UserResponse, error)
	GetByID(id uint) (*dto.UserResponse, error)
	Create(req dto.CreateUserRequest) (*dto.UserResponse, error)
	Update(id uint, req dto.UpdateUserRequest) (*dto.UserResponse, error)
	Delete(id uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) GetAll() ([]dto.UserResponse, error) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	result := make([]dto.UserResponse, len(users))
	for i, u := range users {
		result[i] = dto.ToUserResponse(u)
	}
	return result, nil
}

func (s *userService) GetByID(id uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	res := dto.ToUserResponse(*user)
	return &res, nil
}

func (s *userService) Create(req dto.CreateUserRequest) (*dto.UserResponse, error) {
	role := model.RoleUser
	if req.Role == string(model.RoleAdmin) {
		role = model.RoleAdmin
	}

	user := &model.User{
		Name:  req.Name,
		Email: req.Email,
		Role:  role,
	}

	if err := user.HashPassword(req.Password); err != nil {
		return nil, err
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	res := dto.ToUserResponse(*user)
	return &res, nil
}

func (s *userService) Update(id uint, req dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Name = req.Name
	user.Email = req.Email
	user.Role = model.UserRole(req.Role)

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	res := dto.ToUserResponse(*user)
	return &res, nil
}

func (s *userService) Delete(id uint) error {
	return s.userRepo.Delete(id)
}
