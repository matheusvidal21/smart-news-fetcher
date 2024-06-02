package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvidal21/smart-news-fetcher/internal/auth"
	"net/http"
	"strings"
)

func AuthMiddleware(jwtService auth.JWTServiceInterface) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			context.Abort()
			return
		}

		token := strings.Split(authHeader, "Bearer ")[1]
		if token == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			context.Abort()
			return
		}

		_, err := jwtService.ValidateToken(token)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
