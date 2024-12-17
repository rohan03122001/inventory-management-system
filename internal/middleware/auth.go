package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authMiddleware struct {
	jwtSecret string
}

func NewAuthMiddleware(jwtSecret string) *authMiddleware {
	return &authMiddleware{
		jwtSecret: jwtSecret,
	}
}

func (m *authMiddleware) Authenticate() gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header Required"})
			c.Abort()
			return
		}

		//Extract bearer token

		bearerToken := strings.Split(authHeader, "Bearer ")
		if(len(bearerToken) != 2){
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token Format"})
			c.Abort()
			return
		}


		//parse And Validate token
	}
}