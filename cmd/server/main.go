package main

import (
	"log"

	_ "github.com/gift-redemption/docs" // swagger generated docs
	"github.com/gift-redemption/internal/config"
	"github.com/gift-redemption/internal/database"
	"github.com/gift-redemption/internal/handler"
	"github.com/gift-redemption/internal/repository"
	"github.com/gift-redemption/internal/service"
	"github.com/gift-redemption/seeds"
)

// @title           Gift Redemption API
// @version         1.0
// @description     REST API for gift redemption system

// @contact.name    API Support
// @contact.email   admin@gift-redemption.com

// @host            localhost:8080
// @BasePath        /

// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization
// @description                 Enter: Bearer {token}
func main() {
	cfg := config.Load()

	database.RunMigrations(cfg)

	db := database.NewPostgresConnection(cfg)
	seeds.Run(db)

	// repositories
	userRepo := repository.NewUserRepository(db)
	giftRepo := repository.NewGiftRepository(db)
	redemptionRepo := repository.NewRedemptionRepository(db)
	ratingRepo := repository.NewRatingRepository(db)

	// services
	authService := service.NewAuthService(userRepo, cfg)
	userService := service.NewUserService(userRepo)
	giftService := service.NewGiftService(giftRepo)
	redemptionService := service.NewRedemptionService(db, giftRepo, redemptionRepo, ratingRepo)

	// handlers
	handlers := Handlers{
		Auth:       handler.NewAuthHandler(authService),
		User:       handler.NewUserHandler(userService),
		Gift:       handler.NewGiftHandler(giftService),
		Redemption: handler.NewRedemptionHandler(redemptionService),
	}

	r := NewRouter(cfg, handlers)

	log.Printf("server running on port %s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
