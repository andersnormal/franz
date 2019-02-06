package config

import (
	log "github.com/sirupsen/logrus"
)

// SetupLogger prepares the logger instance
func (c *Config) SetupLogger() {
	switch c.LogFormat {
	case "text":
		log.SetFormatter(&log.TextFormatter{})
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.JSONFormatter{})
	}

	// Only log the warning severity or above.
	log.SetLevel(c.LogLevel)

	// if we should output verbose
	if c.Verbose {
		log.SetLevel(log.InfoLevel)
	}
}
