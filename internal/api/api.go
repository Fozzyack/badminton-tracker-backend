package api

import (
	"encoding/json"
	"net/http"
)

func DecodeJSON(r *http.Request, out interface{}) error {
	return json.NewDecoder(r.Body).Decode(out)
}

func SendJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func SendError(w http.ResponseWriter, errorMsg string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	SendJSON(w, map[string]string{"error": errorMsg})
}
