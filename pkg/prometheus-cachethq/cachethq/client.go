package cachethq

import (
	"github.com/pkg/errors"

	"github.com/andygrunwald/cachet"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/config"
)

// ErrComponentNotFound Error component not found.
var ErrComponentNotFound = errors.New("component not found")

// ErrComponentGroupNotFound Error component group not found.
var ErrComponentGroupNotFound = errors.New("component group not found")

// ErrStatusNotFound Error status not found.
var ErrStatusNotFound = errors.New("status not found")

// Instance instance interface.
type Client interface {
	// ChangeComponentStatus will change the component status in CachetHQ.
	ChangeComponentStatus(name string, groupName string, stringStatus string) error
	// CreateIncident will create an incident in CachetHQ.
	CreateIncident(
		componentName string,
		componentGroupName string,
		componentStringStatus string,
		incident *config.TargetIncident,
		incidentStringStatus string,
	) error
}

// NewInstance will generate a new instance.
func NewInstance(cfg *config.Config) (Client, error) {
	client, err := cachet.NewClient(cfg.Cachet.URL, nil)
	if err != nil {
		return nil, err
	}
	client.Authentication.SetTokenAuth(cfg.Cachet.APIKey)
	return &instance{cachetCfg: cfg.Cachet, client: client}, nil
}
