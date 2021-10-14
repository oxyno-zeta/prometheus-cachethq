package cachethq

import (
	"github.com/pkg/errors"

	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/config"
)

// ErrComponentNotFound Error component not found.
var ErrComponentNotFound = errors.New("component not found")

// ErrComponentGroupNotFound Error component group not found.
var ErrComponentGroupNotFound = errors.New("component group not found")

// ErrStatusNotFound Error status not found.
var ErrStatusNotFound = errors.New("status not found")

// Instance instance interface.
//go:generate mockgen -destination=./mocks/mock_Client.go -package=mocks github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/cachethq Client
type Client interface {
	// Initialize will initialize service.
	Initialize() error
	// Ping will ping CachetHQ to check availability.
	Ping() error
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
func NewInstance(cfgManager config.Manager) Client {
	return &instance{cfgManager: cfgManager}
}
