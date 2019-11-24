package business

import (
	"testing"

	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheushook"
)

func Test_isAlertMatching(t *testing.T) {
	type args struct {
		matchingLabelsKeys []string
		matchingLabels     map[string]string
		alert              *prometheushook.PrometheusAlertDetail
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
				alert: &prometheushook.PrometheusAlertDetail{
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
				alert: &prometheushook.PrometheusAlertDetail{
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
				alert: &prometheushook.PrometheusAlertDetail{
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
				alert: &prometheushook.PrometheusAlertDetail{
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
				alert: &prometheushook.PrometheusAlertDetail{
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
				alert: &prometheushook.PrometheusAlertDetail{
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
