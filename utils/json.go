package utils

import (
	"encoding/json"
	"net/http"
)

// SendJSON encode and send json
func SendJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

// ParseJSON Parse from JSON
func ParseJSON(r *http.Request, data interface{}) error {
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(&data)
}
