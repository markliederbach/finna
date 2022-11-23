package controllers

import (
	"github.com/gin-gonic/gin"
)

type EndpointPath string

type Base struct {
	PrefixPath string
}

// Path returns the full path for the endpoint, including the prefix.
func (b *Base) Path(endpoint EndpointPath) string {
	return b.PrefixPath + string(endpoint)
}

type Endpoint interface {
	// Register registers any routes for the endpoint on the provided gin.Engine
	Register(r *gin.Engine)
}
