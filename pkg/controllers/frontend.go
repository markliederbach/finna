package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Frontend struct {
	Base
}

type FrontendInput struct {
	Base
}

func NewFrontendController(input FrontendInput) *Frontend {
	return &Frontend{
		Base: input.Base,
	}
}

func (c *Frontend) Register(r *gin.Engine) {
	r.GET(c.Base.Path("/"), c.Handler())
}

func (c *Frontend) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := c.MustGet("logger").(*logrus.Entry)
		logger.Info("frontend")
		c.File("frontend/index.html")
	}
}
