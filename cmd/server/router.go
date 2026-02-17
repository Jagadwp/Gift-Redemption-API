package main

import (
	"github.com/gift-redemption/internal/config"
	"github.com/gift-redemption/internal/handler"
	"github.com/gift-redemption/internal/middleware"
	"github.com/gift-redemption/internal/model"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handlers struct {
	Auth       *handler.AuthHandler
	User       *handler.UserHandler
	Gift       *handler.GiftHandler
	Redemption *handler.RedemptionHandler
}

func NewRouter(cfg *config.Config, h Handlers) *gin.Engine {
	r := gin.Default()

	// swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := middleware.Authenticate(cfg)
	adminOnly := middleware.RequireRole(model.RoleAdmin)

	r.POST("/login", h.Auth.Login)

	gifts := r.Group("/gifts", auth)
	{
		gifts.GET("", h.Gift.GetAll)
		gifts.GET("/:id", h.Gift.GetByID)
		gifts.POST("", adminOnly, h.Gift.Create)
		gifts.PUT("/:id", adminOnly, h.Gift.Update)
		gifts.PATCH("/:id", adminOnly, h.Gift.Patch)
		gifts.DELETE("/:id", adminOnly, h.Gift.Delete)
		gifts.POST("/:id/redeem", h.Redemption.Redeem)
		gifts.POST("/:id/rating", h.Redemption.Rate)
	}

	users := r.Group("/users", auth, adminOnly)
	{
		users.GET("", h.User.GetAll)
		users.GET("/:id", h.User.GetByID)
		users.POST("", h.User.Create)
		users.PUT("/:id", h.User.Update)
		users.DELETE("/:id", h.User.Delete)
	}

	return r
}
