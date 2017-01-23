package main

import (
	"fmt"
	"net/http"
	"os"
	"encoding/json"
)

type Health struct {
	Version string `json:"version" form:"version" binding:"required"`
	Status  string `json:"status" form:"status" binding:"required"`
}

func getHostname(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "The hostname is %s!", hostname)
}

func getHealthCheck(w http.ResponseWriter, r *http.Request) {
	healthCheck := Health{os.Getenv("APP_VERSION"), "OK"}

	response, err := json.Marshal(healthCheck)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func main() {
	http.HandleFunc("/", getHostname)
	http.HandleFunc("/health", getHealthCheck)
	http.ListenAndServe(":8080", nil)
}
