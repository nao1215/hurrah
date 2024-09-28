package config

import (
	"runtime/debug"
	"testing"
)

func TestGetVersion(t *testing.T) {
	t.Run("Get version from ldflags", func(t *testing.T) {
		Version = "1.2.3"
		expected := "1.2.3"
		result := GetVersion()

		if result != expected {
			t.Errorf("Expected %s, but got %s", expected, result)
		}
	})

	t.Run("Get version from buildInfo", func(t *testing.T) {
		Version = ""
		buildInfo, ok := debug.ReadBuildInfo()
		if !ok {
			t.Fatalf("Failed to read build info")
		}

		expected := buildInfo.Main.Version
		result := GetVersion()
		if result != expected {
			t.Errorf("Expected %s, but got %s", expected, result)
		}
	})
}
