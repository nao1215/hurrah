package proxy

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_periodicHealthCheck(t *testing.T) {
	t.Run("health checks are performed periodically", func(t *testing.T) {
		done := make(chan bool, 1) // Create a channel to signal when the test should finish
		var requestCount int

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
			requestCount++
			if requestCount >= 3 {
				done <- true
			}
		}))
		defer server.Close()

		go periodicHealthCheck(server.URL, 2, 100*time.Millisecond)

		select {
		case <-done:
			// Success: 3 health checks were performed
		case <-time.After(1 * time.Second):
			t.Error("Health checks did not execute 3 times within the expected time frame")
		}
	})
}
