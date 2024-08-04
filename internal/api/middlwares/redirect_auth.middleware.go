package middlewares

import (
	"net/http"
)

func RedirectToAppIfAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.PathValue("userId") != "" {
			http.Redirect(w, r, "/app", http.StatusSeeOther)
		}
		next.ServeHTTP(w, r)
	})
}

func RedirectToRootIfNotAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.PathValue("userId") == "" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		next.ServeHTTP(w, r)
	})
}
