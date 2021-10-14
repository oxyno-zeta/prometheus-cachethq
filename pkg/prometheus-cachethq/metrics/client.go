package metrics

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Client Client metrics interface.
//go:generate mockgen -destination=./mocks/mock_Client.go -package=mocks github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/metrics Client
type Client interface {
	// Instrument web server.
	Instrument(serverName string) gin.HandlerFunc
	// Get prometheus handler for http expose.
	PrometheusHTTPHandler() http.Handler
	// IncrementIncidentManagedCounter Increment incident managed counter.
	IncrementIncidentManagedCounter(incidentStatus string, componentStatus string)
	// IncrementComponentManagedCounter Increment component managed counter.
	IncrementComponentManagedCounter(status string)
}

// NewMetricsClient will generate a new Client.
func NewMetricsClient() Client {
	ctx := &prometheusMetrics{}
	// Register
	ctx.register()

	return ctx
}
