package http

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(router *gin.RouterGroup, h *AuthHandler) {
	users := router.Group("/auth")
	{
		users.POST("/login", h.Login)
		users.POST("/register", h.Register)
	}
}
