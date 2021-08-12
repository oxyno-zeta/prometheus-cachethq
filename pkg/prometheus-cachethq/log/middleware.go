package log

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/common/utils"
)

type logContextKey struct {
	name string
}

const loggerGinCtxKey = "LoggerCtxKey"

var loggerCtxKey = &logContextKey{name: "logger"}

const nsToMs = 1000000.0

func GetLoggerFromContext(ctx context.Context) Logger {
	logger, _ := ctx.Value(loggerCtxKey).(Logger)

	return logger
}

func GetLoggerFromGin(c *gin.Context) Logger {
	val, _ := c.Get(loggerGinCtxKey)
	logger, _ := val.(Logger)

	return logger
}

func SetLoggerToContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerCtxKey, logger)
}

func SetLoggerToGin(c *gin.Context, logger Logger) {
	c.Set(loggerGinCtxKey, logger)
}

func Middleware(logger Logger, getRequestID func(c *gin.Context) string, getSpanID func(ctx context.Context) string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t1 := time.Now()
		// Get request
		r := c.Request

		// Create logger fields
		logFields := make(map[string]interface{})

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

		logFields["uri"] = utils.GetRequestURL(c.Request)

		// Log request id
		logFields["request_id"] = getRequestID(c)

		// Get span id
		spanID := getSpanID(c.Request.Context())
		if spanID != "" {
			logFields["span_id"] = spanID
		}

		requestLogger := logger.WithFields(logFields)

		requestLogger.Debugln("request started")

		// Add logger to request
		SetLoggerToGin(c, requestLogger)
		c.Request = c.Request.WithContext(SetLoggerToContext(c.Request.Context(), requestLogger))

		// Next
		c.Next()

		// Get status
		status := c.Writer.Status()
		bytes := c.Writer.Size()

		// Create new fields
		endFields := map[string]interface{}{
			"resp_status":       status,
			"resp_bytes_length": bytes,
			"resp_elapsed_ms":   float64(time.Since(t1).Nanoseconds()) / nsToMs,
		}

		endRequestLogger := requestLogger.WithFields(endFields)

		logFunc := endRequestLogger.Infoln

		if status >= http.StatusMultipleChoices && status < http.StatusInternalServerError {
			logFunc = endRequestLogger.Warnln
		}

		if status >= http.StatusInternalServerError {
			logFunc = endRequestLogger.Errorln
		}

		logFunc("request complete")
	}
}
