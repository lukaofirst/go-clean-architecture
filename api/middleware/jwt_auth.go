package middlewares

import (
	"net/http"
	"strings"

	"go-clean-architecture/internal/auth"

	"github.com/gin-gonic/gin"
)

func JWTAuth(cfg auth.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if h == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			return
		}

		const prefix = "Bearer "
		if !strings.HasPrefix(h, prefix) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header (expected Bearer token)"})
			return
		}

		tokenString := strings.TrimPrefix(h, prefix)

		claims, err := auth.ValidateAccessToken(cfg, tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		// Make user info available to handlers
		c.Set("userId", claims.UserID)

		c.Next()
	}
}
