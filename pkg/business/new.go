package business

import (
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/cachethq"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/config"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/metrics"
)

// NewContext Will generate a new business context
func NewContext(cfg *config.Config, metricsCtx metrics.Instance) (*Context, error) {
	cachethqCtx, err := cachethq.NewInstance(cfg)
	if err != nil {
		return nil, err
	}
	return &Context{cachethqCtx: cachethqCtx, metricsCtx: metricsCtx, cfg: cfg}, nil
}
