package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//respond encodes `value` into JSON and writes that to the response
func respond(w http.ResponseWriter, value interface{}) {
	w.Header().Add(headerContentType, contentTypeJSON)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response value to JSON: %v", err),
			http.StatusInternalServerError)
	}
}
