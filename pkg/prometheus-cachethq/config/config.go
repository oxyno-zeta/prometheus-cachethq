package config

// DefaultLogLevel Default log level.
const DefaultLogLevel = "info"

// DefaultLogFormat Default Log format.
const DefaultLogFormat = "json"

// DefaultPort Default port.
const DefaultPort = 8080

// DefaultInternalPort Default internal port.
const DefaultInternalPort = 9090

const (
	// IncidentInvestigatingStatus Incident investigating status.
	IncidentInvestigatingStatus = "INVESTIGATING"
	// IncidentIdentifiedStatus Incident identified status.
	IncidentIdentifiedStatus = "IDENTIFIED"
	// IncidentWatchingStatus Incident watching status.
	IncidentWatchingStatus = "WATCHING"
	// IncidentFixedStatus Incident fixed status.
	IncidentFixedStatus = "FIXED"
)

const (
	// ComponentOperationalStatus Component operational status.
	ComponentOperationalStatus = "OPERATIONAL"
	// ComponentPerformanceIssuesStatus Component performance issues status.
	ComponentPerformanceIssuesStatus = "PERFORMANCE_ISSUES"
	// ComponentPartialOutageStatus Component partial outage status.
	ComponentPartialOutageStatus = "PARTIAL_OUTAGE"
	// ComponentMajorOutageStatus Component major outage status.
	ComponentMajorOutageStatus = "MAJOR_OUTAGE"
)

// Config Configuration object.
type Config struct {
	Log            *LogConfig     `mapstructure:"log"`
	Tracing        *TracingConfig `mapstructure:"tracing"`
	Server         *ServerConfig  `mapstructure:"server"`
	InternalServer *ServerConfig  `mapstructure:"internalServer"`
	Cachet         *CachetConfig  `mapstructure:"cachet" validate:"required"`
	Targets        []*Target      `mapstructure:"targets" validate:"required,dive,required"`
}

// CachetConfig CachetHQ Configuration.
type CachetConfig struct {
	URL    string `mapstructure:"url" validate:"required,uri"`
	APIKey string `mapstructure:"apiKey" validate:"required"`
}

// Target Target configuration.
type Target struct {
	Component *TargetComponent `mapstructure:"component" validate:"required"`
	Alerts    []*TargetAlerts  `mapstructure:"alerts" validate:"required,dive,required"`
	Incident  *TargetIncident  `mapstructure:"incident" validate:"omitempty"`
}

// TargetIncident incident.
type TargetIncident struct {
	Name    string `mapstructure:"name" validate:"required"`
	Content string `mapstructure:"content" validate:"required"`
	Status  string `mapstructure:"status" validate:"required,oneof=INVESTIGATING IDENTIFIED WATCHING FIXED"`
	Public  bool   `mapstructure:"public"`
}

// TargetAlerts Target Prometheus alerts.
type TargetAlerts struct {
	Name   string            `mapstructure:"name" validate:"required_without_all=Labels"`
	Labels map[string]string `mapstructure:"labels" validate:"required_without_all=Name"`
}

// TargetComponent Target component.
type TargetComponent struct {
	Name      string `mapstructure:"name" validate:"required"`
	GroupName string `mapstructure:"groupName"`
	Status    string `mapstructure:"status" validate:"required,oneof=OPERATIONAL PERFORMANCE_ISSUES PARTIAL_OUTAGE MAJOR_OUTAGE"`
}

// TracingConfig represents the Tracing configuration structure.
type TracingConfig struct {
	Enabled       bool                   `mapstructure:"enabled"`
	LogSpan       bool                   `mapstructure:"logSpan"`
	FlushInterval string                 `mapstructure:"flushInterval"`
	UDPHost       string                 `mapstructure:"udpHost"`
	QueueSize     int                    `mapstructure:"queueSize"`
	FixedTags     map[string]interface{} `mapstructure:"fixedTags"`
}

// LogConfig Log configuration.
type LogConfig struct {
	Level    string `mapstructure:"level" validate:"required"`
	Format   string `mapstructure:"format" validate:"required"`
	FilePath string `mapstructure:"filePath"`
}

// ServerConfig Server configuration.
type ServerConfig struct {
	ListenAddr string            `mapstructure:"listenAddr"`
	Port       int               `mapstructure:"port" validate:"required"`
	CORS       *ServerCorsConfig `mapstructure:"cors" validate:"omitempty"`
}

// ServerCorsConfig Server CORS configuration.
type ServerCorsConfig struct {
	AllowOrigins            []string `mapstructure:"allowOrigins"`
	AllowMethods            []string `mapstructure:"allowMethods"`
	AllowHeaders            []string `mapstructure:"allowHeaders"`
	ExposeHeaders           []string `mapstructure:"exposeHeaders"`
	MaxAgeDuration          string   `mapstructure:"maxAgeDuration"`
	AllowCredentials        *bool    `mapstructure:"allowCredentials"`
	AllowWildcard           *bool    `mapstructure:"allowWildcard"`
	AllowBrowserExtensions  *bool    `mapstructure:"allowBrowserExtensions"`
	AllowWebSockets         *bool    `mapstructure:"allowWebSockets"`
	AllowFiles              *bool    `mapstructure:"allowFiles"`
	AllowAllOrigins         *bool    `mapstructure:"allowAllOrigins"`
	UseDefaultConfiguration bool     `mapstructure:"useDefaultConfiguration"`
}

// CredentialConfig Credential Configurations.
type CredentialConfig struct {
	Path  string `mapstructure:"path" validate:"required_without_all=Env Value"`
	Env   string `mapstructure:"env" validate:"required_without_all=Path Value"`
	Value string `mapstructure:"value" validate:"required_without_all=Path Env"`
}
