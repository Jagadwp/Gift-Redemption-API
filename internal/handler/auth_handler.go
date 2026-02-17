package handler

import (
	"errors"

	"github.com/gift-redemption/internal/dto"
	"github.com/gift-redemption/internal/pkg/apperror"
	"github.com/gift-redemption/internal/pkg/response"
	"github.com/gift-redemption/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

// Login godoc
// @Summary      User login
// @Description  Authenticate user and return JWT token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body      dto.LoginRequest  true  "Login credentials"
// @Success      200   {object}  response.envelope{data=dto.LoginResponse}
// @Failure      400   {object}  response.envelope
// @Failure      401   {object}  response.envelope
// @Router       /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body", err.Error())
		return
	}

	result, err := h.authService.Login(req)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			response.Unauthorized(c, "invalid email or password")
			return
		}
		response.InternalServerError(c, "something went wrong")
		return
	}

	response.Success(c, "login successful", result)
}
