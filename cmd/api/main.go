package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/markliederbach/finna/pkg/controllers"
	"github.com/markliederbach/finna/pkg/middleware"
	"github.com/sirupsen/logrus"
)

var (
	Version   = "latest"
	APP_PORT  string
	LOG_LEVEL string
	GIN_MODE  string
	BASE_URL  string
)

func init() {
	APP_PORT = getOrDefaultEnv("APP_PORT", "8784")
	LOG_LEVEL = getOrDefaultEnv("LOG_LEVEL", "INFO")
	GIN_MODE = getOrDefaultEnv("GIN_MODE", gin.ReleaseMode)
	BASE_URL = getOrDefaultEnv("BASE_URL", fmt.Sprintf("http://127.0.0.1:%s", APP_PORT))
}

const (
	APIPathPrefix      = "/api"
	FrontendPathPrefix = "/"
)

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
	r.Use(middleware.Cors())
	r.Use(middleware.Uuid())
	r.Use(middleware.InjectLogger())
	r.Use(middleware.RequestLogger())

	// initialize endpoint controllers
	baseAPIController := controllers.Base{
		PrefixPath: APIPathPrefix,
	}

	baseFrontendController := controllers.Base{
		PrefixPath: FrontendPathPrefix,
	}
	endpoints := []controllers.Endpoint{
		// API endpoints
		controllers.NewPingController(controllers.PingInput{Base: baseAPIController}),
		controllers.NewFormatController(controllers.FormatInput{Base: baseAPIController}),

		// Frontend endpoints
		controllers.NewFrontendController(controllers.FrontendInput{Base: baseFrontendController, BaseUrl: BASE_URL}),
	}

	for _, endpoint := range endpoints {
		endpoint.Register(r)
	}

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
