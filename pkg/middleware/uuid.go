package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	// ContextKeyRequestID is the key used to store the request ID in the context
	ContextKeyRequestID ContextKey = "request_id"
)

// Uuid generates a UUID for the current request and adds it to the context
func Uuid() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New()
		c.Set(ContextKeyRequestID.String(), uuid.String())
	}
}
