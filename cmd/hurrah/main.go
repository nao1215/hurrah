// Package main is the hurrah command entry point.
// The hurrah command is API Gateway for microservices.
package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

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
func newHurrah() (*hurrah, error) {
	flag := config.NewFlag()
	cfg, err := config.NewConfig(flag.ConfigFile)
	if err != nil {
		return nil, err
	}
	slog.SetDefault(config.NewStructuredLogger(os.Stderr, flag.Debug || cfg.Server.Debug))

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

	server := &http.Server{
		Addr:              h.port(),
		Handler:           h.mux,
		ReadHeaderTimeout: time.Duration(10) * time.Second, // TODO: Use can be configured.
	}
	slog.Info("starting the server", slog.String("address", server.Addr))
	return server.ListenAndServe()
}

// port returns the port number to listen on.
func (h *hurrah) port() string {
	if h.config.Server.Port != "" {
		if !strings.HasPrefix(h.config.Server.Port, ":") {
			return ":" + h.config.Server.Port
		}
		return h.config.Server.Port
	}

	if h.flag.Port == "" {
		return config.DefaultPort
	}
	if !strings.HasPrefix(h.flag.Port, ":") {
		return ":" + h.flag.Port
	}
	return h.flag.Port
}

// logStartupInfo logs the startup information of the hurrah command.
// It's only printed in debug mode.
func (h *hurrah) logStartupInfo() {
	var builder strings.Builder
	for _, route := range h.config.Routes {
		builder.WriteString(route.Path)
		builder.WriteString(" -> ")
		builder.WriteString(route.Backend)
		builder.WriteString(fmt.Sprintf(" (timeout: %d[s])", route.Timeout))
		builder.WriteString(" ")
	}
	routing := builder.String()

	slog.Debug(
		"running condition",
		slog.String("version", config.GetVersion()),
		slog.String("port", h.flag.Port),
		slog.String("config_file", h.flag.ConfigFile),
		slog.Bool("debug", h.flag.Debug),
		slog.String("routing", routing),
	)
}
