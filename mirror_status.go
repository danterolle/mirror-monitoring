package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type MirrorStatus struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

// checkMirror function takes a URL as a parameter and uses the `http.Head` function to check the status of the mirror.
// It returns the status of the mirror as a string, "online", "offline" or "unknown"
func checkMirror(url string) (string, error) {
	resp, err := http.Head(url)
	if err != nil {
		return "unknown", err
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return "online", nil
	} else {
		return "offline", nil
	}
}

// mirrorStatusesHandler function reads the "mirrors.json" file, checks the status of each mirror, and returns the status of each mirror as a JSON response
func mirrorStatusesHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("mirrors.json")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading mirrors file: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var mirrors []map[string]interface{}
	if err := json.NewDecoder(file).Decode(&mirrors); err != nil && err != io.EOF {
		http.Error(w, fmt.Sprintf("Error decoding mirrors file: %v", err), http.StatusInternalServerError)
		return
	}

	var mirrorStatuses []MirrorStatus
	for _, mirror := range mirrors {
		url := mirror["url"].(string)
		status, _ := checkMirror(url)
		mirrorStatuses = append(mirrorStatuses, MirrorStatus{URL: url, Status: status})
	}

	mirrorStatusesJSON, err := json.MarshalIndent(mirrorStatuses, "", "    ")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshalling mirror statuses: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(mirrorStatusesJSON)
}
