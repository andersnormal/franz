package config

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"syscall"

	log "github.com/sirupsen/logrus"
)

// Config contains a configuration for franz
type Config struct {
	// Verbose toggles the verbosity
	Verbose bool

	// LogLevel is the level with with to log for this config
	LogLevel log.Level

	// ReloadSignal
	ReloadSignal syscall.Signal

	// TermSignal
	TermSignal syscall.Signal

	// KillSignal
	KillSignal syscall.Signal

	// Timeout of the runtime
	Timeout time.Duration

	// StatusAddr is the addr of the debug listener
	StatusAddr string

	// Addr is the address to listen on
	Addr string

	// Debug ...
	Debug bool

	// GRPCAddr ...
	GRPCAddr string
}

const (
	// DefaultLogLevel is the default logging level.
	DefaultLogLevel = log.WarnLevel

	// DefaultTermSignal is the signal to term the agent.
	DefaultTermSignal = syscall.SIGTERM

	// DefaultReloadSignal is the default signal for reload.
	DefaultReloadSignal = syscall.SIGHUP

	// DefaultKillSignal is the default signal for termination.
	DefaultKillSignal = syscall.SIGINT

	// DefaultVerbose is the default verbosity.
	DefaultVerbose = false

	// DefaultStatusAddr is the default addrs for debug listener
	DefaultStatusAddr = ":8443"

	// DefaultAddr is the default addrs to listen on
	DefaultAddr = ":443"

	// DefaultDebug is the default debug status.
	DefaultDebug = false
)

// New returns a new Config
func New() *Config {
	return &Config{
		Verbose:      DefaultVerbose,
		LogLevel:     DefaultLogLevel,
		ReloadSignal: DefaultReloadSignal,
		TermSignal:   DefaultTermSignal,
		KillSignal:   DefaultKillSignal,
		StatusAddr:   DefaultStatusAddr,
		Debug:        DefaultDebug,
		Addr:         DefaultAddr,
		GRPCAddr:     DefaultGRPCAddr,
	}
}
