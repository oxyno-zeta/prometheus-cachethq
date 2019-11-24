package cachethq

import (
	"github.com/andygrunwald/cachet"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/config"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/errors"
)

func (ctx *instance) findComponent(name string) (*cachet.Component, error) {
	// Create query params for name filter
	queryParams := &cachet.ComponentsQueryParams{
		Name: name,
	}
	c, _, err := ctx.client.Components.GetAll(queryParams)
	if err != nil {
		return nil, errors.NewInternalServerError(err)
	}

	// Check if length is 0
	if len(c.Components) == 0 {
		return nil, errors.NewNotFoundError(ErrComponentNotFound)
	}

	// Get component
	compo := c.Components[0]
	return &compo, nil
}

func getCachetHQComponentStatus(statusString string) (int, error) {
	switch statusString {
	case config.ComponentMajorOutageStatus:
		return cachet.ComponentStatusMajorOutage, nil
	case config.ComponentOperationalStatus:
		return cachet.ComponentStatusOperational, nil
	case config.ComponentPartialOutageStatus:
		return cachet.ComponentStatusPartialOutage, nil
	case config.ComponentPerformanceIssuesStatus:
		return cachet.ComponentStatusPerformanceIssues, nil
	default:
		return 0, ErrStatusNotFound
	}
}

func getCachetHQIncidentStatus(statusString string) (int, error) {
	switch statusString {
	case config.IncidentFixedStatus:
		return cachet.IncidentStatusFixed, nil
	case config.IncidentIdentifiedStatus:
		return cachet.IncidentStatusIdentified, nil
	case config.IncidentInvestigatingStatus:
		return cachet.IncidentStatusInvestigating, nil
	case config.IncidentWatchingStatus:
		return cachet.IncidentStatusWatching, nil
	default:
		return 0, ErrStatusNotFound
	}
}

func getCachetHQIncidentVisibility(visible bool) int {
	if visible {
		return cachet.IncidentVisibilityPublic
	}
	return cachet.IncidentVisibilityLoggedIn
}
