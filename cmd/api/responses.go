package api

import (
	"encoding/json"
	"net/http"
)

func WriteJsonResp(w http.ResponseWriter, status int, data any, envelopeKey string) {
	out := map[string]any{envelopeKey: data}
	jsonOut, err := json.Marshal(out)
	if err != nil {
		http.Error(w, "error encoding json data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(jsonOut); err != nil {
		http.Error(w, "error writing json data to the response", http.StatusInternalServerError)
	}
}

func WriteJsonError(w http.ResponseWriter, status int, msg string) {
	WriteJsonResp(w, status, msg, "error")
}

func WriteJsonServerError(w http.ResponseWriter) {
	WriteJsonError(w, http.StatusInternalServerError, "internal server error")
}
