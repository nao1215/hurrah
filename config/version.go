// Package config provides configuration for the hurrah command.
package config

import (
	"fmt"
	"runtime/debug"
)

// Version value is set by ldflags
var Version string

// Name is command name
const Name = "hurrah"

// GetVersion return bba command version.
// Version global variable is set by ldflags.
func GetVersion() string {
	version := "unknown"
	if Version != "" {
		version = Version
	} else if buildInfo, ok := debug.ReadBuildInfo(); ok {
		version = buildInfo.Main.Version
	}
	return fmt.Sprintf("%s version %s", Name, version)
}
