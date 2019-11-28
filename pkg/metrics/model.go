package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// Instance Intance metrics interface
type Instance interface {
	Instrument() gin.HandlerFunc
	IncrementIncidentManagedCounter(string, string)
	IncrementComponentManagedCounter(string)
}

type instance struct {
	reqCnt                *prometheus.CounterVec
	resSz                 prometheus.Summary
	reqDur                prometheus.Summary
	reqSz                 prometheus.Summary
	up                    prometheus.Gauge
	incidentManagedTotal  *prometheus.CounterVec
	componentManagedTotal *prometheus.CounterVec
}

// NewInstance will generate a new Instance
func NewInstance() Instance {
	ctx := &instance{}
	ctx.register()
	return ctx
}
