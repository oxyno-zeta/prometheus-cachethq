package prometheushook

import (
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/business/prometheushook/models"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/cachethq"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/config"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/metrics"
)

type Client interface {
	ManageHook(promAlertHook *models.PrometheusAlertHook) error
}

func New(cfgManager config.Manager, metricsCtx metrics.Client, cachethqCtx cachethq.Client) Client {
	return &service{cachethqCtx: cachethqCtx, metricsCtx: metricsCtx, cfgManager: cfgManager}
}
