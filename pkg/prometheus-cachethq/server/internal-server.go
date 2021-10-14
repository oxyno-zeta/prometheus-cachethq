package server

import (
	"net/http"
	"strconv"
	"time"

	gosundheit "github.com/AppsFlyer/go-sundheit"
	healthhttp "github.com/AppsFlyer/go-sundheit/http"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/config"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/log"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/metrics"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/server/middlewares"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/signalhandler"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/tracing"
	"github.com/pkg/errors"
)

const (
	DefaultHealthCheckTimeout = time.Second
)

type InternalServer struct {
	logger           log.Logger
	cfgManager       config.Manager
	metricsCl        metrics.Client
	checkers         []*CheckerInput
	server           *http.Server
	signalHandlerSvc signalhandler.Client
}

type CheckerInput struct {
	Name         string
	CheckFn      func() error
	Interval     time.Duration
	Timeout      time.Duration
	InitialDelay time.Duration
}

func NewInternalServer(logger log.Logger, cfgManager config.Manager, metricsCl metrics.Client, signalHandlerSvc signalhandler.Client) *InternalServer {
	return &InternalServer{
		logger:           logger,
		cfgManager:       cfgManager,
		metricsCl:        metricsCl,
		signalHandlerSvc: signalHandlerSvc,
		checkers:         make([]*CheckerInput, 0),
	}
}

// AddChecker allow to add a health checker.
func (svr *InternalServer) AddChecker(chI *CheckerInput) {
	// Append
	svr.checkers = append(svr.checkers, chI)
}

func (svr *InternalServer) generateInternalRouter() (http.Handler, error) {
	// Get configuration
	cfg := svr.cfgManager.GetConfig()

	// Set release mod
	gin.SetMode(gin.ReleaseMode)
	// Create router
	router := gin.New()
	// Add middlewares
	router.Use(gin.Recovery())
	router.Use(helmet.Default())
	router.Use(middlewares.RequestID(svr.logger))
	router.Use(log.Middleware(svr.logger, middlewares.GetRequestIDFromGin, tracing.GetSpanIDFromContext))
	router.Use(svr.metricsCl.Instrument("internal"))
	// Add cors if configured
	err := manageCORS(router, cfg.InternalServer)
	// Check error
	if err != nil {
		return nil, err
	}

	// create a new health instance
	h2 := gosundheit.New()

	for _, it := range svr.checkers {
		// Create logger
		logger := svr.logger.WithField("health-check-target", it.Name)

		// Initialize check options
		options := make([]gosundheit.CheckOption, 0)

		// Check if timeout is set, otherwise put a default value
		if it.Timeout == 0 {
			options = append(options, gosundheit.ExecutionTimeout(DefaultHealthCheckTimeout))
		} else {
			options = append(options, gosundheit.ExecutionTimeout(it.Timeout))
		}

		// Check if initial delay is set, otherwise ignore it
		if it.InitialDelay != 0 {
			options = append(options, gosundheit.InitialDelay(it.InitialDelay))
		}

		// Set interval
		options = append(options, gosundheit.ExecutionPeriod(it.Interval))

		// Register check
		err = h2.RegisterCheck(
			&customHealthChecker{
				logger: logger,
				name:   it.Name,
				fn:     it.CheckFn,
			},
			options...,
		)
		// Check error
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	// Add metrics path
	router.GET("/metrics", gin.WrapH(svr.metricsCl.PrometheusHTTPHandler()))
	router.GET("/health", gin.WrapH(healthhttp.HandleHealthJSON(h2)))
	router.GET("/ready", func(c *gin.Context) {
		// Check if system is in shutdown
		if svr.signalHandlerSvc.IsStoppingSystem() {
			// Response with service unavailable flag
			c.JSON(http.StatusServiceUnavailable, gin.H{"reason": "system stopping"})

			return
		}

		// Otherwise, send health check result
		gin.WrapH(healthhttp.HandleHealthJSON(h2))(c)
	})

	return router, nil
}

func (svr *InternalServer) Listen() error {
	svr.logger.Infof("Internal server listening on %s", svr.server.Addr)
	err := svr.server.ListenAndServe()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	// Default
	return nil
}

func (svr *InternalServer) GenerateServer() error {
	// Get configuration
	cfg := svr.cfgManager.GetConfig()
	// Generate internal router
	r, err := svr.generateInternalRouter()
	// Check error
	if err != nil {
		return err
	}
	// Create server
	addr := cfg.InternalServer.ListenAddr + ":" + strconv.Itoa(cfg.InternalServer.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	// Store server
	svr.server = server

	return nil
}
