package signalhandler

import (
	"github.com/gin-gonic/gin"
)

func (s *service) initializeServerMode() {
	// Start a go routine
	go func() {
		// Infinite loop over channel
		for x := range s.activeRequestCounterChan {
			// Add value
			s.activeRequestCounter += x
		}
	}()
}

func (s *service) ActiveRequestCounterMiddleware() gin.HandlerFunc {
	// Check if server mode isn't enabled
	if !s.serverMode {
		return func(c *gin.Context) { c.Next() }
	}

	// Middleware
	return func(c *gin.Context) {
		// Send +1 to active request counter channel
		s.activeRequestCounterChan <- 1

		// Send -1 to active request counter channel when request is finished
		defer func() { s.activeRequestCounterChan <- -1 }()

		// Next
		c.Next()
	}
}
