package tracing

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/version"
)

func (s *service) HTTPMiddleware(getRequestID func(ctx context.Context) string) gin.HandlerFunc {
	// Get version
	v := version.GetVersion()
	// Add more metadata to span
	opt := ginhttp.MWSpanObserver(func(span opentracing.Span, r *http.Request) {
		// Add request host
		span.SetTag("http.request_host", r.Host)
		// Add request id
		span.SetTag("http.request_id", getRequestID(r.Context()))
		// Add request path
		span.SetTag("http.request_path", r.URL.Path)
		// Add version
		span.SetTag("service.version", v.Version)
		// Add service name
		span.SetTag("service.name", "prometheus-cachethq")
	})

	return ginhttp.Middleware(s.tracer, opt)
}
