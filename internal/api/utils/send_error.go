package utils

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func SendError(w http.ResponseWriter, err error, responseType int) {
	w.WriteHeader(responseType)
	json.NewEncoder(w).Encode(Error{err.Error()})
}
