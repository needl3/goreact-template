package api

import (
	"net/http"
)

func (a *api) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
