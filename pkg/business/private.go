package business

import "github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheushook"

func isAlertMatching(matchingLabelsKeys []string, matchingLabels map[string]string, alert *prometheushook.PrometheusAlertDetail) bool {
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
