package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	PingPath EndpointPath = "/ping"
)

type Ping struct {
	Base
}

type PingInput struct {
	Base
}

func NewPingController(input PingInput) *Ping {
	return &Ping{
		Base: input.Base,
	}
}

func (c *Ping) Register(r *gin.Engine) {
	r.GET(c.Base.Path(PingPath), c.Handler())
}

func (c *Ping) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := c.MustGet("logger").(*logrus.Entry)
		logger.Info("ping")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
