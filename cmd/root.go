package cmd

import (
	"math/rand"
	"time"

	"github.com/andersnormal/franz/config"
	"github.com/andersnormal/franz/version"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cfg *config.Config
)

var root = &cobra.Command{
	Use:     "franz",
	Version: version.Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	// seed
	rand.Seed(time.Now().UnixNano())

	// init config
	// create config
	cfg = config.New()

	// add flags
	cfg.AddFlags(root)

	// set default formatter
	log.SetFormatter(&log.TextFormatter{})

	// silence on the root cmd
	root.SilenceErrors = true
	root.SilenceUsage = true

	// initialize upon running commands
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// setup logger
	cfg.SetupLogger()
}

func Execute() {
	if err := root.Execute(); err != nil {
		log.Error(err)
	}
}
