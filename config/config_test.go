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
					Path:            "/service1",
					Backend:         "http://localhost:8081",
					Timeout:         10,
					HealthCheckPath: "/health",
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

func TestRoute_HealthCheckEnabled(t *testing.T) {
	t.Parallel()

	type fields struct {
		HealthCheckPath string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Health check enabled",
			fields: fields{
				HealthCheckPath: "/health",
			},
			want: true,
		},
		{
			name: "Health check disabled",
			fields: fields{
				HealthCheckPath: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := Route{
				HealthCheckPath: tt.fields.HealthCheckPath,
			}
			if got := r.HealthCheckEnabled(); got != tt.want {
				t.Errorf("Route.HealthCheckEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoute_HealthCheckURL(t *testing.T) {
	t.Parallel()

	type fields struct {
		Backend         string
		HealthCheckPath string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "get health check URL",
			fields: fields{
				Backend:         "http://localhost:8080",
				HealthCheckPath: "/health",
			},
			want:    "http://localhost:8080/health",
			wantErr: false,
		},
		{
			name: "health check is not enabled",
			fields: fields{
				Backend:         "http://localhost:8080",
				HealthCheckPath: "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "failed to parse backend URL",
			fields: fields{
				Backend:         ":",
				HealthCheckPath: "/health",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "failed to parse health check path",
			fields: fields{
				Backend:         "http://localhost:8080",
				HealthCheckPath: ":",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := Route{
				Backend:         tt.fields.Backend,
				HealthCheckPath: tt.fields.HealthCheckPath,
			}
			got, err := r.HealthCheckURL()
			if (err != nil) != tt.wantErr {
				t.Errorf("Route.HealthCheckURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Route.HealthCheckURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
