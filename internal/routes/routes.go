package routes

import (
	"github.com/JscorpTech/auth/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			authHandler := handlers.NewLoginhandler()
			auth.POST("/login", authHandler.Login)
		}
	}

}
