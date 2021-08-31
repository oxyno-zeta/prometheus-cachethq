package prometheushook

import (
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/business/prometheushook/models"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/cachethq"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/config"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/metrics"
	"github.com/thoas/go-funk"
)

const alertNameKey = "alertname"

type service struct {
	cfgManager  config.Manager
	cachethqCtx cachethq.Client
	metricsCtx  metrics.Client
}

func (ctx *service) ManageHook(promAlertHook *models.PrometheusAlertHook) error {
	// Get configuration
	cfg := ctx.cfgManager.GetConfig()

	// Loop over targets
	for _, target := range cfg.Targets {
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
			matchingLabelsKeys, _ := funk.Keys(matchingLabels).([]string)

			// Check if there is an alert target that match this alert
			for _, alert := range promAlertHook.Alerts {
				// Check if alert if matching current target
				if isAlertMatching(matchingLabelsKeys, matchingLabels, alert) {
					// Alert is matching
					componentStatus := target.Component.Status
					// Check if alert have status resolved
					if alert.Status == models.PrometheusStatusResolved {
						componentStatus = config.ComponentOperationalStatus
					}

					// Check if target have an incident configuration
					if target.Incident != nil {
						// Store incident status
						incidentStatus := target.Incident.Status
						// Check if alert have status resolved
						if alert.Status == models.PrometheusStatusResolved {
							// Store incident status to fixed
							incidentStatus = config.IncidentFixedStatus
						}

						// Create incident
						err := ctx.cachethqCtx.CreateIncident(
							target.Component.Name,
							target.Component.GroupName,
							componentStatus,
							target.Incident,
							incidentStatus,
						)
						// Check error
						if err != nil {
							return err
						}

						// Increment incident counter
						ctx.metricsCtx.IncrementIncidentManagedCounter(incidentStatus, componentStatus)
					} else {
						// Change component status
						err := ctx.cachethqCtx.ChangeComponentStatus(target.Component.Name, target.Component.GroupName, componentStatus)
						// Check error
						if err != nil {
							return err
						}

						// Increment component counter
						ctx.metricsCtx.IncrementComponentManagedCounter(componentStatus)
					}
				}
			}
		}
	}

	return nil
}

func isAlertMatching(matchingLabelsKeys []string, matchingLabels map[string]string, alert *models.PrometheusAlertDetail) bool {
	// Check if args exists
	if matchingLabelsKeys == nil || matchingLabels == nil || alert == nil {
		return false
	}

	// Check matching labels keys, if length is 0 => stop because cannot be possible
	if len(matchingLabelsKeys) == 0 {
		return false
	}

	// Loop over keys
	for _, key := range matchingLabelsKeys {
		// Get key
		value, ok := alert.Labels[key]
		if !ok {
			// Key doesn't exists => stop
			return false
		}
		// Check if values are identical
		if value != matchingLabels[key] {
			// Values aren't identical => stop
			return false
		}
	}

	// Default result
	return true
}
