package business

import (
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/cachethq"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/config"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/metrics"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheushook"
)

type testCachetInstance struct {
	err error
}

func (tci *testCachetInstance) ChangeComponentStatus(name string, groupName string, stringStatus string) error {
	return tci.err
}

func (tci *testCachetInstance) CreateIncident(
	componentName string, componentGroupName string, componentStringStatus string,
	incident *config.TargetIncident, incidentStringStatus string) error {
	return tci.err
}

type testMetricsInstance struct{}

func (tmi *testMetricsInstance) Instrument() gin.HandlerFunc {
	return nil
}
func (tmi *testMetricsInstance) IncrementIncidentManagedCounter(string1 string, string2 string) {}
func (tmi *testMetricsInstance) IncrementComponentManagedCounter(string)                        {}

func TestContext_ManageHook(t *testing.T) {
	type fields struct {
		cfg         *config.Config
		cachethqCtx cachethq.Instance
		metricsCtx  metrics.Instance
	}
	type args struct {
		promAlertHook *prometheushook.PrometheusAlertHook
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "No input",
			fields: fields{
				cfg:         &config.Config{},
				cachethqCtx: nil,
				metricsCtx:  nil,
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "No targets list",
			fields: fields{
				cfg: &config.Config{
					Targets: []*config.Target{},
				},
				cachethqCtx: nil,
				metricsCtx:  nil,
			},
			args:    args{},
			wantErr: false,
		},
		{
			name: "Non matching alert",
			fields: fields{
				cfg: &config.Config{
					Targets: []*config.Target{{
						Component: &config.TargetComponent{
							Name:   "Component1",
							Status: "OPERATIONAL",
						},
						Alerts: []*config.TargetAlerts{{
							Name: "Alert1",
						}},
					}},
				},
				cachethqCtx: nil,
				metricsCtx:  nil,
			},
			args: args{
				promAlertHook: &prometheushook.PrometheusAlertHook{
					Version: "4",
					Alerts: []*prometheushook.PrometheusAlertDetail{{
						Status: "firing",
						Labels: map[string]string{
							"label1": "value1",
						},
					}},
				},
			},
			wantErr: false,
		},
		{
			name: "Matching firing alert for component",
			fields: fields{
				cfg: &config.Config{
					Targets: []*config.Target{{
						Component: &config.TargetComponent{
							Name:   "Component1",
							Status: "OPERATIONAL",
						},
						Alerts: []*config.TargetAlerts{{
							Name: "Alert1",
						}},
					}},
				},
				cachethqCtx: &testCachetInstance{err: nil},
				metricsCtx:  &testMetricsInstance{},
			},
			args: args{
				promAlertHook: &prometheushook.PrometheusAlertHook{
					Version: "4",
					Alerts: []*prometheushook.PrometheusAlertDetail{{
						Status: "firing",
						Labels: map[string]string{
							"alertname": "Alert1",
							"label1":    "value1",
						},
					}},
				},
			},
			wantErr: false,
		},
		{
			name: "Matching firing alert for component with error",
			fields: fields{
				cfg: &config.Config{
					Targets: []*config.Target{{
						Component: &config.TargetComponent{
							Name:   "Component1",
							Status: "OPERATIONAL",
						},
						Alerts: []*config.TargetAlerts{{
							Name: "Alert1",
						}},
					}},
				},
				cachethqCtx: &testCachetInstance{err: errors.New("Error1")},
				metricsCtx:  &testMetricsInstance{},
			},
			args: args{
				promAlertHook: &prometheushook.PrometheusAlertHook{
					Version: "4",
					Alerts: []*prometheushook.PrometheusAlertDetail{{
						Status: "firing",
						Labels: map[string]string{
							"alertname": "Alert1",
							"label1":    "value1",
						},
					}},
				},
			},
			wantErr: true,
		},
		{
			name: "Matching firing alert for component and incident",
			fields: fields{
				cfg: &config.Config{
					Targets: []*config.Target{{
						Component: &config.TargetComponent{
							Name:   "Component1",
							Status: "OPERATIONAL",
						},
						Incident: &config.TargetIncident{
							Name:    "Incident1",
							Content: "Content1",
							Status:  "INVESTIGATING",
							Public:  false,
						},
						Alerts: []*config.TargetAlerts{{
							Name: "Alert1",
						}},
					}},
				},
				cachethqCtx: &testCachetInstance{err: nil},
				metricsCtx:  &testMetricsInstance{},
			},
			args: args{
				promAlertHook: &prometheushook.PrometheusAlertHook{
					Version: "4",
					Alerts: []*prometheushook.PrometheusAlertDetail{{
						Status: "firing",
						Labels: map[string]string{
							"alertname": "Alert1",
							"label1":    "value1",
						},
					}},
				},
			},
			wantErr: false,
		},
		{
			name: "Matching firing alert for component and incident with error",
			fields: fields{
				cfg: &config.Config{
					Targets: []*config.Target{{
						Component: &config.TargetComponent{
							Name:   "Component1",
							Status: "OPERATIONAL",
						},
						Incident: &config.TargetIncident{
							Name:    "Incident1",
							Content: "Content1",
							Status:  "INVESTIGATING",
							Public:  false,
						},
						Alerts: []*config.TargetAlerts{{
							Name: "Alert1",
						}},
					}},
				},
				cachethqCtx: &testCachetInstance{err: errors.New("Error1")},
				metricsCtx:  &testMetricsInstance{},
			},
			args: args{
				promAlertHook: &prometheushook.PrometheusAlertHook{
					Version: "4",
					Alerts: []*prometheushook.PrometheusAlertDetail{{
						Status: "firing",
						Labels: map[string]string{
							"alertname": "Alert1",
							"label1":    "value1",
						},
					}},
				},
			},
			wantErr: true,
		},
		{
			name: "Matching resolved alert for component",
			fields: fields{
				cfg: &config.Config{
					Targets: []*config.Target{{
						Component: &config.TargetComponent{
							Name:   "Component1",
							Status: "OPERATIONAL",
						},
						Alerts: []*config.TargetAlerts{{
							Name: "Alert1",
						}},
					}},
				},
				cachethqCtx: &testCachetInstance{err: nil},
				metricsCtx:  &testMetricsInstance{},
			},
			args: args{
				promAlertHook: &prometheushook.PrometheusAlertHook{
					Version: "4",
					Alerts: []*prometheushook.PrometheusAlertDetail{{
						Status: "resolved",
						Labels: map[string]string{
							"alertname": "Alert1",
							"label1":    "value1",
						},
					}},
				},
			},
			wantErr: false,
		},
		{
			name: "Matching resolved alert for component and incident",
			fields: fields{
				cfg: &config.Config{
					Targets: []*config.Target{{
						Component: &config.TargetComponent{
							Name:   "Component1",
							Status: "OPERATIONAL",
						},
						Incident: &config.TargetIncident{
							Name:    "Incident1",
							Content: "Content1",
							Status:  "INVESTIGATING",
							Public:  false,
						},
						Alerts: []*config.TargetAlerts{{
							Name: "Alert1",
						}},
					}},
				},
				cachethqCtx: &testCachetInstance{err: nil},
				metricsCtx:  &testMetricsInstance{},
			},
			args: args{
				promAlertHook: &prometheushook.PrometheusAlertHook{
					Version: "4",
					Alerts: []*prometheushook.PrometheusAlertDetail{{
						Status: "resolved",
						Labels: map[string]string{
							"alertname": "Alert1",
							"label1":    "value1",
						},
					}},
				},
			},
			wantErr: false,
		},
		{
			name: "Matching firing alert for component with alert labels",
			fields: fields{
				cfg: &config.Config{
					Targets: []*config.Target{{
						Component: &config.TargetComponent{
							Name:   "Component1",
							Status: "OPERATIONAL",
						},
						Alerts: []*config.TargetAlerts{{
							Name: "Alert1",
							Labels: map[string]string{
								"label1": "value1",
							},
						}},
					}},
				},
				cachethqCtx: &testCachetInstance{err: nil},
				metricsCtx:  &testMetricsInstance{},
			},
			args: args{
				promAlertHook: &prometheushook.PrometheusAlertHook{
					Version: "4",
					Alerts: []*prometheushook.PrometheusAlertDetail{{
						Status: "firing",
						Labels: map[string]string{
							"alertname": "Alert1",
							"label1":    "value1",
						},
					}},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &Context{
				cfg:         tt.fields.cfg,
				cachethqCtx: tt.fields.cachethqCtx,
				metricsCtx:  tt.fields.metricsCtx,
			}
			if err := ctx.ManageHook(tt.args.promAlertHook); (err != nil) != tt.wantErr {
				t.Errorf("Context.ManageHook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
