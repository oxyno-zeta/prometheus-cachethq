package cachethq

import (
	"github.com/andygrunwald/cachet"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/config"
)

// NewContext will generate a new context
func NewContext(cfg *config.Config) (*Context, error) {
	client, err := cachet.NewClient(cfg.Cachet.URL, nil)
	if err != nil {
		return nil, err
	}
	client.Authentication.SetTokenAuth(cfg.Cachet.APIKey)
	return &Context{cachetCfg: cfg.Cachet, client: client}, nil
}
