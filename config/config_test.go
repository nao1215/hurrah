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
			Routes: []Route{
				{Path: "/service1", Backend: "http://localhost:8081"},
				{Path: "/service2", Backend: "http://localhost:8082"},
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
