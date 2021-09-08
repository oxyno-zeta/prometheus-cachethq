// +build unit

package prometheushook

import (
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/business/prometheushook/models"
	cachethqmocks "github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/cachethq/mocks"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/config"
	configmocks "github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/config/mocks"
	metricsmocks "github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/metrics/mocks"
)

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

func Test_service_ManageHook(t *testing.T) {
	type cfgManagerGetConfigMockResult struct {
		res   *config.Config
		times int
	}
	type cachetCreateIncidentMockResult struct {
		input1 string
		input2 string
		input3 string
		input4 *config.TargetIncident
		input5 string
		err    error
		times  int
	}
	type cachetChangeComponentStatusMockResult struct {
		input1 string
		input2 string
		input3 string
		err    error
		times  int
	}
	type metricsIncrementIncidentManagedCounterMockResult struct {
		input1 string
		input2 string
		times  int
	}
	type metricsIncrementComponentManagedCounterMockResult struct {
		input1 string
		times  int
	}
	type args struct {
		promAlertHook *models.PrometheusAlertHook
	}
	tests := []struct {
		name                                        string
		args                                        args
		cfgManagerGetConfigMock                     cfgManagerGetConfigMockResult
		cachetCreateIncidentMock                    cachetCreateIncidentMockResult
		cachetChangeComponentStatusMock             cachetChangeComponentStatusMockResult
		metricsIncrementIncidentManagedCounterMock  metricsIncrementIncidentManagedCounterMockResult
		metricsIncrementComponentManagedCounterMock metricsIncrementComponentManagedCounterMockResult
		wantErr                                     bool
		errorString                                 string
	}{
		{
			name:    "no input",
			args:    args{},
			wantErr: true,
		},
		{
			name: "no targets list",
			cfgManagerGetConfigMock: cfgManagerGetConfigMockResult{
				res: &config.Config{
					Targets: []*config.Target{},
				},
				times: 1,
			},
			args: args{
				promAlertHook: &models.PrometheusAlertHook{
					Version: "4",
					Alerts: []*models.PrometheusAlertDetail{
						{Status: "firing", Labels: map[string]string{"alertname": "alert1"}},
					},
				},
			},
		},
		{
			name: "non matching alert",
			cfgManagerGetConfigMock: cfgManagerGetConfigMockResult{
				res: &config.Config{
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
				times: 1,
			},
			args: args{
				promAlertHook: &models.PrometheusAlertHook{
					Version: "4",
					Alerts: []*models.PrometheusAlertDetail{
						{
							Status: "firing",
							Labels: map[string]string{
								"alertname": "alert1",
							},
						},
					},
				},
			},
		},
		{
			name: "matching firing alert for component",
			cfgManagerGetConfigMock: cfgManagerGetConfigMockResult{
				res: &config.Config{
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
				times: 1,
			},
			args: args{
				promAlertHook: &models.PrometheusAlertHook{
					Version: "4",
					Alerts: []*models.PrometheusAlertDetail{{
						Status: "firing",
						Labels: map[string]string{
							"alertname": "Alert1",
							"label1":    "value1",
						},
					}},
				},
			},
			cachetChangeComponentStatusMock: cachetChangeComponentStatusMockResult{
				input1: "Component1",
				input2: "",
				input3: "OPERATIONAL",
				times:  1,
			},
			metricsIncrementComponentManagedCounterMock: metricsIncrementComponentManagedCounterMockResult{
				input1: "OPERATIONAL",
				times:  1,
			},
		},
		{
			name: "matching firing alert for component with error on cachethq call",
			cfgManagerGetConfigMock: cfgManagerGetConfigMockResult{
				res: &config.Config{
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
				times: 1,
			},
			args: args{
				promAlertHook: &models.PrometheusAlertHook{
					Version: "4",
					Alerts: []*models.PrometheusAlertDetail{{
						Status: "firing",
						Labels: map[string]string{
							"alertname": "Alert1",
							"label1":    "value1",
						},
					}},
				},
			},
			cachetChangeComponentStatusMock: cachetChangeComponentStatusMockResult{
				input1: "Component1",
				input2: "",
				input3: "OPERATIONAL",
				times:  1,
				err:    errors.New("fail"),
			},
			wantErr:     true,
			errorString: "fail",
		},
		{
			name: "matching firing alert for component and incident",
			cfgManagerGetConfigMock: cfgManagerGetConfigMockResult{
				res: &config.Config{
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
				times: 1,
			},
			args: args{
				promAlertHook: &models.PrometheusAlertHook{
					Version: "4",
					Alerts: []*models.PrometheusAlertDetail{{
						Status: "firing",
						Labels: map[string]string{
							"alertname": "Alert1",
							"label1":    "value1",
						},
					}},
				},
			},
			cachetCreateIncidentMock: cachetCreateIncidentMockResult{
				input1: "Component1",
				input2: "",
				input3: "OPERATIONAL",
				input4: &config.TargetIncident{
					Name:    "Incident1",
					Content: "Content1",
					Status:  "INVESTIGATING",
					Public:  false,
				},
				input5: "INVESTIGATING",
				times:  1,
			},
			metricsIncrementIncidentManagedCounterMock: metricsIncrementIncidentManagedCounterMockResult{
				input1: "INVESTIGATING",
				input2: "OPERATIONAL",
				times:  1,
			},
		},
		{
			name: "matching firing alert for component and incident with error on cachethq call",
			cfgManagerGetConfigMock: cfgManagerGetConfigMockResult{
				res: &config.Config{
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
				times: 1,
			},
			args: args{
				promAlertHook: &models.PrometheusAlertHook{
					Version: "4",
					Alerts: []*models.PrometheusAlertDetail{{
						Status: "firing",
						Labels: map[string]string{
							"alertname": "Alert1",
							"label1":    "value1",
						},
					}},
				},
			},
			cachetCreateIncidentMock: cachetCreateIncidentMockResult{
				input1: "Component1",
				input2: "",
				input3: "OPERATIONAL",
				input4: &config.TargetIncident{
					Name:    "Incident1",
					Content: "Content1",
					Status:  "INVESTIGATING",
					Public:  false,
				},
				input5: "INVESTIGATING",
				err:    errors.New("fail"),
				times:  1,
			},
			wantErr:     true,
			errorString: "fail",
		},
		{
			name: "matching resolved alert for component",
			cfgManagerGetConfigMock: cfgManagerGetConfigMockResult{
				res: &config.Config{
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
				times: 1,
			},
			args: args{
				promAlertHook: &models.PrometheusAlertHook{
					Version: "4",
					Alerts: []*models.PrometheusAlertDetail{{
						Status: "resolved",
						Labels: map[string]string{
							"alertname": "Alert1",
							"label1":    "value1",
						},
					}},
				},
			},
			cachetChangeComponentStatusMock: cachetChangeComponentStatusMockResult{
				input1: "Component1",
				input2: "",
				input3: "OPERATIONAL",
				times:  1,
			},
			metricsIncrementComponentManagedCounterMock: metricsIncrementComponentManagedCounterMockResult{
				input1: "OPERATIONAL",
				times:  1,
			},
		},
		{
			name: "matching resolved alert for component and incident",
			cfgManagerGetConfigMock: cfgManagerGetConfigMockResult{
				res: &config.Config{
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
				times: 1,
			},
			args: args{
				promAlertHook: &models.PrometheusAlertHook{
					Version: "4",
					Alerts: []*models.PrometheusAlertDetail{{
						Status: "resolved",
						Labels: map[string]string{
							"alertname": "Alert1",
							"label1":    "value1",
						},
					}},
				},
			},
			cachetCreateIncidentMock: cachetCreateIncidentMockResult{
				input1: "Component1",
				input2: "",
				input3: "OPERATIONAL",
				input4: &config.TargetIncident{
					Name:    "Incident1",
					Content: "Content1",
					Status:  "INVESTIGATING",
					Public:  false,
				},
				input5: "FIXED",
				times:  1,
			},
			metricsIncrementIncidentManagedCounterMock: metricsIncrementIncidentManagedCounterMockResult{
				input1: "FIXED",
				input2: "OPERATIONAL",
				times:  1,
			},
		},
		{
			name: "matching firing alert for component with alert labels",
			cfgManagerGetConfigMock: cfgManagerGetConfigMockResult{
				res: &config.Config{
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
				times: 1,
			},
			args: args{
				promAlertHook: &models.PrometheusAlertHook{
					Version: "4",
					Alerts: []*models.PrometheusAlertDetail{{
						Status: "firing",
						Labels: map[string]string{
							"alertname": "Alert1",
							"label1":    "value1",
						},
					}},
				},
			},
			cachetChangeComponentStatusMock: cachetChangeComponentStatusMockResult{
				input1: "Component1",
				input2: "",
				input3: "OPERATIONAL",
				times:  1,
			},
			metricsIncrementComponentManagedCounterMock: metricsIncrementComponentManagedCounterMockResult{
				input1: "OPERATIONAL",
				times:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			cfgMock := configmocks.NewMockManager(ctrl)
			metricsMock := metricsmocks.NewMockClient(ctrl)
			cachetMock := cachethqmocks.NewMockClient(ctrl)

			cfgMock.EXPECT().GetConfig().
				Return(tt.cfgManagerGetConfigMock.res).
				Times(tt.cfgManagerGetConfigMock.times)

			metricsMock.EXPECT().
				IncrementIncidentManagedCounter(
					tt.metricsIncrementIncidentManagedCounterMock.input1,
					tt.metricsIncrementIncidentManagedCounterMock.input2,
				).
				Times(
					tt.metricsIncrementIncidentManagedCounterMock.times,
				)

			metricsMock.EXPECT().
				IncrementComponentManagedCounter(
					tt.metricsIncrementComponentManagedCounterMock.input1,
				).
				Times(
					tt.metricsIncrementComponentManagedCounterMock.times,
				)

			cachetMock.EXPECT().
				CreateIncident(
					tt.cachetCreateIncidentMock.input1,
					tt.cachetCreateIncidentMock.input2,
					tt.cachetCreateIncidentMock.input3,
					tt.cachetCreateIncidentMock.input4,
					tt.cachetCreateIncidentMock.input5,
				).
				Return(
					tt.cachetCreateIncidentMock.err,
				).
				Times(
					tt.cachetCreateIncidentMock.times,
				)

			cachetMock.EXPECT().
				ChangeComponentStatus(
					tt.cachetChangeComponentStatusMock.input1,
					tt.cachetChangeComponentStatusMock.input2,
					tt.cachetChangeComponentStatusMock.input3,
				).
				Return(
					tt.cachetChangeComponentStatusMock.err,
				).
				Times(
					tt.cachetChangeComponentStatusMock.times,
				)

			ctx := &service{
				validate:    validator.New(),
				cfgManager:  cfgMock,
				cachethqCtx: cachetMock,
				metricsCtx:  metricsMock,
			}

			err := ctx.ManageHook(tt.args.promAlertHook)

			if (err != nil) != tt.wantErr {
				t.Errorf("service.ManageHook() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.errorString != "" && err != nil && err.Error() != tt.errorString {
				t.Errorf("service.ManageHook() error = %v, wantErr %v", err, tt.errorString)
				return
			}
		})
	}
}
