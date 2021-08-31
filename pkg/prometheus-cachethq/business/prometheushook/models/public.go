package models

import "github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/common/errors"

// Validate validate prometheus alert hook.
func (pah *PrometheusAlertHook) Validate() error {
	// Check hook version
	if pah.Version != "4" {
		return errors.NewInvalidInputErrorWithError(ErrHookVersionNotSupported)
	}

	return nil
}
