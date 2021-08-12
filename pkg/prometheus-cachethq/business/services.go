package business

import (
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/business/prometheushook"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/config"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/log"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/metrics"
)

type Services struct {
	systemLogger      log.Logger
	PrometheusHookSvc prometheushook.Client
}

func NewServices(systemLogger log.Logger, cfgManager config.Manager, metricsCl metrics.Client) *Services {
	// Get configuration
	cfg := cfgManager.GetConfig()

	// Create prometheus hook service
	promSvc, _ := prometheushook.New(cfg, metricsCl)

	return &Services{
		systemLogger:      systemLogger,
		PrometheusHookSvc: promSvc,
	}
}
