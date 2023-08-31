package main

import (
	"os"

	"github.com/pingidentity/pingctl/cmd"
	"github.com/pingidentity/pingctl/internal/config"
	"github.com/pingidentity/pingctl/internal/logger"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"

	// goreleaser can also pass the specific commit if you want
	commit string = ""
)

func main() {
	l := logger.Get()

	l.Debug().Msg("Starting pingctl")

	cmd.Execute()
}

func init() {
	l := logger.Get()

	if err := config.Init(); err != nil {
		l.Error().Err(err).Msg("Error initializing configuration")
		os.Exit(1)
	}
}
