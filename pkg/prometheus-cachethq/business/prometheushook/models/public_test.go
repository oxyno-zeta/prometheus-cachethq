package models

import "testing"

func TestPrometheusAlertHook_Validate(t *testing.T) {
	type fields struct {
		Version string
		Alerts  []*PrometheusAlertDetail
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Valid",
			fields: fields{
				Version: "4",
				Alerts:  []*PrometheusAlertDetail{{Status: "fake"}},
			},
			wantErr: false,
		},
		{
			name: "Not valid",
			fields: fields{
				Version: "3",
				Alerts:  []*PrometheusAlertDetail{{Status: "fake"}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pah := &PrometheusAlertHook{
				Version: tt.fields.Version,
				Alerts:  tt.fields.Alerts,
			}
			if err := pah.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("PrometheusAlertHook.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
