package server

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/business"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/business/prometheushook/models"
	cerrors "github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/common/errors"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/common/utils"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/config"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/log"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/metrics"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/server/middlewares"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/signalhandler"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/tracing"
)

type Server struct {
	logger           log.Logger
	cfgManager       config.Manager
	metricsCl        metrics.Client
	tracingSvc       tracing.Service
	busiServices     *business.Services
	signalHandlerSvc signalhandler.Client
	server           *http.Server
}

func NewServer(
	logger log.Logger, cfgManager config.Manager, metricsCl metrics.Client,
	tracingSvc tracing.Service, busiServices *business.Services,
	signalHandlerSvc signalhandler.Client,
) *Server {
	return &Server{
		logger:           logger,
		cfgManager:       cfgManager,
		metricsCl:        metricsCl,
		tracingSvc:       tracingSvc,
		busiServices:     busiServices,
		signalHandlerSvc: signalHandlerSvc,
	}
}

func (svr *Server) GenerateServer() error {
	// Get configuration
	cfg := svr.cfgManager.GetConfig()
	// Generate router
	r, err := svr.generateRouter()
	if err != nil {
		return err
	}

	// Create server
	addr := cfg.Server.ListenAddr + ":" + strconv.Itoa(cfg.Server.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	// Prepare for configuration onChange
	svr.cfgManager.AddOnChangeHook(func() {
		// Generate router
		r, err2 := svr.generateRouter()
		if err2 != nil {
			svr.logger.Fatal(err2)
		}
		// Change server handler
		server.Handler = r
		svr.logger.Info("Server handler reloaded")
	})

	// Store server
	svr.server = server

	return nil
}

func (svr *Server) generateRouter() (http.Handler, error) {
	// Get configuration
	cfg := svr.cfgManager.GetConfig()
	// Set release mod
	gin.SetMode(gin.ReleaseMode)
	// Create router
	router := gin.New()
	// Manage no route
	router.NoRoute(func(c *gin.Context) {
		// Create not found error
		err := cerrors.NewNotFoundError("404 not found")

		// Answer
		utils.AnswerWithError(c, err)
	})
	// Add middlewares
	router.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithDecompressFn(gzip.DefaultDecompressHandle)))
	router.Use(gin.Recovery())
	router.Use(svr.signalHandlerSvc.ActiveRequestCounterMiddleware())
	router.Use(middlewares.RequestID(svr.logger))
	router.Use(svr.tracingSvc.HTTPMiddleware(middlewares.GetRequestIDFromContext))
	router.Use(log.Middleware(svr.logger, middlewares.GetRequestIDFromGin, tracing.GetSpanIDFromContext))
	router.Use(svr.metricsCl.Instrument("business"))
	// Add helmet for security
	router.Use(helmet.Default())
	// Add cors if configured
	err := manageCORS(router, cfg.Server)
	// Check error
	if err != nil {
		return nil, err
	}

	// Add routes
	router.POST("/prometheus/webhook", func(c *gin.Context) {
		// Get logger from request
		reqLogger := log.GetLoggerFromGin(c)

		var alerts models.PrometheusAlertHook
		// Try to map data
		err := c.ShouldBindJSON(&alerts)
		// Check if error exists
		if err != nil {
			reqLogger.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Validate input
		err = alerts.Validate()
		if err != nil {
			reqLogger.Error(err)

			utils.AnswerWithError(c, err)
			return
		}
		err = svr.busiServices.PrometheusHookSvc.ManageHook(&alerts)
		if err != nil {
			reqLogger.Error(err)

			utils.AnswerWithError(c, err)
			return
		}
		c.Status(http.StatusNoContent)
	})

	return router, nil
}

func (svr *Server) Listen() error {
	svr.logger.Infof("Server listening on %s", svr.server.Addr)
	err := svr.server.ListenAndServe()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Default
	return nil
}
