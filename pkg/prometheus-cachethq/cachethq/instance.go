package cachethq

import (
	"github.com/andygrunwald/cachet"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/common/errors"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/config"
)

type instance struct {
	client    *cachet.Client
	cachetCfg *config.CachetConfig
}

func (ctx *instance) ChangeComponentStatus(name string, groupName string, stringStatus string) error {
	// Find component
	compo, err := ctx.findComponent(name, groupName)
	if err != nil {
		return err
	}
	status, err := getCachetHQComponentStatus(stringStatus)
	if err != nil {
		return errors.NewInternalServerErrorWithError(err)
	}
	// Store component ID
	id := compo.ID
	// Change component status
	compo.Status = status
	// Run update request
	_, _, err = ctx.client.Components.Update(id, compo)
	if err != nil {
		return errors.NewInternalServerErrorWithError(err)
	}
	return nil
}

func (ctx *instance) CreateIncident(
	componentName string, componentGroupName string, componentStringStatus string,
	incident *config.TargetIncident, incidentStringStatus string) error {
	// Find component
	compo, err := ctx.findComponent(componentName, componentGroupName)
	if err != nil {
		return err
	}
	componentStatus, err := getCachetHQComponentStatus(componentStringStatus)
	if err != nil {
		return errors.NewInternalServerErrorWithError(err)
	}
	incidentStatus, err := getCachetHQIncidentStatus(incidentStringStatus)
	if err != nil {
		return errors.NewInternalServerErrorWithError(err)
	}
	visibility := getCachetHQIncidentVisibility(incident.Public)
	inci := &cachet.Incident{
		Name:            incident.Name,
		Message:         incident.Content,
		Visible:         visibility,
		Status:          incidentStatus,
		ComponentID:     compo.ID,
		ComponentStatus: componentStatus,
	}
	_, _, err = ctx.client.Incidents.Create(inci)
	if err != nil {
		return errors.NewInternalServerErrorWithError(err)
	}
	return nil
}

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
		return nil, errors.NewNotFoundErrorWithError(ErrComponentGroupNotFound)
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
		return nil, errors.NewInternalServerErrorWithError(err)
	}

	// Check if length is 0
	if len(c.Components) == 0 {
		return nil, errors.NewNotFoundErrorWithError(ErrComponentNotFound)
	}

	// Filter components by groups
	// Client doesn't manage group id equal to 0 for no groups...
	for _, comp := range c.Components {
		if comp.GroupID == grpID {
			return &comp, nil
		}
	}

	// Default case
	return nil, errors.NewNotFoundErrorWithError(ErrComponentNotFound)
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
