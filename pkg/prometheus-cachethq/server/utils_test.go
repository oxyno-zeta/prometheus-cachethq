// +build unit

package server

import (
	"testing"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/config"
	"github.com/stretchr/testify/assert"
)

func Test_generateCORSConfiguration(t *testing.T) {
	starBool := func(v bool) *bool { return &v }
	type args struct {
		cfg *config.ServerCorsConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *cors.Config
		wantErr bool
	}{
		{
			name: "nil",
			want: nil,
		},
		{
			name: "Empty configuration",
			args: args{
				cfg: &config.ServerCorsConfig{},
			},
			wantErr: true,
		},
		{
			name: "Use default configuration only",
			args: args{
				cfg: &config.ServerCorsConfig{
					UseDefaultConfiguration: true,
				},
			},
			wantErr: true,
		},
		{
			name: "Parse duration error",
			args: args{
				cfg: &config.ServerCorsConfig{
					MaxAgeDuration: "fake",
				},
			},
			wantErr: true,
		},
		{
			name: "default configuration and allow all origins only",
			args: args{
				cfg: &config.ServerCorsConfig{
					UseDefaultConfiguration: true,
					AllowAllOrigins:         starBool(true),
				},
			},
			want: &cors.Config{
				AllowAllOrigins: true,
				AllowMethods: []string{
					"GET",
					"POST",
					"PUT",
					"PATCH",
					"DELETE",
					"HEAD",
				},
				AllowHeaders: []string{
					"Origin",
					"Content-Length",
					"Content-Type",
				},
				MaxAge: 12 * time.Hour,
			},
		},
		{
			name: "all settings with wildcard",
			args: args{
				cfg: &config.ServerCorsConfig{
					UseDefaultConfiguration: true,
					AllowBrowserExtensions:  starBool(true),
					AllowCredentials:        starBool(true),
					AllowFiles:              starBool(true),
					AllowWebSockets:         starBool(true),
					AllowWildcard:           starBool(true),
					AllowHeaders:            []string{"fake"},
					AllowMethods:            []string{"FAKE"},
					AllowOrigins:            []string{"http://*.com"},
					ExposeHeaders:           []string{"EXPOSED"},
					MaxAgeDuration:          "2h",
				},
			},
			want: &cors.Config{
				AllowBrowserExtensions: true,
				AllowCredentials:       true,
				AllowFiles:             true,
				AllowWebSockets:        true,
				AllowWildcard:          true,
				ExposeHeaders:          []string{"EXPOSED"},
				AllowOrigins:           []string{"http://*.com"},
				AllowMethods:           []string{"FAKE"},
				AllowHeaders:           []string{"fake"},
				MaxAge:                 2 * time.Hour,
			},
		},
		{
			name: "all settings without wildcard",
			args: args{
				cfg: &config.ServerCorsConfig{
					UseDefaultConfiguration: true,
					AllowBrowserExtensions:  starBool(true),
					AllowCredentials:        starBool(true),
					AllowFiles:              starBool(true),
					AllowWebSockets:         starBool(true),
					AllowHeaders:            []string{"fake"},
					AllowMethods:            []string{"FAKE"},
					AllowOrigins:            []string{"http://localhost"},
					ExposeHeaders:           []string{"EXPOSED"},
					MaxAgeDuration:          "2h",
				},
			},
			want: &cors.Config{
				AllowBrowserExtensions: true,
				AllowCredentials:       true,
				AllowFiles:             true,
				AllowWebSockets:        true,
				ExposeHeaders:          []string{"EXPOSED"},
				AllowOrigins:           []string{"http://localhost"},
				AllowMethods:           []string{"FAKE"},
				AllowHeaders:           []string{"fake"},
				MaxAge:                 2 * time.Hour,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateCORSConfiguration(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateCORSConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
