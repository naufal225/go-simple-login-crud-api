package router

import (
	"github.com/gin-gonic/gin"
	"github.com/naufal225/go-simple-login-crud-api/internal/config"
	"github.com/naufal225/go-simple-login-crud-api/internal/handler"
	"github.com/naufal225/go-simple-login-crud-api/internal/middleware"
)

func SetupRouter(
	cfg *config.Config,
	authHandler *handler.AuthHandler,
	itemHandler *handler.ItemHandler,
) *gin.Engine {
	
	r := gin.Default()

	// === Public Routes ===

	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
	}

	// === Protected Routes ===

	protected := r.Group("/")
	protected.Use(middleware.JWTAuth(cfg))
	{
		protected.GET("/items", itemHandler.List)
		protected.POST("/items", itemHandler.Create)
		protected.PUT("/items/:id", itemHandler.Update)
		protected.DELETE("/items/:id", itemHandler.Delete)
	}

	return r

}