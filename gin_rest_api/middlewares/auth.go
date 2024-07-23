package middlewares

import (
	"net/http"
	"rest-api/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	
	// Using AbortWithStatusJSON to abort the request from moving forward
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
	}
	
	// Validate token
	userId, err := utils.VerifyToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Attach the userId in the gin context
	c.Set("userId", userId)

	c.Next()
}