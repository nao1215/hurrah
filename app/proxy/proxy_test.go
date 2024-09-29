// Package proxy
package proxy

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nao1215/hurrah/config"
)

func TestSetProxy(t *testing.T) {

	t.Run("SetProxy with valid routes", func(t *testing.T) {
		// Create two test backend server
		backendServer1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("backend 1"))
		}))
		defer backendServer1.Close()

		backendServer2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("backend 2"))
		}))
		defer backendServer2.Close()

		// routing configuration
		routes := []config.Route{
			{
				Path:    "/service1",
				Backend: backendServer1.URL,
			},
			{
				Path:    "/service2",
				Backend: backendServer2.URL,
			},
		}

		mux := http.NewServeMux()
		err := SetProxy(mux, routes)
		if err != nil {
			t.Errorf("SetProxy() error = %v", err)
		}

		// run proxy server
		testServer := httptest.NewServer(mux)
		defer testServer.Close()

		// request to /service1 and check the response
		resp1, err := http.Get(testServer.URL + "/service1")
		if err != nil {
			t.Errorf("http.Get() error = %v", err)
		}
		defer resp1.Body.Close()

		body1 := make([]byte, 128)
		n1, err := resp1.Body.Read(body1)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				t.Errorf("resp1.Body.Read() error = %v", err)
			}
		}
		if diff := cmp.Diff("backend 1", string(body1[:n1])); diff != "" {
			t.Errorf("resp1.Body.Read() mismatch (-got +want):\n%s", diff)
		}
		if diff := cmp.Diff(http.StatusOK, resp1.StatusCode); diff != "" {
			t.Errorf("resp1.StatusCode mismatch (-got +want):\n%s", diff)
		}

		// request to /service2 and check the response
		resp2, err := http.Get(testServer.URL + "/service2")
		if err != nil {
			t.Errorf("http.Get() error = %v", err)
		}
		defer resp2.Body.Close()

		body2 := make([]byte, 128)
		n2, err := resp2.Body.Read(body2)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				t.Errorf("resp1.Body.Read() error = %v", err)
			}
		}
		if diff := cmp.Diff("backend 2", string(body2[:n2])); diff != "" {
			t.Errorf("resp2.Body.Read() mismatch (-got +want):\n%s", diff)
		}
		if diff := cmp.Diff(http.StatusOK, resp2.StatusCode); diff != "" {
			t.Errorf("resp2.StatusCode mismatch (-got +want):\n%s", diff)
		}
	})

	t.Run("SetProxy with invalid backend URL", func(t *testing.T) {
		routes := []config.Route{
			{
				Path:    "/service1",
				Backend: "postgres://user:abc{DEf1=ghi@example.com:5432/bad",
			},
		}
		mux := http.NewServeMux()
		if err := SetProxy(mux, routes); err == nil {
			t.Error("SetProxy() error = nil, want error")
		}
	})

	t.Run("SetProxy without route settings", func(t *testing.T) {
		mux := http.NewServeMux()
		if err := SetProxy(mux, nil); err != nil {
			t.Errorf("SetProxy() error = %v", err)
		}
	})
}
