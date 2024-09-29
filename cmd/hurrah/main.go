// Package main is the hurrah command entry point.
// The hurrah command is API Gateway for microservices.
package main

import (
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/nao1215/hurrah/app/proxy"
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
	if err := hurrah.run(); err != nil {
		slog.Error("failed to run hurrah command", slog.String("error", err.Error()))
		return ExitCodeError
	}
	return ExitCodeOK
}

// hurrah is the main struct of the hurrah command.
type hurrah struct {
	flag   *config.Flag   // flag is the flag at command startup.
	config *config.Config // config is the configuration of the hurrah command.
	mux    *http.ServeMux // mux is the HTTP request multiplexer.
}

// newHurrah reads the command line flags and returns a new hurrah.
func newHurrah() (*hurrah, error) { //nolint:unparam
	flag := config.NewFlag()
	slog.SetDefault(config.NewStructuredLogger(os.Stderr, flag.Debug))

	cfg, err := config.NewConfig(flag.ConfigFile)
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	if err := proxy.SetProxy(mux, cfg.Routes); err != nil {
		return nil, err
	}

	return &hurrah{
		flag:   flag,
		config: cfg,
		mux:    mux,
	}, nil
}

// run runs the main logic of the hurrah command.
func (h *hurrah) run() error {
	h.logStartupInfo()
	return http.ListenAndServe(h.port(), h.mux)
}

// port returns the port number to listen on.
func (h *hurrah) port() string {
	if !strings.HasPrefix(h.flag.Port, ":") {
		return ":" + h.flag.Port
	}
	return h.flag.Port
}

// logStartupInfo logs the startup information of the hurrah command.
// It's only printed in debug mode.
func (h *hurrah) logStartupInfo() {
	routing := ""
	for _, route := range h.config.Routes {
		routing += route.Path + " -> " + route.Backend + " "
	}

	slog.Debug(
		"running condition",
		slog.String("version", config.GetVersion()),
		slog.String("port", h.flag.Port),
		slog.String("config_file", h.flag.ConfigFile),
		slog.Bool("debug", h.flag.Debug),
		slog.String("routing", routing),
	)
}
