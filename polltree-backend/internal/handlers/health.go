package handlers

import (
	"encoding/json"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Server is healthy")

	response := map[string]string {
		"status": "ok",
		"message": "server is healthy",
	}

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(response)
}