// +build integration

package config

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"

	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/log"
	"github.com/stretchr/testify/assert"
)

func Test_managercontext_Load(t *testing.T) {
	tests := []struct {
		name           string
		configs        map[string]string
		envVariables   map[string]string
		secretFiles    map[string]string
		expectedResult *Config
		wantErr        bool
	}{
		{
			name: "Configuration not found",
			configs: map[string]string{
				"config": "",
			},
			wantErr: true,
		},
		{
			name: "Not a yaml",
			configs: map[string]string{
				"config.yaml": "notayaml",
			},
			wantErr: true,
		},
		{
			name: "Empty",
			configs: map[string]string{
				"config.yaml": "",
			},
			wantErr: true,
		},
		{
			name: "default config",
			configs: map[string]string{
				"log.yaml": `
log:
  level: error
  format: text
`,
				"database.yaml": `
database:
  connectionUrl:
    value: host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable

`,
				"cachet.yaml": `
cachet:
  url: http://localhost:8000
  apiKey: fake

targets:
  - component:
      name: component1
      status: PARTIAL_OUTAGE
    alerts:
      - name: test1
`,
			},
			expectedResult: &Config{
				Log: &LogConfig{
					Format: "text",
					Level:  "error",
				},
				Tracing:        &TracingConfig{Enabled: false},
				Server:         &ServerConfig{Port: 8080},
				InternalServer: &ServerConfig{Port: 9090},
				Cachet: &CachetConfig{
					URL:    "http://localhost:8000",
					APIKey: "fake",
				},
				Targets: []*Target{{
					Component: &TargetComponent{
						Name:   "component1",
						Status: "PARTIAL_OUTAGE",
					},
					Alerts: []*TargetAlerts{{
						Name: "test1",
					}},
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir, err := ioutil.TempDir("", "config")
			if err != nil {
				t.Error(err)
				return
			}

			defer os.RemoveAll(dir) // clean up
			for k, v := range tt.configs {
				tmpfn := filepath.Join(dir, k)
				err = ioutil.WriteFile(tmpfn, []byte(v), 0666)
				if err != nil {
					t.Error(err)
					return
				}
			}

			// Set environment variables
			for k, v := range tt.envVariables {
				os.Setenv(k, v)
				defer os.Unsetenv(k)
			}

			// Create secret files
			for k, v := range tt.secretFiles {
				dirToCr := filepath.Dir(k)
				err = os.MkdirAll(dirToCr, 0666)
				if err != nil {
					t.Error(err)
					return
				}
				err = ioutil.WriteFile(k, []byte(v), 0666)
				if err != nil {
					t.Error(err)
					return
				}
				defer os.Remove(k)
			}

			// Change var for main configuration file
			mainConfigFolderPath = dir

			ctx := &managercontext{
				logger: log.NewLogger(),
			}

			// Load config
			err = ctx.Load()

			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Get configuration
			res := ctx.GetConfig()

			assert.Equal(t, tt.expectedResult, res)
		})
	}
}

func Test_Load_reload_config(t *testing.T) {
	// Channel for wait watch
	waitCh := make(chan bool)

	dir, err := ioutil.TempDir("", "config-reload")
	assert.NoError(t, err)

	configs := map[string]string{
		"log.yaml": `
log:
  level: error
`,
		"database.yaml": `
database:
  connectionUrl:
    value: host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable

`,
		"tracing.yaml": `
tracing:
  enabled: true
`,
		"cachet.yaml": `
cachet:
  url: http://localhost:8000
  apiKey: fake

targets:
  - component:
      name: component1
      status: PARTIAL_OUTAGE
    alerts:
      - name: test1
`,
	}

	defer os.RemoveAll(dir) // clean up
	for k, v := range configs {
		tmpfn := filepath.Join(dir, k)
		err = ioutil.WriteFile(tmpfn, []byte(v), 0666)
		assert.NoError(t, err)
	}

	secretFiles := map[string]string{
		os.TempDir() + "/secret1": "VALUE1",
	}
	// Create secret files
	for k, v := range secretFiles {
		dirToCr := filepath.Dir(k)
		err = os.MkdirAll(dirToCr, 0666)
		assert.NoError(t, err)
		err = ioutil.WriteFile(k, []byte(v), 0666)
		assert.NoError(t, err)
		defer os.Remove(k)
	}

	// Change var for main configuration file
	mainConfigFolderPath = dir

	ctx := &managercontext{
		logger: log.NewLogger(),
	}

	ctx.AddOnChangeHook(func() {
		waitCh <- true
	})

	// Load config
	err = ctx.Load()
	assert.NoError(t, err)
	// Get configuration
	res := ctx.GetConfig()

	assert.Equal(t, &Config{
		Log: &LogConfig{
			Level:  "error",
			Format: "json",
		},
		Server: &ServerConfig{
			Port: 8080,
		},
		InternalServer: &ServerConfig{
			Port: 9090,
		},
		Tracing: &TracingConfig{Enabled: true},
		Cachet: &CachetConfig{
			URL:    "http://localhost:8000",
			APIKey: "fake",
		},
		Targets: []*Target{{
			Component: &TargetComponent{
				Name:   "component1",
				Status: "PARTIAL_OUTAGE",
			},
			Alerts: []*TargetAlerts{{
				Name: "test1",
			}},
		}},
	}, res)

	configs = map[string]string{
		"log.yaml": `
log:
  level: debug
  format: text
`,
	}

	defer os.RemoveAll(dir) // clean up
	for k, v := range configs {
		tmpfn := filepath.Join(dir, k)
		err = ioutil.WriteFile(tmpfn, []byte(v), 0666)
		assert.NoError(t, err)
	}

	select {
	case <-waitCh:
		// Get configuration
		res = ctx.GetConfig()

		assert.Equal(t, &Config{
			Log: &LogConfig{
				Level:  "debug",
				Format: "text",
			},
			Server: &ServerConfig{
				Port: 8080,
			},
			InternalServer: &ServerConfig{
				Port: 9090,
			},
			Tracing: &TracingConfig{Enabled: true},
			Cachet: &CachetConfig{
				URL:    "http://localhost:8000",
				APIKey: "fake",
			},
			Targets: []*Target{{
				Component: &TargetComponent{
					Name:   "component1",
					Status: "PARTIAL_OUTAGE",
				},
				Alerts: []*TargetAlerts{{
					Name: "test1",
				}},
			}},
		}, res)
		return
	case <-time.After(5 * time.Second):
		assert.FailNow(t, "shouldn't call this")
	}
}

func Test_Load_reload_config_with_wrong_config(t *testing.T) {
	// Channel for wait watch
	waitCh := make(chan bool)

	dir, err := ioutil.TempDir("", "config-reload-wrong-config")
	assert.NoError(t, err)

	configs := map[string]string{
		"log.yaml": `
log:
  level: error
  format: text
`,
		"database.yaml": `
database:
  connectionUrl:
    value: host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable

`,
		"tracing.yaml": `
tracing:
  enabled: true
`,
		"cachet.yaml": `
cachet:
  url: http://localhost:8000
  apiKey: fake

targets:
  - component:
      name: component1
      status: PARTIAL_OUTAGE
    alerts:
      - name: test1
`,
	}

	defer os.RemoveAll(dir) // clean up
	for k, v := range configs {
		tmpfn := filepath.Join(dir, k)
		err = ioutil.WriteFile(tmpfn, []byte(v), 0666)
		assert.NoError(t, err)
	}

	secretFiles := map[string]string{
		os.TempDir() + "/secret1": "VALUE1",
	}
	// Create secret files
	for k, v := range secretFiles {
		dirToCr := filepath.Dir(k)
		err = os.MkdirAll(dirToCr, 0666)
		assert.NoError(t, err)
		err = ioutil.WriteFile(k, []byte(v), 0666)
		assert.NoError(t, err)
		defer os.Remove(k)
	}

	// Change var for main configuration file
	mainConfigFolderPath = dir

	ctx := &managercontext{
		logger: log.NewLogger(),
	}

	ctx.AddOnChangeHook(func() {
		waitCh <- true
	})

	// Load config
	err = ctx.Load()
	assert.NoError(t, err)
	// Get configuration
	res := ctx.GetConfig()

	assert.Equal(t, &Config{
		Log: &LogConfig{
			Level:  "error",
			Format: "text",
		},
		Server: &ServerConfig{
			Port: 8080,
		},
		InternalServer: &ServerConfig{
			Port: 9090,
		},
		Tracing: &TracingConfig{Enabled: true},
		Cachet: &CachetConfig{
			URL:    "http://localhost:8000",
			APIKey: "fake",
		},
		Targets: []*Target{{
			Component: &TargetComponent{
				Name:   "component1",
				Status: "PARTIAL_OUTAGE",
			},
			Alerts: []*TargetAlerts{{
				Name: "test1",
			}},
		}},
	}, res)

	configs = map[string]string{
		"log.yaml": `
configuration with error
`,
	}

	defer os.RemoveAll(dir) // clean up
	for k, v := range configs {
		tmpfn := filepath.Join(dir, k)
		err = ioutil.WriteFile(tmpfn, []byte(v), 0666)
		assert.NoError(t, err)
	}

	select {
	case <-waitCh:
		assert.FailNow(t, "shouldn't call this")
		return
	case <-time.After(5 * time.Second):
		// Get configuration
		res = ctx.GetConfig()

		assert.Equal(t, &Config{
			Log: &LogConfig{
				Level:  "error",
				Format: "text",
			},
			Server: &ServerConfig{
				Port: 8080,
			},
			InternalServer: &ServerConfig{
				Port: 9090,
			},
			Tracing: &TracingConfig{Enabled: true},
			Cachet: &CachetConfig{
				URL:    "http://localhost:8000",
				APIKey: "fake",
			},
			Targets: []*Target{{
				Component: &TargetComponent{
					Name:   "component1",
					Status: "PARTIAL_OUTAGE",
				},
				Alerts: []*TargetAlerts{{
					Name: "test1",
				}},
			}},
		}, res)
	}
}

func Test_Load_reload_config_ignore_hidden_file_and_directory(t *testing.T) {
	// Channel for wait watch
	waitCh := make(chan bool)

	dir, err := ioutil.TempDir("", "config-reload-ignore")
	assert.NoError(t, err)
	err = os.MkdirAll(path.Join(dir, "dir1"), os.ModePerm)
	assert.NoError(t, err)

	configs := map[string]string{
		"..log.yaml": `
log:
  level: error
`,
		".log2.yaml": `
log:
  format: fake
`,
		"dir1/log2.yaml": `
server:
  port: 8181
`,
		"log.yaml": `
log:
  format: humanfriendly
  level: debug
`,
		"database.yaml": `
database:
  connectionUrl:
    value: host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable

`,
		"cachet.yaml": `
cachet:
  url: http://localhost:8000
  apiKey: fake

targets:
  - component:
      name: component1
      status: PARTIAL_OUTAGE
    alerts:
      - name: test1
`,
	}

	defer os.RemoveAll(dir) // clean up
	for k, v := range configs {
		tmpfn := filepath.Join(dir, k)
		err = ioutil.WriteFile(tmpfn, []byte(v), 0666)
		assert.NoError(t, err)
	}

	secretFiles := map[string]string{
		os.TempDir() + "/secret1": "VALUE1",
	}
	// Create secret files
	for k, v := range secretFiles {
		dirToCr := filepath.Dir(k)
		err = os.MkdirAll(dirToCr, 0666)
		assert.NoError(t, err)
		err = ioutil.WriteFile(k, []byte(v), 0666)
		assert.NoError(t, err)
		defer os.Remove(k)
	}

	// Change var for main configuration file
	mainConfigFolderPath = dir

	ctx := &managercontext{
		logger: log.NewLogger(),
	}

	ctx.AddOnChangeHook(func() {
		waitCh <- true
	})

	// Load config
	err = ctx.Load()
	assert.NoError(t, err)
	// Get configuration
	res := ctx.GetConfig()

	assert.Equal(t, &Config{
		Log: &LogConfig{
			Level:  "debug",
			Format: "humanfriendly",
		},
		Server: &ServerConfig{
			Port: 8080,
		},
		InternalServer: &ServerConfig{
			Port: 9090,
		},
		Tracing: &TracingConfig{Enabled: false},
		Cachet: &CachetConfig{
			URL:    "http://localhost:8000",
			APIKey: "fake",
		},
		Targets: []*Target{{
			Component: &TargetComponent{
				Name:   "component1",
				Status: "PARTIAL_OUTAGE",
			},
			Alerts: []*TargetAlerts{{
				Name: "test1",
			}},
		}},
	}, res)
}
