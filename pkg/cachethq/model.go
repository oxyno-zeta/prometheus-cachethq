package cachethq

import (
	"errors"

	"github.com/andygrunwald/cachet"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/config"
)

// Instance instance interface
type Instance interface {
	ChangeComponentStatus(name string, groupName string, stringStatus string) error
	CreateIncident(
		componentName string,
		componentGroupName string,
		componentStringStatus string,
		incident *config.TargetIncident,
		incidentStringStatus string,
	) error
}

// instance CachetHQ instance
type instance struct {
	client    *cachet.Client
	cachetCfg *config.CachetConfig
}

// ErrComponentNotFound Error component not found
var ErrComponentNotFound = errors.New("component not found")

// ErrComponentGroupNotFound Error component group not found
var ErrComponentGroupNotFound = errors.New("component group not found")

// ErrStatusNotFound Error status not found
var ErrStatusNotFound = errors.New("status not found")
