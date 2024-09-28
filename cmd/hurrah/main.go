package main

import (
	"flag"
	"fmt"
	"os"
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
	flag := newFlag()
	fmt.Printf("Port: %d\n", flag.Port)
	fmt.Printf("ConfigFile: %s\n", flag.ConfigFile)
	fmt.Printf("Debug: %t\n", flag.Debug)

	return ExitCodeOK
}

// Flag represents a flag at command startup.
type Flag struct {
	Port       int    // Port is the port number to listen on.
	ConfigFile string // ConfigFile is the path to the configuration file.
	Debug      bool   // Debug is whether to run in debug mode.
}

// newFlag creates a new Flag.
// It reads the command line flags and returns a new Flag.
func newFlag() *Flag {
	fs := flag.NewFlagSet("", flag.ExitOnError)

	port := fs.Int("port", 8080, "a port number to listen on")
	configFile := fs.String("config", "config.yaml", "a path to the configuration file")
	debugMode := fs.Bool("debug", false, "whether to run in debug mode")

	_ = fs.Parse(os.Args[1:]) // If an error occurs, it will be handled by flag.ExitOnError.

	return &Flag{
		Port:       *port,
		ConfigFile: *configFile,
		Debug:      *debugMode,
	}
}
