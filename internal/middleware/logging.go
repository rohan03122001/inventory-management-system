package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

type logMiddleware struct {
}

func NewLogMiddleware() *logMiddleware {
	return &logMiddleware{}
}

// logs request details

func (m *logMiddleware) Logger() gin.HandlerFunc{
	return func(c *gin.Context) {
		//start timer

		start:= time.Now()

		c.Next()

		duration := time.Since(start)
		status:=c.Writer.Status()
		path:= c.Request.URL.Path
		method:= c.Request.Method

		
		println("[API]", method, path, status, duration)
	}
}