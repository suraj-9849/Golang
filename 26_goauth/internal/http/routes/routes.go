package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suraj-9849/Golang/internal/http/handlers"
	"github.com/suraj-9849/Golang/internal/middleware"
)

func SetupRoutes(r *gin.Engine, handler *handlers.Handler) {
	r.GET("/", handler.Home)
	r.POST("/signup", handler.Signup)
	r.POST("/login", handler.Login)
	protected := r.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/profile", handler.GetProfile)
		protected.GET("/dashboard", handler.GetDashboard)
	}
}
