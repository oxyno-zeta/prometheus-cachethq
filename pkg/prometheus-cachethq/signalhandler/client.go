package signalhandler

import (
	"os"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/log"
	"github.com/thoas/go-funk"
)

type Client interface {
	// Initialize will initialize service.
	// Important note: this must be called only once.
	Initialize() error
	// OnSignal will add a hook on specific signal.
	OnSignal(signal os.Signal, hook func())
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
		signalChan:               make(chan os.Signal),
		activeRequestCounter:     0,
		activeRequestCounterChan: make(chan int64),
	}
}
