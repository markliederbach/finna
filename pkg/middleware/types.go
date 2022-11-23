package middleware

import "github.com/gin-gonic/gin"

// Middleware creates a new middleware handler for Gin.
type Middleware func() (gin.HandlerFunc, error)

type ContextKey string

// String returns the string representation of the context key.
func (c ContextKey) String() string {
	return string(c)
}
