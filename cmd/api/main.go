package main

import (
	"math"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	loggerMiddleware, err := configureLogger()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	logrus.WithField("version", Version).Info("running finna API")

	gin.SetMode(GIN_MODE)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(uuidMiddleware())
	r.Use(loggerMiddleware)
	r.Use(ginLoggerMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		logger := c.MustGet("logger").(*logrus.Entry)
		logger.Info("ping")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err = r.Run(":" + APP_PORT)
	if err != nil {
		logrus.WithError(err).Error("failed to start API")
		os.Exit(1)
	}
}

func configureLogger() (gin.HandlerFunc, error) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	level, err := logrus.ParseLevel(LOG_LEVEL)
	if err != nil {
		return nil, err
	}
	logrus.SetLevel(level)
	return func(c *gin.Context) {
		// Context logger with default fields every request should have
		contextLogger := logrus.WithFields(logrus.Fields{
			"request_id": c.GetString("request_id"),
		})
		c.Set("logger", contextLogger)
		c.Next()
	}, nil
}

func ginLoggerMiddleware() gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	return func(c *gin.Context) {
		contextLogger := c.MustGet("logger").(*logrus.Entry)
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()

		// calculate latency in milliseconds
		latency := int(math.Ceil(float64(time.Since(start).Nanoseconds()) / 1000000.0))

		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}
		logger := contextLogger.WithFields(logrus.Fields{
			"hostname":   hostname,
			"statusCode": statusCode,
			"latency":    latency, // time to process
			"clientIP":   clientIP,
			"method":     c.Request.Method,
			"path":       path,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
		})
		if len(c.Errors) > 0 {
			logger.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
			return
		}
		if statusCode >= http.StatusInternalServerError {
			logger.Error("internal server error")
		} else if statusCode >= http.StatusBadRequest {
			logger.Warn("bad request")
		} else {
			logger.Debug("request complete")
		}
	}
}

// middleware to inject uuid into gin context
func uuidMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New()
		c.Set("request_id", uuid.String())
		c.Next()
	}
}

func getOrDefaultEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
