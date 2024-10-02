package config

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewConfig(t *testing.T) {
	t.Run("Read config file", func(t *testing.T) {
		got, err := NewConfig(filepath.Join("testdata", "config.toml"))
		if err != nil {
			t.Errorf("NewConfig() error = %v", err)
		}

		want := &Config{
			Server: Server{
				Port:  "9191",
				Debug: true,
			},
			Routes: []Route{
				{
					Path:    "/service1",
					Backend: "http://localhost:8081",
					Timeout: 10,
				},
				{
					Path:    "/service2",
					Backend: "http://localhost:8082",
					Timeout: 30,
				},
			},
		}

		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("NewConfig() mismatch (-got +want):\n%s", diff)
		}
	})

	t.Run("Read config file with default values", func(t *testing.T) {
		got, err := NewConfig(filepath.Join("testdata", "default.toml"))
		if err != nil {
			t.Errorf("NewConfig() error = %v", err)
		}

		want := &Config{
			Server: Server{
				Port:  DefaultPort,
				Debug: false,
			},
			Routes: []Route{
				{
					Path:    "/service1",
					Backend: "http://localhost:8081",
					Timeout: 30,
				},
			},
		}

		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("NewConfig() mismatch (-got +want):\n%s", diff)
		}
	})

	t.Run("Read config file that not exist", func(t *testing.T) {
		_, err := NewConfig(filepath.Join("testdata", "not-exist.toml"))
		if err == nil {
			t.Error("NewConfig() error = nil, want error")
		}
	})

	t.Run("Read invalid config file", func(t *testing.T) {
		_, err := NewConfig(filepath.Join("testdata", "invalid.toml"))
		if err == nil {
			t.Error("NewConfig() error = nil, want error")
		}
	})
}
