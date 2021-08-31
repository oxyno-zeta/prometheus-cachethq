package business

import (
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/business/prometheushook"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/cachethq"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/config"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/log"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/metrics"
)

type Services struct {
	systemLogger      log.Logger
	PrometheusHookSvc prometheushook.Client
}

func NewServices(systemLogger log.Logger, cfgManager config.Manager, metricsCl metrics.Client, cachethqCtx cachethq.Client) *Services {
	// Create prometheus hook service
	promSvc := prometheushook.New(cfgManager, metricsCl, cachethqCtx)

	return &Services{
		systemLogger:      systemLogger,
		PrometheusHookSvc: promSvc,
	}
}
