package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{}

func NewLoginhandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
