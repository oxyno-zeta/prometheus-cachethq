package tracing

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/config"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/log"
)

// Service Tracing service.
//go:generate mockgen -destination=./mocks/mock_Service.go -package=mocks github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/tracing Service
type Service interface {
	// Reload service
	Reload() error
	// Get opentracing tracer
	GetTracer() opentracing.Tracer
	// Http Gin HttpMiddleware to add trace per request
	HTTPMiddleware(getRequestID func(ctx context.Context) string) gin.HandlerFunc
}

// Trace structure.
//go:generate mockgen -destination=./mocks/mock_Trace.go -package=mocks github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/tracing Trace
type Trace interface {
	// Add tag to trace
	SetTag(key string, value interface{})
	// Get a child trace
	GetChildTrace(operationName string) Trace
	// End the trace
	Finish()
	// Get the trace ID
	GetTraceID() string
	// InjectInHTTPHeader will inject span in http header for forwarding.
	InjectInHTTPHeader(header http.Header) error
}

func New(cfgManager config.Manager, logger log.Logger) (Service, error) {
	return newService(cfgManager, logger)
}
