// Package proxy provides a way to set the proxy server settings.
package proxy

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/nao1215/hurrah/app/middleware"
	"github.com/nao1215/hurrah/config"
)

// SetProxy sets the proxy server settings.
func SetProxy(mux *http.ServeMux, routes []config.Route, middlewares ...middleware.Middleware) error {
	for _, route := range routes {
		proxy, err := newReverseProxy(route.Backend, route.Timeout)
		if err != nil {
			return fmt.Errorf("proxy: failed to create a reverse proxy for route %s: %w", route.Path, err)
		}
		if route.HealthCheckEnabled() {
			u, err := route.HealthCheckURL()
			if err != nil {
				return fmt.Errorf("proxy: failed to get health check URL for route %s: %w", route.Path, err)
			}
			ctx := context.Background()
			go periodicHealthCheck(ctx, u, route.Timeout, 1*time.Second)
		}

		handlerWithMiddleware := middleware.Chain(middleware.ToHandlerWithCtx(proxy), middlewares...)
		mux.Handle(route.Path, handlerWithMiddleware.AdaptHandler())
		slog.Debug("proxy: set a reverse proxy", slog.String("path", route.Path), slog.String("backend", route.Backend))
	}
	return nil
}

// newReverseProxy creates a reverse proxy to the given backend URL
func newReverseProxy(target string, timeout int64) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(target)
	if err != nil {
		return nil, fmt.Errorf("failed to parse target URL: %w", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(url)

	proxy.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout: time.Duration(timeout) * time.Second,
		}).DialContext,
		ResponseHeaderTimeout: time.Duration(timeout) * time.Second,
		TLSHandshakeTimeout:   time.Duration(timeout) * time.Second,
	}
	return proxy, nil
}
