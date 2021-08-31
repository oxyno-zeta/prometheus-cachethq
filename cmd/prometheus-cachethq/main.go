package main

import (
	"os"
	"syscall"
	"time"

	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/business"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/cachethq"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/config"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/log"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/metrics"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/server"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/signalhandler"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/tracing"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/version"
	"golang.org/x/sync/errgroup"
)

func main() {
	// Create new logger
	logger := log.NewLogger()

	// Create configuration manager
	cfgManager := config.NewManager(logger)

	// Load configuration
	err := cfgManager.Load()
	// Check error
	if err != nil {
		logger.Fatal(err)
	}

	// Get configuration
	cfg := cfgManager.GetConfig()
	// Configure logger
	err = logger.Configure(cfg.Log.Level, cfg.Log.Format, cfg.Log.FilePath)
	// Check error
	if err != nil {
		logger.Fatal(err)
	}

	// Watch change for logger (special case)
	cfgManager.AddOnChangeHook(func() {
		// Get configuration
		cfg := cfgManager.GetConfig()
		// Configure logger
		err = logger.Configure(cfg.Log.Level, cfg.Log.Format, cfg.Log.FilePath)
		// Check error
		if err != nil {
			logger.Fatal(err)
		}
	})

	// Getting version
	v := version.GetVersion()

	logger.Infof("Starting version: %s (git commit: %s) built on %s", v.Version, v.GitCommit, v.BuildDate)

	// Create metrics client
	metricsCl := metrics.NewMetricsClient()

	// Generate tracing service instance
	tracingSvc, err := tracing.New(cfgManager, logger)
	// Check error
	if err != nil {
		logger.Fatal(err)
	}
	// Prepare on reload hook
	cfgManager.AddOnChangeHook(func() {
		err = tracingSvc.Reload()
		// Check error
		if err != nil {
			logger.Fatal(err)
		}
	})

	// Create cachetHQ service
	cachetHQCl := cachethq.NewInstance(cfgManager)
	// Initialize
	err = cachetHQCl.Initialize()
	// Check error
	if err != nil {
		logger.Fatal(err)
	}
	// Prepare on reload hook
	cfgManager.AddOnChangeHook(func() {
		err = cachetHQCl.Initialize()
		// Check error
		if err != nil {
			logger.Fatal(err)
		}
	})

	// Create signal handler service
	signalHandlerSvc := signalhandler.NewClient(logger, true, []os.Signal{syscall.SIGTERM, syscall.SIGINT})
	// Initialize service
	err = signalHandlerSvc.Initialize()
	// Check error
	if err != nil {
		logger.Fatal(err)
	}

	// Create business services
	busServices := business.NewServices(logger, cfgManager, metricsCl, cachetHQCl)

	// Create servers
	svr := server.NewServer(logger, cfgManager, metricsCl, tracingSvc, busServices, signalHandlerSvc)
	intSvr := server.NewInternalServer(logger, cfgManager, metricsCl, signalHandlerSvc)

	intSvr.AddChecker(&server.CheckerInput{
		Name:     "cachetHQ",
		CheckFn:  cachetHQCl.Ping,
		Interval: time.Second,
		Timeout:  time.Second,
	})

	// Generate server
	err = svr.GenerateServer()
	if err != nil {
		logger.Fatal(err)
	}
	// Generate internal server
	err = intSvr.GenerateServer()
	if err != nil {
		logger.Fatal(err)
	}

	var g errgroup.Group

	g.Go(svr.Listen)
	g.Go(intSvr.Listen)

	if err := g.Wait(); err != nil {
		logger.Fatal(err)
	}
}
