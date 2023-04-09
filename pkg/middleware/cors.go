package middleware

import (
	"github.com/gin-gonic/gin"
)

// Uuid generates a UUID for the current request and adds it to the context
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
		c.Next()
	}
}
