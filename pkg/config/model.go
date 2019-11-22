package config

// MainConfigPath Configuration path
const MainConfigPath = "conf/config.yaml"

// DefaultPort Default port
const DefaultPort = 8080

// DefaultInternalPort Default internal port
const DefaultInternalPort = 9090

// DefaultLogLevel Default log level
const DefaultLogLevel = "info"

// DefaultLogFormat Default Log format
const DefaultLogFormat = "json"

// Config Application Configuration
type Config struct {
	Log            *LogConfig    `koanf:"log"`
	InternalServer *ServerConfig `koanf:"internalServer"`
	Server         *ServerConfig `koanf:"server"`
	Cachet         *CachetConfig `koanf:"cachet" validate:"required"`
	Targets        []*Target     `koanf:"targets" validate:"required,dive,required"`
}

// ServerConfig Server configuration
type ServerConfig struct {
	ListenAddr string `koanf:"listenAddr"`
	Port       int    `koanf:"port" validate:"required"`
}

// LogConfig Log configuration
type LogConfig struct {
	Level  string `koanf:"level"`
	Format string `koanf:"format"`
}

// CachetConfig CachetHQ Configuration
type CachetConfig struct {
	URL    string `koanf:"url" validate:"required,uri"`
	APIKey string `koanf:"apiKey" validate:"required"`
}

// Target Target configuration
type Target struct {
	Component *TargetComponent `koanf:"component" validate:"required"`
	Alerts    []*TargetAlerts  `koanf:"alerts" validate:"required,dive,required"`
	Incident  *TargetIncident  `koanf:"incident" validate:"omitempty"`
}

// TargetIncident incident
type TargetIncident struct {
	Name    string `koanf:"name" validate:"required"`
	Content string `koanf:"content" validate:"required"`
	Status  string `koanf:"status" validate:"required,oneof=INVESTIGATING IDENTIFIED WATCHING FIXED"`
	Public  bool   `koanf:"public"`
}

// TargetAlerts Target Prometheus alerts
type TargetAlerts struct {
	Name   string            `koanf:"name" validate:"required_without_all=Labels"`
	Labels map[string]string `koanf:"labels" validate:"required_without_all=Name"`
}

// TargetComponent Target component
type TargetComponent struct {
	Name   string `koanf:"name" validate:"required"`
	Status string `koanf:"status" validate:"required,oneof=OPERATIONAL PERFORMANCE_ISSUES PARTIAL_OUTAGE MAJOR_OUTAGE"`
}

const (
	// IncidentInvestigatingStatus Incident investigating status
	IncidentInvestigatingStatus = "INVESTIGATING"
	// IncidentIdentifiedStatus Incident identified status
	IncidentIdentifiedStatus = "IDENTIFIED"
	// IncidentWatchingStatus Incident watching status
	IncidentWatchingStatus = "WATCHING"
	// IncidentFixedStatus Incident fixed status
	IncidentFixedStatus = "FIXED"
)
const (
	// ComponentOperationalStatus Component operational status
	ComponentOperationalStatus = "OPERATIONAL"
	// ComponentPerformanceIssuesStatus Component performance issues status
	ComponentPerformanceIssuesStatus = "PERFORMANCE_ISSUES"
	// ComponentPartialOutageStatus Component partial outage status
	ComponentPartialOutageStatus = "PARTIAL_OUTAGE"
	// ComponentMajorOutageStatus Component major outage status
	ComponentMajorOutageStatus = "MAJOR_OUTAGE"
)
