package cachethq

import (
	"github.com/andygrunwald/cachet"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/config"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/errors"
)

func (ctx *instance) findComponentGroup(name string) (*cachet.ComponentGroup, error) {
	// Create query params for name filter
	queryParams := &cachet.ComponentGroupsQueryParams{
		Name: name,
	}
	c, _, err := ctx.client.ComponentGroups.GetAll(queryParams)
	if err != nil {
		return nil, err
	}

	// Check if length is 0
	if len(c.ComponentGroups) == 0 {
		return nil, errors.NewNotFoundError(ErrComponentGroupNotFound)
	}

	// Get component group
	grp := c.ComponentGroups[0]
	return &grp, nil
}

func (ctx *instance) findComponent(name string, groupName string) (*cachet.Component, error) {
	// Creating default group id
	grpID := 0
	// Find group if possible
	if groupName != "" {
		grp, err := ctx.findComponentGroup(groupName)
		// Check error
		if err != nil {
			return nil, err
		}

		grpID = grp.ID
	}

	// Create query params for name filter
	queryParams := &cachet.ComponentsQueryParams{
		Name:         name,
		GroupID:      grpID,
		QueryOptions: cachet.QueryOptions{PerPage: 10000},
	}
	c, _, err := ctx.client.Components.GetAll(queryParams)
	if err != nil {
		return nil, errors.NewInternalServerError(err)
	}

	// Check if length is 0
	if len(c.Components) == 0 {
		return nil, errors.NewNotFoundError(ErrComponentNotFound)
	}

	// Filter components by groups
	// Client doesn't manage group id equal to 0 for no groups...
	for _, comp := range c.Components {
		if comp.GroupID == grpID {
			return &comp, nil
		}
	}

	// Default case
	return nil, errors.NewNotFoundError(ErrComponentNotFound)
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
