package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"traunseenet.com/rest-api/utils"
)

func Authenticate(context *gin.Context) {
	authHeader := context.Request.Header.Get("Authorization")
	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is required"})
		return
	}

	userId, err := utils.ValidateToken(authHeader)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid JWT token", "error": err.Error()})
		return
	}

	context.Set("userId", userId)
	context.Next()
}