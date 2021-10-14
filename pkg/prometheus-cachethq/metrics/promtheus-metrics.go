package metrics

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type prometheusMetrics struct {
	reqCnt                *prometheus.CounterVec
	resSz                 *prometheus.SummaryVec
	reqDur                *prometheus.SummaryVec
	reqSz                 *prometheus.SummaryVec
	up                    prometheus.Gauge
	incidentManagedTotal  *prometheus.CounterVec
	componentManagedTotal *prometheus.CounterVec
}

func (ctx *prometheusMetrics) PrometheusHTTPHandler() http.Handler {
	return promhttp.Handler()
}

func (ctx *prometheusMetrics) IncrementIncidentManagedCounter(incidentStatus string, componentStatus string) {
	ctx.incidentManagedTotal.WithLabelValues(incidentStatus).Inc()
	ctx.IncrementComponentManagedCounter(componentStatus)
}

func (ctx *prometheusMetrics) IncrementComponentManagedCounter(status string) {
	ctx.componentManagedTotal.WithLabelValues(status).Inc()
}

// Instrument will instrument gin routes.
func (ctx *prometheusMetrics) Instrument(serverName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		reqSz := computeApproximateRequestSize(c.Request)

		c.Next()

		status := strconv.Itoa(c.Writer.Status())
		elapsed := float64(time.Since(start)) / float64(time.Second)
		resSz := float64(c.Writer.Size())

		ctx.reqDur.WithLabelValues(serverName).Observe(elapsed)
		ctx.reqCnt.WithLabelValues(serverName, status, c.Request.Method, c.Request.Host, c.Request.URL.Path).Inc()
		ctx.reqSz.WithLabelValues(serverName).Observe(float64(reqSz))
		ctx.resSz.WithLabelValues(serverName).Observe(resSz)
	}
}

// From https://github.com/DanielHeckrath/gin-prometheus/blob/master/gin_prometheus.go
func computeApproximateRequestSize(r *http.Request) int {
	s := 0
	if r.URL != nil {
		s = len(r.URL.Path)
	}

	s += len(r.Method)
	s += len(r.Proto)

	for name, values := range r.Header {
		s += len(name)
		for _, value := range values {
			s += len(value)
		}
	}

	s += len(r.Host)

	// N.B. r.Form and r.MultipartForm are assumed to be included in r.URL.

	if r.ContentLength != -1 {
		s += int(r.ContentLength)
	}

	return s
}

func (ctx *prometheusMetrics) register() {
	ctx.reqCnt = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
		},
		[]string{"server_name", "status_code", "method", "host", "path"},
	)
	prometheus.MustRegister(ctx.reqCnt)

	ctx.reqDur = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_request_duration_seconds",
			Help: "The HTTP request latencies in seconds.",
		},
		[]string{"server_name"},
	)
	prometheus.MustRegister(ctx.reqDur)

	ctx.reqSz = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_request_size_bytes",
			Help: "The HTTP request sizes in bytes.",
		},
		[]string{"server_name"},
	)
	prometheus.MustRegister(ctx.reqSz)

	ctx.resSz = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_response_size_bytes",
			Help: "The HTTP response sizes in bytes.",
		},
		[]string{"server_name"},
	)
	prometheus.MustRegister(ctx.resSz)

	ctx.up = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "up",
			Help: "1 = up, 0 = down",
		},
	)
	ctx.up.Set(1)
	prometheus.MustRegister(ctx.up)

	ctx.incidentManagedTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "incident_managed_total",
			Help: "How many incident managed in total ?",
		},
		[]string{"status"},
	)
	prometheus.MustRegister(ctx.incidentManagedTotal)

	ctx.componentManagedTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "component_managed_total",
			Help: "How many component managed in total ?",
		},
		[]string{"status"},
	)
	prometheus.MustRegister(ctx.componentManagedTotal)
}
