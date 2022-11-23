package middleware

import (
	"math"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	// ContextKeyLogger is the key used to store the logger in the context
	ContextKeyLogger ContextKey = "logger"
)

// InjectLogger creates a contextual logger containing the request ID
func InjectLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Context logger with default fields every request should have
		contextLogger := logrus.WithFields(logrus.Fields{
			ContextKeyRequestID.String(): c.GetString(ContextKeyRequestID.String()),
		})
		c.Set(ContextKeyLogger.String(), contextLogger)
	}
}

// RequestLogger debug logs the current request and response.
// Requires the Logger middleware to be added first.
func RequestLogger() gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	return func(c *gin.Context) {
		contextLogger := logrus.WithFields(logrus.Fields{})
		foundLogger, exists := c.Get(ContextKeyLogger.String())
		if exists {
			contextLogger = foundLogger.(*logrus.Entry)
		}
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
			"latency_ms": latency, // time to process
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
