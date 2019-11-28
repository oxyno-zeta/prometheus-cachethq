package metrics

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Instrument will instrument gin routes
func (ctx *instance) Instrument() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		reqSz := computeApproximateRequestSize(c.Request)

		c.Next()

		status := strconv.Itoa(c.Writer.Status())
		elapsed := float64(time.Since(start)) / float64(time.Second)
		resSz := float64(c.Writer.Size())

		ctx.reqDur.Observe(elapsed)
		ctx.reqCnt.WithLabelValues(status, c.Request.Method, c.Request.Host, c.Request.URL.Path).Inc()
		ctx.reqSz.Observe(float64(reqSz))
		ctx.resSz.Observe(resSz)
	}
}

// IncrementIncidentManagedCounter Increment incident managed counter
func (ctx *instance) IncrementIncidentManagedCounter(incidentStatus string, componentStatus string) {
	ctx.incidentManagedTotal.WithLabelValues(incidentStatus).Inc()
	ctx.IncrementComponentManagedCounter(componentStatus)
}

// IncrementComponentManagedCounter Increment component managed counter
func (ctx *instance) IncrementComponentManagedCounter(status string) {
	ctx.componentManagedTotal.WithLabelValues(status).Inc()
}
