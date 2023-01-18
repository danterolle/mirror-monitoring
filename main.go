package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/mirrors/status", mirrorStatusesHandler)
	http.ListenAndServe(":8080", nil)
}
