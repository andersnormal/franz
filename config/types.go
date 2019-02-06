package config

import (
	log "github.com/sirupsen/logrus"
)

// Config ...
type Config struct {
	// LogFormat
	LogFormat string

	// LogLevel
	LogLevel log.Level

	// Verbose
	Verbose bool
}
