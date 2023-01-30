package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	InfoPath EndpointPath = "/info"
)

type Info struct {
	Base
}

type InfoInput struct {
	Base
}

func NewInfoController(input InfoInput) *Info {
	return &Info{
		Base: input.Base,
	}
}

func (c *Info) Register(r *gin.Engine) {
	r.POST(c.Base.Path(InfoPath), c.Handler())
}

func (c *Info) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := c.MustGet("logger").(*logrus.Entry)
		logger.Info("ping")
		c.JSON(http.StatusOK, map[string]interface{}{
			"item_id":      itemID,
			"access_token": accessToken,
			"products":     strings.Split(PLAID_PRODUCTS, ","),
		})
	}
}
