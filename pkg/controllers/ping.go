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
	return func(ctx *gin.Context) {
		logger := ctx.MustGet("logger").(*logrus.Entry)
		logger.Info("ping")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
