package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/sirupsen/logrus"
)

var k = koanf.New(".")
var validate = validator.New()

// Load Load configuration
func Load() (*Config, error) {
	// Load default configuration
	err := k.Load(confmap.Provider(map[string]interface{}{
		"log.level":           DefaultLogLevel,
		"log.format":          DefaultLogFormat,
		"server.port":         DefaultPort,
		"internalServer.port": DefaultInternalPort,
	}, "."), nil)
	if err != nil {
		return nil, err
	}

	// Try to load main configuration file
	err = k.Load(file.Provider(MainConfigPath), yaml.Parser())
	if err != nil {
		return nil, err
	}

	// Prepare configuration object
	var out Config
	// Quick unmarshal.
	err = k.Unmarshal("", &out)
	if err != nil {
		return nil, err
	}

	// Configuration validation
	err = validate.Struct(out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

// ConfigureLogger Configure logger instance
func ConfigureLogger(logger *logrus.Logger, logConfig *LogConfig) error {
	// Manage log format
	if logConfig.Format == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{})
	}

	// Manage log level
	lvl, err := logrus.ParseLevel(logConfig.Level)
	if err != nil {
		return err
	}
	logger.SetLevel(lvl)

	return nil
}
