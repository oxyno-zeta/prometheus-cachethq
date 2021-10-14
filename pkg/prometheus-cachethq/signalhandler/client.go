package signalhandler

import (
	"os"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/log"
	"github.com/thoas/go-funk"
)

//go:generate mockgen -destination=./mocks/mock_Client.go -package=mocks github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/signalhandler Client
type Client interface {
	// Initialize will initialize service.
	// Important note: this must be called only once.
	Initialize() error
	// OnSignal will add a hook on specific signal.
	// Hooks for SIGINT and SIGTERM are called here before checking that all requests are finished and calling onExit hooks.
	OnSignal(signal os.Signal, hook func())
	// OnExit will add a hook that will be called when a SIGINT or SIGTERM is caught and when the application will be closed
	// That will be launched only when all incoming requests are finished.
	OnExit(hook func())
	// Middleware to count active requests.
	ActiveRequestCounterMiddleware() gin.HandlerFunc
	// Is stopping system will return true if the application is stopping.
	IsStoppingSystem() bool
}

func NewClient(logger log.Logger, serverMode bool, signalListToNotify []os.Signal) Client {
	// Create signal list to notify
	signalListToNotifyInternal := []os.Signal{syscall.SIGTERM, syscall.SIGINT}
	// Append all items from input inside
	signalListToNotifyInternal = append(signalListToNotifyInternal, signalListToNotify...)
	// Filter to unique
	signalListToNotifyInternal, _ = funk.Uniq(signalListToNotifyInternal).([]os.Signal)

	return &service{
		logger:                   logger,
		serverMode:               serverMode,
		signalListToNotify:       signalListToNotifyInternal,
		hooksStorage:             map[os.Signal][]func(){},
		onExitHookStorage:        []func(){},
		activeRequestCounter:     0,
		activeRequestCounterChan: make(chan int64),
	}
}
