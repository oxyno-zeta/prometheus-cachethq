package prometheushook

import "github.com/oxyno-zeta/prometheus-cachethq/pkg/errors"

func (pah *PrometheusAlertHook) Validate() error {
	// Check hook version
	if pah.Version != "4" {
		return errors.NewBadInputError(ErrHookVersionNotSupported)
	}
	return nil
}
