package signalhandler

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/log"
)

const (
	TimeBetweenChecks = 200 * time.Millisecond
)

type service struct {
	logger                   log.Logger
	serverMode               bool
	signalListToNotify       []os.Signal
	signalChan               chan os.Signal
	hooksStorage             map[os.Signal][]func()
	activeRequestCounter     int64
	activeRequestCounterChan chan int64
	stoppingSysInProgress    bool
}

func (s *service) IsStoppingSystem() bool {
	return s.stoppingSysInProgress
}

func (s *service) Initialize() error {
	// Create signal channel
	signalChan := make(chan os.Signal, 1)

	// Notify watcher
	signal.Notify(
		signalChan,
		s.signalListToNotify...,
	)

	go func() {
		for sig := range signalChan {
			// Log
			s.logger.Infof("Catching signal \"%s\"", sig)

			// Get hook list for signal
			hooks := s.hooksStorage[sig]

			// Run all hooks
			for _, h := range hooks {
				// Start hook
				h()
			}

			// Check if signal is SIGTERM or SIGINT
			if sig == syscall.SIGINT || sig == syscall.SIGTERM {
				// Run stop hook
				s.stoppingAppHook()
			}
		}
	}()

	// Initialize server mode
	s.initializeServerMode()

	// Default
	return nil
}

func (s *service) OnSignal(signal os.Signal, hook func()) {
	// Check if array exist
	if s.hooksStorage[signal] != nil {
		// Create list
		s.hooksStorage[signal] = make([]func(), 0)
	}

	// Add item
	s.hooksStorage[signal] = append(s.hooksStorage[signal], hook)
}

func (s *service) stoppingAppHook() {
	// Check if application is already marked as in stopping mode
	if s.stoppingSysInProgress {
		// Avoid starting a new go routine for the same thing
		return
	}

	// Create ticker
	ticker := time.NewTicker(TimeBetweenChecks)

	// Starting the go routine
	go func() {
		// Loop
		for range ticker.C {
			// Log
			s.logger.Debug("Checking is application can be stopped")
			// Check if requests still in progress
			if s.activeRequestCounter == 0 {
				// Log
				s.logger.Info("Stopping application")
				// Stopping application
				os.Exit(0)
			}
		}
	}()

	// Updating the stopping flag
	s.stoppingSysInProgress = true
}
