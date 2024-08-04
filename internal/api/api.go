package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	middlewares "github.com/needl3/goreact-template/internal/api/middlwares"
	"github.com/needl3/goreact-template/internal/core"
)

type Middleware func(http.Handler) http.Handler

type api struct {
	authCore    *core.AuthCore
}

func New(ctx context.Context, db *pgxpool.Pool) *api {
	return &api{
		authCore:    core.NewAuthCore(db),
	}
}

func applyMiddlewares(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}

func (a *api) Server(ctx context.Context, port int) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: a.Routers(),
	}
}

func (a *api) Routers() *http.ServeMux {
	mux := http.NewServeMux()

	////////////////////// Rendering section ///////////////////

	// React routes
	// Registered here for authentication check
	registeredPages := map[string]http.Handler{
		"/":            applyMiddlewares(http.HandlerFunc(a.React), middlewares.RedirectToAppIfAuthenticated, middlewares.AuthCheckMiddlware),
		"/app":         applyMiddlewares(http.HandlerFunc(a.React), middlewares.RedirectToRootIfNotAuthenticated, middlewares.AuthCheckMiddlware),
		"/app/new":     applyMiddlewares(http.HandlerFunc(a.React), middlewares.RedirectToRootIfNotAuthenticated, middlewares.AuthCheckMiddlware),
		"/app/history": applyMiddlewares(http.HandlerFunc(a.React), middlewares.RedirectToRootIfNotAuthenticated, middlewares.AuthCheckMiddlware),
		"/app/prompts": applyMiddlewares(http.HandlerFunc(a.React), middlewares.RedirectToRootIfNotAuthenticated, middlewares.AuthCheckMiddlware),
		"/app/{id}":    applyMiddlewares(http.HandlerFunc(a.React), middlewares.RedirectToRootIfNotAuthenticated, middlewares.AuthCheckMiddlware),
	}
	for route, handler := range registeredPages {
		mux.Handle(fmt.Sprintf("GET %s", route), handler)
	}
	// Assets
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("frontend/public/assets")))
	mux.Handle("GET /assets/", fs)

	////////////////////// API Section ///////////////////
	mux.HandleFunc("GET /v1/auth/signin", a.Auth)
	mux.HandleFunc("GET /v1/auth/signout", a.Signout)
	mux.HandleFunc("GET /v1/health", a.Health)
	mux.HandleFunc("GET /v1/auth/callback/google", a.AuthCallback)

	return mux
}
