package business

import (
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/cachethq"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/config"
)

// NewContext Will generate a new business context
func NewContext(cfg *config.Config) (*Context, error) {
	cachethqCtx, err := cachethq.NewContext(cfg)
	if err != nil {
		return nil, err
	}
	return &Context{cachethqCtx: cachethqCtx, cfg: cfg}, nil
}
