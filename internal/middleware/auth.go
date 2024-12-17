package middleware

import (
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
		// authHeader := c.GetHeader("Authorization")
		// if authHeader == ""{
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header Required"})
		// 	c.Abort()
		// 	return
		// }

		// //Extract bearer token

		// bearerToken := strings.Split(authHeader, "Bearer ")
		// if(len(bearerToken) != 2){
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token Format"})
		// 	c.Abort()
		// 	return
		// }
		// //parse And Validate token
	
		// token, err:= jwt.Parse(bearerToken[1], func(token *jwt.Token)(interface{}, error){
		// 	return []byte(m.jwtSecret),nil
		// })
		
		
		// if err!=nil || !token.Valid {
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"error": "Invalid Token",
		// 	})
		// 	c.Abort()
		// 	return
		// }

		// //add claims to context ??

		// if claims, ok := token.Claims.(jwt.MapClaims); ok{
		// 	c.Set("user_id", claims["user_id"])
        // }

        c.Next()

	}
}