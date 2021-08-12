// +build integration

package server

import (
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/metrics"
)

// Generate metrics instance
var metricsCtx = metrics.NewMetricsClient()
