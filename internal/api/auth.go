package api

import (
	"fmt"
	"net/http"
)

func (a *api) Auth(w http.ResponseWriter, r *http.Request) {
	url, err := a.authCore.GenerateAuthUrl()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}

func (a *api) AuthCallback(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	code := queries.Get("code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accessToken, err := a.authCore.Signup(r.Context(), code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{
		Name:     "Authorization",
		Value:    fmt.Sprintf("Bearer %s", accessToken),
		HttpOnly: true,
		Path:     "/",
		MaxAge:   3600 * 14,
	}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/app", http.StatusFound)
}

func (a *api) Signout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1,
	})
	http.Redirect(w, r, "/", http.StatusFound)
}
