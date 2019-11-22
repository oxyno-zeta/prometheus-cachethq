package cachethq

import (
	"errors"
	"github.com/andygrunwald/cachet"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/config"
)

// Context CachetHQ Configuration
type Context struct {
	client    *cachet.Client
	cachetCfg *config.CachetConfig
}

// ErrComponentNotFound Error component not found
var ErrComponentNotFound = errors.New("component not found")

// ErrStatusNotFound Error status not found
var ErrStatusNotFound = errors.New("status not found")
