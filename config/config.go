package config

import (
	"fmt"
	"net/url"

	"github.com/BurntSushi/toml"
)

const (
	// DefaultTimeout is the default timeout of the route.
	DefaultTimeout int64 = 30
	// DefaultPort is the default port number to listen on.
	DefaultPort string = ":8080"
)

// Route is a struct that represents a route.
type Route struct {
	Path            string `toml:"path"`              // Path is the path of the route. e.g., /api/v1/users
	Backend         string `toml:"backend"`           // Backend is the backend URL of the route. e.g., http://localhost:8080
	Timeout         int64  `toml:"timeout"`           // Timeout is the timeout of the route. e.g., 10
	HealthCheckPath string `toml:"health_check_path"` // HealthCheckPath is the path of the health check. e.g., /health
}

// HealthCheckEnabled returns true if the health check is enabled.
func (r Route) HealthCheckEnabled() bool {
	return r.HealthCheckPath != ""
}

// HealthCheckURL returns the URL of the health check.
func (r Route) HealthCheckURL() (string, error) {
	if !r.HealthCheckEnabled() {
		return "", fmt.Errorf("config: health check is not enabled")
	}
	backendURL, err := url.Parse(r.Backend)
	if err != nil {
		return "", fmt.Errorf("config: failed to parse backend URL for health check: %w", err)
	}
	healthCheckURL, err := backendURL.Parse(r.HealthCheckPath)
	if err != nil {
		return "", fmt.Errorf("config: failed to parse health check path: %w", err)
	}
	return healthCheckURL.String(), nil
}

// Server is a struct that represents a server.
type Server struct {
	Port  string `toml:"port"`  // Port is the port number to listen on.
	Debug bool   `toml:"debug"` // Debug is whether to run in debug mode. By default, only output info/warning/error logs.
}

// Config is a struct that represents a configuration.
type Config struct {
	Server Server  `toml:"server"`
	Routes []Route `toml:"routes"`
}

// NewConfig creates a new Config.
// It reads the configuration file and returns a new Config.
func NewConfig(cfgFilePath string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(cfgFilePath, &cfg); err != nil {
		return nil, fmt.Errorf("config: failed to decode config file: %w", err)
	}

	if cfg.Server.Port == "" {
		cfg.Server.Port = DefaultPort
	}

	for i, route := range cfg.Routes {
		if route.Timeout <= 0 {
			cfg.Routes[i].Timeout = DefaultTimeout
		}
	}
	return &cfg, nil
}
