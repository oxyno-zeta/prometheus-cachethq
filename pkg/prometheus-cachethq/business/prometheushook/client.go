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

func New(cfg *config.Config, metricsCtx metrics.Client) (Client, error) {
	cachethqCtx, err := cachethq.NewInstance(cfg)
	if err != nil {
		return nil, err
	}
	return &service{cachethqCtx: cachethqCtx, metricsCtx: metricsCtx, cfg: cfg}, nil
}
