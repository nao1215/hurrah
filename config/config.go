package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// Route is a struct that represents a route.
type Route struct {
	Path    string `toml:"path"`    // Path is the path of the route. e.g., /api/v1/users
	Backend string `toml:"backend"` // Backend is the backend URL of the route. e.g., http://localhost:8080
}

// Config is a struct that represents a configuration.
type Config struct {
	Routes []Route `toml:"routes"`
}

// NewConfig creates a new Config.
// It reads the configuration file and returns a new Config.
func NewConfig(cfgFilePath string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(cfgFilePath, &cfg); err != nil {
		return nil, fmt.Errorf("config: failed to decode config file: %w", err)
	}
	return &cfg, nil
}
