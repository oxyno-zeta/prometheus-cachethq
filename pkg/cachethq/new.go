package cachethq

import (
	"github.com/andygrunwald/cachet"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/config"
)

// NewInstance will generate a new instance
func NewInstance(cfg *config.Config) (Instance, error) {
	client, err := cachet.NewClient(cfg.Cachet.URL, nil)
	if err != nil {
		return nil, err
	}
	client.Authentication.SetTokenAuth(cfg.Cachet.APIKey)
	return &instance{cachetCfg: cfg.Cachet, client: client}, nil
}
