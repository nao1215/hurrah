package config

import (
	"flag"
	"os"
)

// Flag represents a flag at command startup.
type Flag struct {
	Port       int    // Port is the port number to listen on.
	ConfigFile string // ConfigFile is the path to the configuration file.
	Debug      bool   // Debug is whether to run in debug mode.
}

// NewFlag creates a new Flag.
// It reads the command line flags and returns a new Flag.
func NewFlag() Flag {
	fs := flag.NewFlagSet("", flag.ExitOnError)

	port := fs.Int("port", 8080, "a port number to listen on")
	configFile := fs.String("config", "config.yaml", "a path to the configuration file")
	debugMode := fs.Bool("debug", false, "whether to run in debug mode. By default, only output warning/error logs")

	_ = fs.Parse(os.Args[1:]) // If an error occurs, it will be handled by flag.ExitOnError.

	return Flag{
		Port:       *port,
		ConfigFile: *configFile,
		Debug:      *debugMode,
	}
}
