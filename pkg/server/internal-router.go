package server

import (
	"net/http"

	"github.com/dimiro1/health"
	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/config"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

// GenerateInternalRouter Generate internal router
func GenerateInternalRouter(logger *logrus.Logger, cfg *config.Config, metricsCtx metrics.Instance) http.Handler {
	// Set release mod
	gin.SetMode(gin.ReleaseMode)
	// Create router
	router := gin.New()
	// Add middlewares
	router.Use(gin.Recovery())
	router.Use(metricsCtx.Instrument())
	router.Use(requestID(logger))
	router.Use(logMiddleware(logger))
	// Add routes
	healthHandler := health.NewHandler()
	router.GET("/health", gin.WrapH(healthHandler))
	h := promhttp.Handler()
	router.GET("/metrics", gin.WrapH(h))
	return router
}
