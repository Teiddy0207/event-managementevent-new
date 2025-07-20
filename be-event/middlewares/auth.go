package middlewares

import (
	"be-event/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"fmt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Thiếu Authorization header"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token không hợp lệ"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		userID, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token hết hạn hoặc không hợp lệ"})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
