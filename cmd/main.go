package main

import (
	m "mirror-monitoring/pkg/mirror_status"
	"net/http"
)

func main() {
	http.HandleFunc("/mirrors/status", m.MirrorStatusesHandler)
	http.ListenAndServe(":8080", nil)
}
