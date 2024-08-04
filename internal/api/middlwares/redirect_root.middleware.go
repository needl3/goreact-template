package middlewares

import (
	"net/http"
)

// Used to redirect to home if not authenticated
func RedirectIfAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.PathValue("userId") != "" {
			http.Redirect(w, r, "/app", http.StatusSeeOther)
		}
	})
}
