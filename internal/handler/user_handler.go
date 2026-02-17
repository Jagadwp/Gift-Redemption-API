package handler

import (
	"errors"
	"strconv"

	"github.com/gift-redemption/internal/dto"
	"github.com/gift-redemption/internal/pkg/apperror"
	"github.com/gift-redemption/internal/pkg/response"
	"github.com/gift-redemption/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

// GetUsers godoc
// @Summary      Get all users
// @Description  Returns list of all users (admin only)
// @Tags         Users
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  response.envelope{data=[]dto.UserResponse}
// @Failure      403  {object}  response.envelope
// @Router       /users [get]
func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.userService.GetAll()
	if err != nil {
		response.InternalServerError(c, "failed to fetch users")
		return
	}
	response.Success(c, "users retrieved successfully", users)
}

// GetUser godoc
// @Summary      Get user by ID
// @Description  Returns a single user by ID (admin only)
// @Tags         Users
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  response.envelope{data=dto.UserResponse}
// @Failure      404  {object}  response.envelope
// @Router       /users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	user, err := h.userService.GetByID(id)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			response.NotFound(c, "user not found")
			return
		}
		response.InternalServerError(c, "failed to fetch user")
		return
	}

	response.Success(c, "user retrieved successfully", user)
}

// CreateUser godoc
// @Summary      Create user
// @Description  Create a new user (admin only)
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      dto.CreateUserRequest  true  "User data"
// @Success      201   {object}  response.envelope{data=dto.UserResponse}
// @Failure      400   {object}  response.envelope
// @Failure      409   {object}  response.envelope
// @Router       /users [post]
func (h *UserHandler) Create(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body", err.Error())
		return
	}

	user, err := h.userService.Create(req)
	if err != nil {
		if errors.Is(err, apperror.ErrDuplicateEntry) {
			response.UnprocessableEntity(c, "email already registered", nil)
			return
		}
		response.InternalServerError(c, "failed to create user")
		return
	}

	response.Created(c, "user created successfully", user)
}

// UpdateUser godoc
// @Summary      Update user
// @Description  Update user data (admin only)
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int                    true  "User ID"
// @Param        body  body      dto.UpdateUserRequest  true  "User data"
// @Success      200   {object}  response.envelope{data=dto.UserResponse}
// @Failure      400   {object}  response.envelope
// @Failure      404   {object}  response.envelope
// @Router       /users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body", err.Error())
		return
	}

	user, err := h.userService.Update(id, req)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			response.NotFound(c, "user not found")
			return
		}
		response.InternalServerError(c, "failed to update user")
		return
	}

	response.Success(c, "user updated successfully", user)
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Soft delete a user (admin only)
// @Tags         Users
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  response.envelope
// @Failure      404  {object}  response.envelope
// @Router       /users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	if err := h.userService.Delete(id); err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			response.NotFound(c, "user not found")
			return
		}
		response.InternalServerError(c, "failed to delete user")
		return
	}

	response.Success(c, "user deleted successfully", nil)
}

func parseID(c *gin.Context, param string) (uint, error) {
	raw, err := strconv.ParseUint(c.Param(param), 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid id parameter", nil)
		return 0, err
	}
	return uint(raw), nil
}
