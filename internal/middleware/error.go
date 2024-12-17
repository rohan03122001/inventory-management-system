package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorMiddleware struct{}

func NewErrorMiddleware() *ErrorMiddleware {
	return &ErrorMiddleware{}
}

func (m *ErrorMiddleware) ErrorHandler() gin.HandlerFunc{
	return func(c *gin.Context) {
		defer func ()  {
				if err:= recover(); err!= nil{
					println("Panic Recovered", err)
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": "An internal error occurred",
					})
					c.Abort()
				}
		}()

		c.Next()
	}
}