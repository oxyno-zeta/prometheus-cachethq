//+build unit

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_loadBusinessDefaultValues(t *testing.T) {
	type args struct {
		out *Config
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		expectedCfg *Config
	}{
		{
			name: "Empty",
			args: args{out: &Config{}},
			expectedCfg: &Config{
				Tracing: &TracingConfig{Enabled: false},
			},
		},
		{
			name: "oidc",
			args: args{
				out: &Config{
					OIDCAuthentication: &OIDCAuthConfig{},
				},
			},
			expectedCfg: &Config{
				Tracing: &TracingConfig{Enabled: false},
				OIDCAuthentication: &OIDCAuthConfig{
					Scopes:     DefaultOIDCScopes,
					CookieName: DefaultCookieName,
				},
			},
		},
		{
			name: "opa",
			args: args{
				out: &Config{
					OPAServerAuthorization: &OPAServerAuthorization{},
				},
			},
			expectedCfg: &Config{
				Tracing: &TracingConfig{Enabled: false},
				OPAServerAuthorization: &OPAServerAuthorization{
					Tags: map[string]string{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := loadBusinessDefaultValues(tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("loadBusinessDefaultValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.expectedCfg, tt.args.out)
		})
	}
}
