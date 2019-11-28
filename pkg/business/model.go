package business

import (
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/cachethq"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/config"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/metrics"
)

// Context business context
type Context struct {
	cfg         *config.Config
	cachethqCtx cachethq.Instance
	metricsCtx  metrics.Instance
}
