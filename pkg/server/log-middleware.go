package server

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func logMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		t1 := time.Now()
		// Get request
		r := c.Request

		// Create logger fields
		logFields := logrus.Fields{}

		// Check if it is http or https
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}
		logFields["http_scheme"] = scheme
		logFields["http_proto"] = r.Proto
		logFields["http_method"] = r.Method

		logFields["remote_addr"] = r.RemoteAddr
		logFields["user_agent"] = r.UserAgent()
		logFields["client_ip"] = c.ClientIP()

		logFields["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

		requestLogger := logger.WithFields(logFields)

		requestLogger.Debugln("request started")

		// Next
		c.Next()

		// Get status
		status := c.Writer.Status()
		bytes := c.Writer.Size()
		endRequestLogger := requestLogger.WithFields(logrus.Fields{
			"resp_status":       status,
			"resp_bytes_length": bytes,
			"resp_elapsed_ms":   float64(time.Since(t1).Nanoseconds()) / 1000000.0,
		})
		logFunc := endRequestLogger.Infoln
		if status >= 300 && status < 400 {
			logFunc = endRequestLogger.Warnln
		}
		if status >= 400 {
			logFunc = endRequestLogger.Errorln
		}

		logFunc("request complete")
	}
}
