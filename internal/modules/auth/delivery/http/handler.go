package http

import (
	"net/http"

	"github.com/JscorpTech/auth/internal/modules/auth"
	"github.com/JscorpTech/auth/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	usecase auth.AuthUsecase
}

func NewAuthHandler(usecase auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		usecase: usecase,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var userPayload auth.AuthRegisterRequest
	if err := c.ShouldBindJSON(&userPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": utils.FormatValidationErrors(err, &userPayload)})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "register api",
	})
}
