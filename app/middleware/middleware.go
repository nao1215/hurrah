// Package middleware provides a way to chain multiple middlewares.
package middleware

import (
	"context"
	"net/http"

	"log/slog"
)

type (
	// HandlerWithCtx type defines a function that handles an HTTP request with a context.
	HandlerWithCtx func(context.Context, http.ResponseWriter, *http.Request) error
	// Middleware type defines a function that wraps an HTTP handler.
	Middleware func(HandlerWithCtx) HandlerWithCtx
)

// Chain creates a new Middleware by chaining the middlewares.
// The returned Middleware executes the middlewares in the order they are passed.
func Chain(handler HandlerWithCtx, middlewares ...Middleware) HandlerWithCtx {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

// ToHandlerWithCtx converts an http.Handler to a HandlerWithCtx.
func ToHandlerWithCtx(handler http.Handler) HandlerWithCtx {
	return func(_ context.Context, w http.ResponseWriter, r *http.Request) error {
		handler.ServeHTTP(w, r)
		return nil
	}
}

// AdaptHandler converts a HandlerWithCtx to an http.Handler.
// The returned http.Handler calls the HandlerWithCtx with the request context.
func (h HandlerWithCtx) AdaptHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if err := h(ctx, w, r); err != nil {
			slog.Error("middleware: failed to handle the request", slog.String("error", err.Error()))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})
}

// Kind represents the kind of the middleware.
type Kind string

const (
	// KindBasicAuth is a middleware that checks the basic authentication.
	KindBasicAuth Kind = "basic_auth"
)

// BasicAuth is a middleware that checks the basic authentication.
func BasicAuth() Middleware {
	return func(next HandlerWithCtx) HandlerWithCtx {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			// TODO: implement basic authentication
			return next(ctx, w, r)
		}
	}
}
