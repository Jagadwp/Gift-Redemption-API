package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	docs "github.com/gift-redemption/docs" // swagger generated docs
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

	if cfg.AppHost != "" {
		docs.SwaggerInfo.Host = cfg.AppHost
		if cfg.AppEnv == "production" {
			docs.SwaggerInfo.Schemes = []string{"https"}
		}
	}

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

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: r,
	}

	go func() {
		log.Printf("server running on port %s", cfg.AppPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown: allow in-flight requests to complete.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("server shutdown error: %v", err)
	}

	if sqlDB, err := db.DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Printf("db close error: %v", err)
		}
	}
}
