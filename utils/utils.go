package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// parses the request body and stores in the variable of type payload
func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("Missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

//dedicated response object handler
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

//dedicated error object handler
func WriteError (w http.ResponseWriter, status int, err error){
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}