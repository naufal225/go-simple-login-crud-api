package main

import (
	"log"

	"github.com/naufal225/go-simple-login-crud-api/internal/config"
	"github.com/naufal225/go-simple-login-crud-api/internal/db"
	"github.com/naufal225/go-simple-login-crud-api/internal/handler"
	"github.com/naufal225/go-simple-login-crud-api/internal/repo"
	"github.com/naufal225/go-simple-login-crud-api/internal/router"
	"github.com/naufal225/go-simple-login-crud-api/internal/service"
)

func main() {
	// Load Config
	cfg := config.Load()

	// DB
	dbConn := db.Connect(cfg)

	// Repositories
	userRepo := repo.NewUserRepository(dbConn)
	itemRepo := repo.NewItemRepository(dbConn)

	// Services
	authService := service.NewAuthService(userRepo, cfg)
	itemService := service.NewItemService(itemRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authService)
	itemHandler := handler.NewItemHandler(itemService)

	// Router
	r := router.SetupRouter(cfg, authHandler, itemHandler)

	log.Printf("🚀 server running on port %s", cfg.AppPort)

	r.Run(":" + cfg.AppPort)
}
