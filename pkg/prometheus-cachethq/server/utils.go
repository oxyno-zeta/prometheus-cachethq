package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/prometheus-cachethq/pkg/prometheus-cachethq/config"
	"github.com/pkg/errors"
)

func manageCORS(router gin.IRoutes, cfg *config.ServerConfig) error {
	// Generate configuration
	ccfg, err := generateCORSConfiguration(cfg.CORS)
	// Check error
	if err != nil {
		return err
	}

	// If cors configuration exists, apply cors middleware
	if ccfg != nil {
		router.Use(cors.New(*ccfg))
	}

	return nil
}

func generateCORSConfiguration(cfg *config.ServerCorsConfig) (*cors.Config, error) {
	// Check if cors configuration exists
	if cfg == nil {
		return nil, nil
	}

	// Prepare configuration
	ccfg := &cors.Config{}
	// Check if should start from default configuration
	if cfg.UseDefaultConfiguration {
		val := cors.DefaultConfig()

		ccfg = &val
	}

	// Check if allow origins exists
	if cfg.AllowOrigins != nil {
		ccfg.AllowOrigins = cfg.AllowOrigins
	}

	// Check if allow methods exists
	if cfg.AllowMethods != nil {
		ccfg.AllowMethods = cfg.AllowMethods
	}

	// Check if allow headers exists
	if cfg.AllowHeaders != nil {
		ccfg.AllowHeaders = cfg.AllowHeaders
	}

	// Check if expose headers exists
	if cfg.ExposeHeaders != nil {
		ccfg.ExposeHeaders = cfg.ExposeHeaders
	}

	// Check if max age duration exists
	if cfg.MaxAgeDuration != "" {
		// Parse max age
		maxAge, err := time.ParseDuration(cfg.MaxAgeDuration)
		// Check error
		if err != nil {
			return nil, errors.WithStack(err)
		}

		ccfg.MaxAge = maxAge
	}

	// Check if allow credentials exists
	if cfg.AllowCredentials != nil {
		ccfg.AllowCredentials = *cfg.AllowCredentials
	}

	// Check if allow wildcard exists
	if cfg.AllowWildcard != nil {
		ccfg.AllowWildcard = *cfg.AllowWildcard
	}

	// Check if allow browser extensions exists
	if cfg.AllowBrowserExtensions != nil {
		ccfg.AllowBrowserExtensions = *cfg.AllowBrowserExtensions
	}

	// Check if allow websockets exists
	if cfg.AllowWebSockets != nil {
		ccfg.AllowWebSockets = *cfg.AllowWebSockets
	}

	// Check if allow files exists
	if cfg.AllowFiles != nil {
		ccfg.AllowFiles = *cfg.AllowFiles
	}

	// Check if allow all origins exists
	if cfg.AllowAllOrigins != nil {
		ccfg.AllowAllOrigins = *cfg.AllowAllOrigins
	}

	// Validate configuration
	err := ccfg.Validate()
	// Check error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return ccfg, nil
}
