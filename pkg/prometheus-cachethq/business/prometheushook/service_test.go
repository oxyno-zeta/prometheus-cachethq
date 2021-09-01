package prometheushook

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/business/prometheushook/models"
)

// func TestContext_ManageHook(t *testing.T) {
// 	type fields struct {
// 		cfg         *config.Config
// 		cachethqCtx cachethq.Client
// 		metricsCtx  metrics.Client
// 	}
// 	type args struct {
// 		promAlertHook *models.PrometheusAlertHook
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "No input",
// 			fields: fields{
// 				cfg:         &config.Config{},
// 				cachethqCtx: nil,
// 				metricsCtx:  nil,
// 			},
// 			args:    args{},
// 			wantErr: false,
// 		},
// 		{
// 			name: "No targets list",
// 			fields: fields{
// 				cfg: &config.Config{
// 					Targets: []*config.Target{},
// 				},
// 				cachethqCtx: nil,
// 				metricsCtx:  nil,
// 			},
// 			args:    args{},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Non matching alert",
// 			fields: fields{
// 				cfg: &config.Config{
// 					Targets: []*config.Target{{
// 						Component: &config.TargetComponent{
// 							Name:   "Component1",
// 							Status: "OPERATIONAL",
// 						},
// 						Alerts: []*config.TargetAlerts{{
// 							Name: "Alert1",
// 						}},
// 					}},
// 				},
// 				cachethqCtx: nil,
// 				metricsCtx:  nil,
// 			},
// 			args: args{
// 				promAlertHook: &models.PrometheusAlertHook{
// 					Version: "4",
// 					Alerts: []*prometheushook.PrometheusAlertDetail{{
// 						Status: "firing",
// 						Labels: map[string]string{
// 							"label1": "value1",
// 						},
// 					}},
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Matching firing alert for component",
// 			fields: fields{
// 				cfg: &config.Config{
// 					Targets: []*config.Target{{
// 						Component: &config.TargetComponent{
// 							Name:   "Component1",
// 							Status: "OPERATIONAL",
// 						},
// 						Alerts: []*config.TargetAlerts{{
// 							Name: "Alert1",
// 						}},
// 					}},
// 				},
// 				cachethqCtx: &testCachetInstance{err: nil},
// 				metricsCtx:  &testMetricsInstance{},
// 			},
// 			args: args{
// 				promAlertHook: &prometheushook.PrometheusAlertHook{
// 					Version: "4",
// 					Alerts: []*prometheushook.PrometheusAlertDetail{{
// 						Status: "firing",
// 						Labels: map[string]string{
// 							"alertname": "Alert1",
// 							"label1":    "value1",
// 						},
// 					}},
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Matching firing alert for component with error",
// 			fields: fields{
// 				cfg: &config.Config{
// 					Targets: []*config.Target{{
// 						Component: &config.TargetComponent{
// 							Name:   "Component1",
// 							Status: "OPERATIONAL",
// 						},
// 						Alerts: []*config.TargetAlerts{{
// 							Name: "Alert1",
// 						}},
// 					}},
// 				},
// 				cachethqCtx: &testCachetInstance{err: errors.New("Error1")},
// 				metricsCtx:  &testMetricsInstance{},
// 			},
// 			args: args{
// 				promAlertHook: &prometheushook.PrometheusAlertHook{
// 					Version: "4",
// 					Alerts: []*prometheushook.PrometheusAlertDetail{{
// 						Status: "firing",
// 						Labels: map[string]string{
// 							"alertname": "Alert1",
// 							"label1":    "value1",
// 						},
// 					}},
// 				},
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "Matching firing alert for component and incident",
// 			fields: fields{
// 				cfg: &config.Config{
// 					Targets: []*config.Target{{
// 						Component: &config.TargetComponent{
// 							Name:   "Component1",
// 							Status: "OPERATIONAL",
// 						},
// 						Incident: &config.TargetIncident{
// 							Name:    "Incident1",
// 							Content: "Content1",
// 							Status:  "INVESTIGATING",
// 							Public:  false,
// 						},
// 						Alerts: []*config.TargetAlerts{{
// 							Name: "Alert1",
// 						}},
// 					}},
// 				},
// 				cachethqCtx: &testCachetInstance{err: nil},
// 				metricsCtx:  &testMetricsInstance{},
// 			},
// 			args: args{
// 				promAlertHook: &prometheushook.PrometheusAlertHook{
// 					Version: "4",
// 					Alerts: []*prometheushook.PrometheusAlertDetail{{
// 						Status: "firing",
// 						Labels: map[string]string{
// 							"alertname": "Alert1",
// 							"label1":    "value1",
// 						},
// 					}},
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Matching firing alert for component and incident with error",
// 			fields: fields{
// 				cfg: &config.Config{
// 					Targets: []*config.Target{{
// 						Component: &config.TargetComponent{
// 							Name:   "Component1",
// 							Status: "OPERATIONAL",
// 						},
// 						Incident: &config.TargetIncident{
// 							Name:    "Incident1",
// 							Content: "Content1",
// 							Status:  "INVESTIGATING",
// 							Public:  false,
// 						},
// 						Alerts: []*config.TargetAlerts{{
// 							Name: "Alert1",
// 						}},
// 					}},
// 				},
// 				cachethqCtx: &testCachetInstance{err: errors.New("Error1")},
// 				metricsCtx:  &testMetricsInstance{},
// 			},
// 			args: args{
// 				promAlertHook: &prometheushook.PrometheusAlertHook{
// 					Version: "4",
// 					Alerts: []*prometheushook.PrometheusAlertDetail{{
// 						Status: "firing",
// 						Labels: map[string]string{
// 							"alertname": "Alert1",
// 							"label1":    "value1",
// 						},
// 					}},
// 				},
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "Matching resolved alert for component",
// 			fields: fields{
// 				cfg: &config.Config{
// 					Targets: []*config.Target{{
// 						Component: &config.TargetComponent{
// 							Name:   "Component1",
// 							Status: "OPERATIONAL",
// 						},
// 						Alerts: []*config.TargetAlerts{{
// 							Name: "Alert1",
// 						}},
// 					}},
// 				},
// 				cachethqCtx: &testCachetInstance{err: nil},
// 				metricsCtx:  &testMetricsInstance{},
// 			},
// 			args: args{
// 				promAlertHook: &prometheushook.PrometheusAlertHook{
// 					Version: "4",
// 					Alerts: []*prometheushook.PrometheusAlertDetail{{
// 						Status: "resolved",
// 						Labels: map[string]string{
// 							"alertname": "Alert1",
// 							"label1":    "value1",
// 						},
// 					}},
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Matching resolved alert for component and incident",
// 			fields: fields{
// 				cfg: &config.Config{
// 					Targets: []*config.Target{{
// 						Component: &config.TargetComponent{
// 							Name:   "Component1",
// 							Status: "OPERATIONAL",
// 						},
// 						Incident: &config.TargetIncident{
// 							Name:    "Incident1",
// 							Content: "Content1",
// 							Status:  "INVESTIGATING",
// 							Public:  false,
// 						},
// 						Alerts: []*config.TargetAlerts{{
// 							Name: "Alert1",
// 						}},
// 					}},
// 				},
// 				cachethqCtx: &testCachetInstance{err: nil},
// 				metricsCtx:  &testMetricsInstance{},
// 			},
// 			args: args{
// 				promAlertHook: &prometheushook.PrometheusAlertHook{
// 					Version: "4",
// 					Alerts: []*prometheushook.PrometheusAlertDetail{{
// 						Status: "resolved",
// 						Labels: map[string]string{
// 							"alertname": "Alert1",
// 							"label1":    "value1",
// 						},
// 					}},
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Matching firing alert for component with alert labels",
// 			fields: fields{
// 				cfg: &config.Config{
// 					Targets: []*config.Target{{
// 						Component: &config.TargetComponent{
// 							Name:   "Component1",
// 							Status: "OPERATIONAL",
// 						},
// 						Alerts: []*config.TargetAlerts{{
// 							Name: "Alert1",
// 							Labels: map[string]string{
// 								"label1": "value1",
// 							},
// 						}},
// 					}},
// 				},
// 				cachethqCtx: &testCachetInstance{err: nil},
// 				metricsCtx:  &testMetricsInstance{},
// 			},
// 			args: args{
// 				promAlertHook: &prometheushook.PrometheusAlertHook{
// 					Version: "4",
// 					Alerts: []*prometheushook.PrometheusAlertDetail{{
// 						Status: "firing",
// 						Labels: map[string]string{
// 							"alertname": "Alert1",
// 							"label1":    "value1",
// 						},
// 					}},
// 				},
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctx := &Context{
// 				cfg:         tt.fields.cfg,
// 				cachethqCtx: tt.fields.cachethqCtx,
// 				metricsCtx:  tt.fields.metricsCtx,
// 			}
// 			if err := ctx.ManageHook(tt.args.promAlertHook); (err != nil) != tt.wantErr {
// 				t.Errorf("Context.ManageHook() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

func Test_isAlertMatching(t *testing.T) {
	type args struct {
		matchingLabelsKeys []string
		matchingLabels     map[string]string
		alert              *models.PrometheusAlertDetail
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "No matching labels keys",
			args: args{
				matchingLabelsKeys: nil,
				matchingLabels:     nil,
				alert: &models.PrometheusAlertDetail{
					Status: "firing",
					Labels: map[string]string{
						"label1": "value1",
					},
				},
			},
			want: false,
		},
		{
			name: "Empty matching labels keys",
			args: args{
				matchingLabelsKeys: []string{},
				matchingLabels:     nil,
				alert: &models.PrometheusAlertDetail{
					Status: "firing",
					Labels: map[string]string{
						"label1": "value1",
					},
				},
			},
			want: false,
		},
		{
			name: "Labels don't match",
			args: args{
				matchingLabelsKeys: []string{"label_found"},
				matchingLabels: map[string]string{
					"label_found": "value",
				},
				alert: &models.PrometheusAlertDetail{
					Status: "firing",
					Labels: map[string]string{
						"label1": "value1",
					},
				},
			},
			want: false,
		},
		{
			name: "Labels don't match",
			args: args{
				matchingLabelsKeys: []string{"label_found", "label2"},
				matchingLabels: map[string]string{
					"label_found": "value",
					"label2":      "value2",
				},
				alert: &models.PrometheusAlertDetail{
					Status: "firing",
					Labels: map[string]string{
						"label1": "value1",
						"label2": "value2",
					},
				},
			},
			want: false,
		},
		{
			name: "Labels match and values don't match",
			args: args{
				matchingLabelsKeys: []string{"label1", "label2"},
				matchingLabels: map[string]string{
					"label1": "value",
					"label2": "value2",
				},
				alert: &models.PrometheusAlertDetail{
					Status: "firing",
					Labels: map[string]string{
						"label1": "value1",
						"label2": "value2",
					},
				},
			},
			want: false,
		},
		{
			name: "Labels and values match",
			args: args{
				matchingLabelsKeys: []string{"label1", "label2"},
				matchingLabels: map[string]string{
					"label1": "value1",
					"label2": "value2",
				},
				alert: &models.PrometheusAlertDetail{
					Status: "firing",
					Labels: map[string]string{
						"label1": "value1",
						"label2": "value2",
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlertMatching(tt.args.matchingLabelsKeys, tt.args.matchingLabels, tt.args.alert); got != tt.want {
				t.Errorf("isAlertMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_validateInputHook(t *testing.T) {
	type args struct {
		promAlertHook *models.PrometheusAlertHook
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "nil case",
			args:    args{},
			wantErr: true,
		},
		{
			name:    "empty case",
			args:    args{promAlertHook: &models.PrometheusAlertHook{}},
			wantErr: true,
		},
		{
			name: "not a valid version",
			args: args{
				promAlertHook: &models.PrometheusAlertHook{
					Version: "3",
					Alerts: []*models.PrometheusAlertDetail{
						{
							Status: "firing",
							Labels: map[string]string{
								"alertname": "fail",
							},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "valid",
			args: args{
				promAlertHook: &models.PrometheusAlertHook{
					Version: "4",
					Alerts: []*models.PrometheusAlertDetail{
						{
							Status: "firing",
							Labels: map[string]string{
								"alertname": "fail",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &service{
				validate: validator.New(),
			}
			if err := ctx.validateInputHook(tt.args.promAlertHook); (err != nil) != tt.wantErr {
				t.Errorf("service.validateInputHook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
