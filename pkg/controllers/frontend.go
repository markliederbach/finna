package controllers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Frontend struct {
	Base
	BaseUrl string
}

type FrontendInput struct {
	Base
	BaseUrl string
}

func NewFrontendController(input FrontendInput) *Frontend {
	return &Frontend{
		Base:    input.Base,
		BaseUrl: input.BaseUrl,
	}
}

func (c *Frontend) Register(r *gin.Engine) {
	r.LoadHTMLFiles("frontend/index.tmpl.html")
	r.GET(c.Base.Path("/"), c.Handler())
}

func (c *Frontend) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger := ctx.MustGet("logger").(*logrus.Entry)
		logger.Info("frontend")
		ctx.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"base_url": template.URL(c.BaseUrl),
		})
	}
}
