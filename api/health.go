package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type health struct {
	Status string    `json:"status"`
	Time   time.Time `json:"time"`
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health{Status: "ok", Time: time.Now().UTC()})
}
func errorJSON(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
