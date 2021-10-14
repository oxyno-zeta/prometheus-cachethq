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
