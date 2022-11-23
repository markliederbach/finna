package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/markliederbach/finna/pkg/middleware"
	"github.com/sirupsen/logrus"
)

var (
	Version   = "latest"
	APP_PORT  string
	LOG_LEVEL string
	GIN_MODE  string
)

func init() {
	APP_PORT = getOrDefaultEnv("APP_PORT", "8080")
	LOG_LEVEL = getOrDefaultEnv("LOG_LEVEL", "INFO")
	GIN_MODE = getOrDefaultEnv("GIN_MODE", gin.ReleaseMode)
}

func main() {
	err := configureLogger()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	logrus.WithField("version", Version).Info("running finna API")

	gin.SetMode(GIN_MODE)
	r := gin.New()
	r.Use(gin.Recovery())

	// custom middleware
	r.Use(middleware.Uuid())
	r.Use(middleware.InjectLogger())
	r.Use(middleware.RequestLogger())

	// routes
	r.GET("/ping", func(c *gin.Context) {
		logger := c.MustGet("logger").(*logrus.Entry)
		logger.Info("ping")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// start server
	err = r.Run(":" + APP_PORT)
	if err != nil {
		logrus.WithError(err).Error("failed to start API")
		os.Exit(1)
	}
}

func configureLogger() error {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	level, err := logrus.ParseLevel(LOG_LEVEL)
	if err != nil {
		return err
	}
	logrus.SetLevel(level)
	return nil
}

func getOrDefaultEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
