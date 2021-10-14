package cachethq

import (
	"testing"

	"github.com/andygrunwald/cachet"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/config"
)

func Test_getCachetHQIncidentStatus(t *testing.T) {
	type args struct {
		statusString string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Incident fixed status",
			args:    args{statusString: config.IncidentFixedStatus},
			want:    cachet.IncidentStatusFixed,
			wantErr: false,
		},
		{
			name:    "Incident identified status",
			args:    args{statusString: config.IncidentIdentifiedStatus},
			want:    cachet.IncidentStatusIdentified,
			wantErr: false,
		},
		{
			name:    "Incident investigating status",
			args:    args{statusString: config.IncidentInvestigatingStatus},
			want:    cachet.IncidentStatusInvestigating,
			wantErr: false,
		},
		{
			name:    "Incident watching status",
			args:    args{statusString: config.IncidentWatchingStatus},
			want:    cachet.IncidentStatusWatching,
			wantErr: false,
		},
		{
			name:    "Other case",
			args:    args{statusString: "OTHER_CASE"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCachetHQIncidentStatus(tt.args.statusString)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCachetHQIncidentStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getCachetHQIncidentStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCachetHQIncidentVisibility(t *testing.T) {
	type args struct {
		visible bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Visibility ok",
			args: args{visible: true},
			want: cachet.IncidentVisibilityPublic,
		},
		{
			name: "Visibility hidden",
			args: args{visible: false},
			want: cachet.IncidentVisibilityLoggedIn,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCachetHQIncidentVisibility(tt.args.visible); got != tt.want {
				t.Errorf("getCachetHQIncidentVisibility() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCachetHQComponentStatus(t *testing.T) {
	type args struct {
		statusString string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Component status Major outage",
			args:    args{statusString: config.ComponentMajorOutageStatus},
			want:    cachet.ComponentStatusMajorOutage,
			wantErr: false,
		},
		{
			name:    "Component status operational",
			args:    args{statusString: config.ComponentOperationalStatus},
			want:    cachet.ComponentStatusOperational,
			wantErr: false,
		},
		{
			name:    "Component status Partial outage",
			args:    args{statusString: config.ComponentPartialOutageStatus},
			want:    cachet.ComponentStatusPartialOutage,
			wantErr: false,
		},
		{
			name:    "Component status Performance issues",
			args:    args{statusString: config.ComponentPerformanceIssuesStatus},
			want:    cachet.ComponentStatusPerformanceIssues,
			wantErr: false,
		},
		{
			name:    "Other case",
			args:    args{statusString: "OTHER_CASE"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCachetHQComponentStatus(tt.args.statusString)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCachetHQComponentStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getCachetHQComponentStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
