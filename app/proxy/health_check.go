package proxy

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

// periodicHealthCheck checks the backend health every specified interval.
// TODO: configuable interval
// TODO: metric for health check
func periodicHealthCheck(ctx context.Context, backend string, timeout int64, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		if ctx.Err() != nil {
			slog.Info("proxy: health check stopped", slog.String("backend", backend))
			return
		}

		func() {
			client := http.Client{
				Timeout: time.Duration(timeout) * time.Second,
			}
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, backend, nil)
			if err != nil {
				slog.Error("proxy: failed to create a health check request", slog.String("backend", backend), slog.String("error", err.Error()))
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				slog.Error("proxy: periodic health check failed", slog.String("backend", backend), slog.String("error", err.Error()))
				return
			}
			defer resp.Body.Close() //nolint:errcheck

			if resp.StatusCode == http.StatusOK {
				slog.Info("proxy: backend is healthy", slog.String("backend", backend))
			} else {
				slog.Error("proxy: backend health check failed", slog.String("backend", backend), slog.Int("status", resp.StatusCode))
			}
		}()
	}
}
