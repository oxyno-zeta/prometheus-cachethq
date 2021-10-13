// +build integration

package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/andygrunwald/cachet"
	"github.com/golang/mock/gomock"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/business"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/cachethq"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/config"
	cmocks "github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/config/mocks"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/log"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/signalhandler"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/tracing"
	"github.com/stretchr/testify/assert"
)

func TestPublicRouter(t *testing.T) {
	body := `{
  "receiver": "webhook",
  "status": "firing",
  "alerts": [
    {
      "status": "firing",
      "labels": {
        "alertname": "test1",
        "dc": "eu-west-1",
        "instance": "localhost:9090",
        "job": "prometheus24"
      },
      "annotations": {
        "description": "some description"
      },
      "startsAt": "2018-08-03T09:52:26.739266876+02:00",
      "endsAt": "0001-01-01T00:00:00Z",
      "generatorURL": "http://my-laptop:9090/graph?g0.expr=go_memstats_alloc_bytes+%3E+0\u0026g0.tab=1"
    }
  ],
  "groupLabels": {
    "alertname": "Test",
    "job": "prometheus24"
  },
  "commonLabels": {
    "alertname": "Test",
    "dc": "eu-west-1",
    "instance": "localhost:9090",
    "job": "prometheus24"
  },
  "commonAnnotations": {
    "description": "some description"
  },
  "externalURL": "http://my-laptop:9093",
  "version": "4",
  "groupKey": "{}:{alertname=\"Test\", job=\"prometheus24\"}"
}`
	cachetCfg := &config.CachetConfig{
		URL:    "http://localhost:8000",
		APIKey: "8I0tUhh4WUZIL7eyFUMX",
	}
	tracingCfg := &config.TracingConfig{}
	serverCfg := &config.ServerConfig{}

	type args struct {
		cfg *config.Config
	}
	tests := []struct {
		name                     string
		args                     args
		inputMethod              string
		inputURL                 string
		inputBody                string
		inputHeaders             map[string]string
		expectedCode             int
		expectedBody             string
		expectedHeaders          map[string]string
		expectedComponentsStatus map[int]int
		expectedIncidents        []*cachet.Incident
		wantErr                  bool
	}{
		{
			name: "Get a not found path",
			args: args{
				cfg: &config.Config{
					Cachet:  cachetCfg,
					Tracing: tracingCfg,
					Server:  serverCfg,
				},
			},
			inputMethod:  "GET",
			inputURL:     "http://localhost/not-found/",
			expectedCode: 404,
			expectedBody: "{\"error\":\"404 not found\",\"extensions\":{\"code\":\"NOT_FOUND\"}}",
			expectedHeaders: map[string]string{
				"Content-Type": "application/json; charset=utf-8",
			},
		},
		{
			name: "shouldn't match anything",
			args: args{
				cfg: &config.Config{
					Cachet:  cachetCfg,
					Tracing: tracingCfg,
					Server:  serverCfg,
					Targets: []*config.Target{{
						Alerts:    []*config.TargetAlerts{{Name: "not-matching"}},
						Component: &config.TargetComponent{Name: "component1", Status: "PARTIAL_OUTAGE"},
					}},
				},
			},
			inputMethod:  "POST",
			inputURL:     "http://localhost/prometheus/webhook",
			inputBody:    body,
			expectedCode: 204,
			expectedComponentsStatus: map[int]int{
				1: cachet.ComponentStatusOperational,
				2: cachet.ComponentStatusOperational,
				3: cachet.ComponentStatusOperational,
				4: cachet.ComponentStatusOperational,
				5: cachet.ComponentStatusOperational,
			},
		},
		{
			name: "should match the alert name for 1 component",
			args: args{
				cfg: &config.Config{
					Cachet:  cachetCfg,
					Tracing: tracingCfg,
					Server:  serverCfg,
					Targets: []*config.Target{{
						Alerts:    []*config.TargetAlerts{{Name: "test1"}},
						Component: &config.TargetComponent{Name: "component1", Status: "PARTIAL_OUTAGE"},
					}},
				},
			},
			inputMethod:  "POST",
			inputURL:     "http://localhost/prometheus/webhook",
			inputBody:    body,
			expectedCode: 204,
			expectedComponentsStatus: map[int]int{
				1: cachet.ComponentStatusOperational,
				2: cachet.ComponentStatusOperational,
				3: cachet.ComponentStatusOperational,
				4: cachet.ComponentStatusPartialOutage,
				5: cachet.ComponentStatusOperational,
			},
		},
		{
			name: "should match the alert name for 2 components",
			args: args{
				cfg: &config.Config{
					Cachet:  cachetCfg,
					Tracing: tracingCfg,
					Server:  serverCfg,
					Targets: []*config.Target{
						{
							Alerts:    []*config.TargetAlerts{{Name: "test1"}},
							Component: &config.TargetComponent{Name: "component1", Status: "PARTIAL_OUTAGE"},
						},
						{
							Alerts:    []*config.TargetAlerts{{Name: "test1"}},
							Component: &config.TargetComponent{Name: "triple1", Status: "PARTIAL_OUTAGE"},
						},
					},
				},
			},
			inputMethod:  "POST",
			inputURL:     "http://localhost/prometheus/webhook",
			inputBody:    body,
			expectedCode: 204,
			expectedComponentsStatus: map[int]int{
				1: cachet.ComponentStatusPartialOutage,
				2: cachet.ComponentStatusOperational,
				3: cachet.ComponentStatusOperational,
				4: cachet.ComponentStatusPartialOutage,
				5: cachet.ComponentStatusOperational,
			},
		},
		{
			name: "should match the alert name for 2 components (one with alert name and the other with label)",
			args: args{
				cfg: &config.Config{
					Cachet:  cachetCfg,
					Tracing: tracingCfg,
					Server:  serverCfg,
					Targets: []*config.Target{
						{
							Alerts:    []*config.TargetAlerts{{Name: "test1"}},
							Component: &config.TargetComponent{Name: "component1", Status: "PARTIAL_OUTAGE"},
						},
						{
							Alerts:    []*config.TargetAlerts{{Labels: map[string]string{"job": "prometheus24"}}},
							Component: &config.TargetComponent{Name: "triple1", Status: "PARTIAL_OUTAGE"},
						},
					},
				},
			},
			inputMethod:  "POST",
			inputURL:     "http://localhost/prometheus/webhook",
			inputBody:    body,
			expectedCode: 204,
			expectedComponentsStatus: map[int]int{
				1: cachet.ComponentStatusPartialOutage,
				2: cachet.ComponentStatusOperational,
				3: cachet.ComponentStatusOperational,
				4: cachet.ComponentStatusPartialOutage,
				5: cachet.ComponentStatusOperational,
			},
		},
		{
			name: "should match component with group",
			args: args{
				cfg: &config.Config{
					Cachet:  cachetCfg,
					Tracing: tracingCfg,
					Server:  serverCfg,
					Targets: []*config.Target{
						{
							Alerts:    []*config.TargetAlerts{{Name: "test1"}},
							Component: &config.TargetComponent{Name: "triple1", GroupName: "group1", Status: "PARTIAL_OUTAGE"},
						},
					},
				},
			},
			inputMethod:  "POST",
			inputURL:     "http://localhost/prometheus/webhook",
			inputBody:    body,
			expectedCode: 204,
			expectedComponentsStatus: map[int]int{
				1: cachet.ComponentStatusOperational,
				2: cachet.ComponentStatusPartialOutage,
				3: cachet.ComponentStatusOperational,
				4: cachet.ComponentStatusOperational,
				5: cachet.ComponentStatusOperational,
			},
		},
		{
			name: "should fail to match component with unknown group",
			args: args{
				cfg: &config.Config{
					Cachet:  cachetCfg,
					Tracing: tracingCfg,
					Server:  serverCfg,
					Targets: []*config.Target{
						{
							Alerts:    []*config.TargetAlerts{{Name: "test1"}},
							Component: &config.TargetComponent{Name: "triple1", GroupName: "fake", Status: "PARTIAL_OUTAGE"},
						},
					},
				},
			},
			inputMethod:  "POST",
			inputURL:     "http://localhost/prometheus/webhook",
			inputBody:    body,
			expectedCode: 404,
			expectedBody: "{\"error\":\"component group not found\",\"extensions\":{\"code\":\"NOT_FOUND\"}}",
			expectedComponentsStatus: map[int]int{
				1: cachet.ComponentStatusOperational,
				2: cachet.ComponentStatusOperational,
				3: cachet.ComponentStatusOperational,
				4: cachet.ComponentStatusOperational,
				5: cachet.ComponentStatusOperational,
			},
		},
		{
			name: "should fail to match a not found component",
			args: args{
				cfg: &config.Config{
					Cachet:  cachetCfg,
					Tracing: tracingCfg,
					Server:  serverCfg,
					Targets: []*config.Target{
						{
							Alerts:    []*config.TargetAlerts{{Name: "test1"}},
							Component: &config.TargetComponent{Name: "fake", Status: "PARTIAL_OUTAGE"},
						},
					},
				},
			},
			inputMethod:  "POST",
			inputURL:     "http://localhost/prometheus/webhook",
			inputBody:    body,
			expectedCode: 404,
			expectedBody: "{\"error\":\"component not found\",\"extensions\":{\"code\":\"NOT_FOUND\"}}",
			expectedComponentsStatus: map[int]int{
				1: cachet.ComponentStatusOperational,
				2: cachet.ComponentStatusOperational,
				3: cachet.ComponentStatusOperational,
				4: cachet.ComponentStatusOperational,
				5: cachet.ComponentStatusOperational,
			},
		},
		{
			name: "should create an incident",
			args: args{
				cfg: &config.Config{
					Cachet:  cachetCfg,
					Tracing: tracingCfg,
					Server:  serverCfg,
					Targets: []*config.Target{
						{
							Alerts:    []*config.TargetAlerts{{Name: "test1"}},
							Component: &config.TargetComponent{Name: "triple1", GroupName: "group2", Status: "PARTIAL_OUTAGE"},
							Incident: &config.TargetIncident{
								Name:    "fake incident",
								Content: "content",
								Status:  "INVESTIGATING",
								Public:  true,
							},
						},
					},
				},
			},
			inputMethod:  "POST",
			inputURL:     "http://localhost/prometheus/webhook",
			inputBody:    body,
			expectedCode: 204,
			expectedComponentsStatus: map[int]int{
				1: cachet.ComponentStatusOperational,
				2: cachet.ComponentStatusOperational,
				3: cachet.ComponentStatusPartialOutage,
				4: cachet.ComponentStatusOperational,
				5: cachet.ComponentStatusOperational,
			},
			expectedIncidents: []*cachet.Incident{{
				Name:        "fake incident",
				Message:     "content",
				Visible:     cachet.IncidentVisibilityPublic,
				Status:      cachet.IncidentStatusInvestigating,
				ComponentID: 3,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create new cachet client
			client, err := cachet.NewClient(cachetCfg.URL, nil)
			assert.NoError(t, err)
			// Set authentication token.
			client.Authentication.SetTokenAuth(cachetCfg.APIKey)

			// Reset cachethq
			for i := 1; i < 6; i += 1 {
				c, _, err := client.Components.Get(i)
				assert.NoError(t, err)
				c.Status = cachet.ComponentStatusOperational
				_, _, err = client.Components.Update(i, c)
				assert.NoError(t, err)
				ic, _, err := client.Incidents.GetAll(&cachet.IncidentsQueryParams{QueryOptions: cachet.QueryOptions{PerPage: 50}})
				assert.NoError(t, err)
				for _, it := range ic.Incidents {
					_, err = client.Incidents.Delete(it.ID)
					assert.NoError(t, err)
				}
			}

			// Create go mock controller
			ctrl := gomock.NewController(t)
			cfgManagerMock := cmocks.NewMockManager(ctrl)

			// Load configuration in manager
			cfgManagerMock.EXPECT().GetConfig().AnyTimes().Return(tt.args.cfg)

			logger := log.NewLogger()
			// Create tracing service
			tsvc, err := tracing.New(cfgManagerMock, logger)
			assert.NoError(t, err)

			// Create new cachethq client
			cachetCl := cachethq.NewInstance(cfgManagerMock)
			// Initialize
			err = cachetCl.Initialize()
			// Check error
			assert.NoError(t, err)

			// Create business services
			busServices := business.NewServices(
				logger,
				cfgManagerMock,
				metricsCtx,
				cachetCl,
			)

			signalHandlerSvcCl := signalhandler.NewClient(logger, false, []os.Signal{})

			svr := &Server{
				logger:           logger,
				cfgManager:       cfgManagerMock,
				metricsCl:        metricsCtx,
				tracingSvc:       tsvc,
				busiServices:     busServices,
				signalHandlerSvc: signalHandlerSvcCl,
			}
			got, err := svr.generateRouter()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateRouter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// If want error at this moment => stop
			if tt.wantErr {
				return
			}
			w := httptest.NewRecorder()
			var req *http.Request

			if tt.inputBody != "" {
				req, err = http.NewRequest(
					tt.inputMethod,
					tt.inputURL,
					bytes.NewBufferString(tt.inputBody),
				)
			} else {
				req, err = http.NewRequest(
					tt.inputMethod,
					tt.inputURL,
					nil,
				)
			}
			if err != nil {
				t.Error(err)
				return
			}

			// Add headers
			if tt.inputHeaders != nil {
				for key, value := range tt.inputHeaders {
					req.Header.Set(key, value)
				}
			}

			got.ServeHTTP(w, req)

			if tt.expectedBody != "" {
				body := w.Body.String()
				assert.Equal(t, tt.expectedBody, body)
			}

			if tt.expectedHeaders != nil {
				for key, val := range tt.expectedHeaders {
					wheader := w.HeaderMap.Get(key)
					assert.Equal(t, val, wheader, key)
				}
			}

			assert.Equal(t, tt.expectedCode, w.Code)

			for k, v := range tt.expectedComponentsStatus {
				c, _, err := client.Components.Get(k)
				assert.NoError(t, err)
				assert.Equal(t, v, c.Status)
			}

			for _, v := range tt.expectedIncidents {
				incRes, _, err := client.Incidents.GetAll(&cachet.IncidentsQueryParams{
					Name: v.Name,
				})
				assert.NoError(t, err)
				assert.Len(t, incRes.Incidents, 1)

				if len(incRes.Incidents) != 0 {
					inc := incRes.Incidents[0]
					assert.Equal(t, v.ComponentID, inc.ComponentID)
					assert.Equal(t, v.Message, inc.Message)
					assert.Equal(t, v.Name, inc.Name)
					assert.Equal(t, v.Status, inc.Status)
					assert.Equal(t, v.Visible, inc.Visible)
				}
			}
		})
	}
}
