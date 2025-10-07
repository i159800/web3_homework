package middleware

import (
  "fmt"
  "github.com/gin-gonic/gin"
	"github.com/i159800/web3_homework/go_study_task/task4/utils"
	"net/http"
	"strings"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		fmt.Println(tokenString)
		claims, err := utils.ParseToken(tokenString)
		if err == nil {
			c.Set("UserID", claims["UserID"])
			c.Set("roles", claims["roles"])
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		}
	}
}
