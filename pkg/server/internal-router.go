package server

import (
	"net/http"

	"github.com/dimiro1/health"
	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/config"
	"github.com/sirupsen/logrus"
)

// GenerateInternalRouter Generate internal router
func GenerateInternalRouter(logger *logrus.Logger, cfg *config.Config) http.Handler {
	// Set release mod
	gin.SetMode(gin.ReleaseMode)
	// Create router
	router := gin.New()
	// Add middlewares
	router.Use(gin.Recovery())
	router.Use(logMiddleware(logger))
	// Add routes
	healthHandler := health.NewHandler()
	router.GET("/health", gin.WrapH(healthHandler))
	return router
}
