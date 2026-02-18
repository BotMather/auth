package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/JscorpTech/auth/internal/config"
	"github.com/JscorpTech/auth/internal/dto"
	"github.com/JscorpTech/auth/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AuthMiddleware(cfg *config.Config, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenRaw := c.Request.Header.Get("Authorization")
		token := strings.Replace(tokenRaw, "Bearer ", "", 1)
		claims, err := utils.VerifyJWT(token, cfg.PublicKey)

		if err != nil {
			dto.JSON(c, http.StatusUnauthorized, nil, err.Error())
			c.Abort()
			return
		}

		exp, ok := claims["exp"]
		if !ok {
			dto.JSON(c, http.StatusUnauthorized, nil, "Invalid token")
			c.Abort()
			return
		}
		if exp.(float64) < float64(time.Now().Unix()) {
			dto.JSON(c, http.StatusUnauthorized, nil, "Token expired")
			c.Abort()
			return
		}
		if claims["token_type"] != "access" {
			dto.JSON(c, http.StatusUnauthorized, nil, "Invalid token")
			c.Abort()
			return
		}
		c.Set("user", claims)
		c.Next()
	}
}
