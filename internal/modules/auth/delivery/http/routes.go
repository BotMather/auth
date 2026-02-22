package http

import (
	"github.com/JscorpTech/auth/internal/config"
	"github.com/JscorpTech/auth/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(cfg *config.Config, router *gin.RouterGroup, h *AuthHandler) {
	public := router.Group("")
	{
		public.POST("/login", h.Login)
		public.POST("/register", h.Register)
		public.POST("/refresh", h.RefreshToken)
		public.POST("/confirm", h.Confirm)
		public.POST("/google", h.Google)
	}
	private := router.Group("")
	private.Use(middlewares.AuthMiddleware(cfg, h.logger))
	{
		private.GET("/me", h.Me)
	}
}
