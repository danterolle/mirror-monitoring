package mirror_status

import "net/http"

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
