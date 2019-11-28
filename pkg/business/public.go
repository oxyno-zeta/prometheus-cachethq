package business

import (
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/config"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheushook"
	"github.com/thoas/go-funk"
)

const alertNameKey = "alertname"

// ManageHook Manage Prometheus hook
func (ctx *Context) ManageHook(promAlertHook *prometheushook.PrometheusAlertHook) error {
	// Loop over targets
	for _, target := range ctx.cfg.Targets {
		// Loop over alerts matching
		for _, alertMatching := range target.Alerts {
			// Begin creation of matching labels
			var matchingLabels map[string]string
			if alertMatching.Labels != nil {
				matchingLabels = alertMatching.Labels
			} else {
				// Create empty map
				matchingLabels = make(map[string]string)
			}

			// Manage name case
			if alertMatching.Name != "" {
				// If name is defined, set it for search
				matchingLabels[alertNameKey] = alertMatching.Name
			}

			// Get matching labels keys
			matchingLabelsKeys := funk.Keys(matchingLabels).([]string)

			// Check if there is an alert target that match this alert
			for _, alert := range promAlertHook.Alerts {
				if isAlertMatching(matchingLabelsKeys, matchingLabels, alert) {
					// Alert is matching
					componentStatus := target.Component.Status
					if alert.Status == prometheushook.PrometheusStatusResolved {
						componentStatus = config.ComponentOperationalStatus
					}
					if target.Incident != nil {
						incidentStatus := target.Incident.Status
						if alert.Status == prometheushook.PrometheusStatusResolved {
							incidentStatus = config.IncidentFixedStatus
						}
						err := ctx.cachethqCtx.CreateIncident(target.Component.Name, componentStatus, target.Incident, incidentStatus)
						if err != nil {
							return err
						}
						ctx.metricsCtx.IncrementIncidentManagedCounter(incidentStatus, componentStatus)
					} else {
						err := ctx.cachethqCtx.ChangeComponentStatus(target.Component.Name, componentStatus)
						if err != nil {
							return err
						}
						ctx.metricsCtx.IncrementComponentManagedCounter(componentStatus)
					}
				}
			}
		}
	}
	return nil
}
