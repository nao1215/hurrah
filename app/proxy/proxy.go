// Package proxy
package proxy

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/nao1215/hurrah/config"
)

// SetProxy sets the proxy server settings.
func SetProxy(mux *http.ServeMux, routes []config.Route) error {
	for _, route := range routes {
		proxy, err := newReverseProxy(route.Backend)
		if err != nil {
			return fmt.Errorf("proxy: failed to create a reverse proxy for route %s: %w", route.Path, err)
		}
		mux.Handle(route.Path, proxy)
		slog.Debug("proxy: set a reverse proxy", slog.String("path", route.Path), slog.String("backend", route.Backend))
	}
	return nil
}

// newReverseProxy creates a reverse proxy to the given backend URL
func newReverseProxy(target string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(target)
	if err != nil {
		return nil, fmt.Errorf("failed to parse target URL: %w", err)
	}
	return httputil.NewSingleHostReverseProxy(url), nil
}
