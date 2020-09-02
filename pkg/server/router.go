package server

import (
	"net/http"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/business"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/config"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/metrics"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheushook"
	"github.com/sirupsen/logrus"
)

// GenerateRouter Generate main router
func GenerateRouter(logger *logrus.Logger, cfg *config.Config, metricsCtx metrics.Instance) (http.Handler, error) {
	// Set release mod
	gin.SetMode(gin.ReleaseMode)
	// Create router
	router := gin.New()
	// Add middlewares
	router.Use(gin.Recovery())
	router.Use(helmet.Default())
	router.Use(requestID(logger))
	router.Use(logMiddleware(logger))
	router.Use(metricsCtx.Instrument())
	// Add routes
	router.POST("/prometheus/webhook", func(c *gin.Context) {
		var alerts prometheushook.PrometheusAlertHook
		// Try to map data
		err := c.ShouldBindJSON(&alerts)
		// Check if error exists
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Validate input
		err = alerts.Validate()
		if err != nil {
			handleError(c, err)
			return
		}
		businessCtx, err := business.NewContext(cfg, metricsCtx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = businessCtx.ManageHook(&alerts)
		if err != nil {
			handleError(c, err)
			return
		}
		c.Status(http.StatusNoContent)
	})
	return router, nil
}
