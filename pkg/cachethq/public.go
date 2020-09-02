package cachethq

import (
	"github.com/andygrunwald/cachet"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/config"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/errors"
)

// ChangeComponentStatus Change component status
func (ctx *instance) ChangeComponentStatus(name string, groupName string, stringStatus string) error {
	// Find component
	compo, err := ctx.findComponent(name, groupName)
	if err != nil {
		return err
	}
	status, err := getCachetHQComponentStatus(stringStatus)
	if err != nil {
		return errors.NewInternalServerError(err)
	}
	// Store component ID
	id := compo.ID
	// Change component status
	compo.Status = status
	// Run update request
	_, _, err = ctx.client.Components.Update(id, compo)
	if err != nil {
		return errors.NewInternalServerError(err)
	}
	return nil
}

// CreateIncident Create incident
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
		return errors.NewInternalServerError(err)
	}
	incidentStatus, err := getCachetHQIncidentStatus(incidentStringStatus)
	if err != nil {
		return errors.NewInternalServerError(err)
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
		return errors.NewInternalServerError(err)
	}
	return nil
}
