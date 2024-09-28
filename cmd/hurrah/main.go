// Package main is the hurrah command entry point.
// The hurrah command is API Gateway for microservices.
package main

import (
	"log/slog"
	"os"

	"github.com/nao1215/hurrah/config"
)

const (
	// ExitCodeOK is the exit code when the program ends successfully.
	ExitCodeOK int = 0
	// ExitCodeError is the exit code when the program ends with an error.
	ExitCodeError int = 1
)

// main is the entry point of the program.
func main() {
	os.Exit(run())
}

// run runs the main logic of the program.
func run() int {
	hurrah, err := newHurrah()
	if err != nil {
		slog.Error("failed to initialize hurrah command", slog.String("error", err.Error()))
		return ExitCodeError
	}

	hurrah.logStartupInfo()
	return ExitCodeOK
}

// hurrah is the main struct of the hurrah command.
type hurrah struct {
	flag config.Flag // flag is the flag at command startup.
}

// newHurrah reads the command line flags and returns a new hurrah.
func newHurrah() (*hurrah, error) { //nolint:unparam
	flag := config.NewFlag()
	// TODO: Implement the configuration file reading process.
	// Use can change log output destination.
	slog.SetDefault(config.NewStructuredLogger(os.Stderr, flag.Debug))
	return &hurrah{
		flag: flag,
	}, nil
}

// logStartupInfo logs the startup information of the hurrah command.
// It's only printed in debug mode.
func (h *hurrah) logStartupInfo() {
	slog.Debug(
		"running condition",
		slog.String("version", config.GetVersion()),
		slog.Int("port", h.flag.Port),
		slog.String("config_file", h.flag.ConfigFile),
		slog.Bool("debug", h.flag.Debug),
	)
}
