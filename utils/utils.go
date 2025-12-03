package utils

import "strings"

func ProcessWSURL(url string) string {

	// force ws:// schema
	if strings.HasPrefix(url, "http://") {
		url = "ws://" + strings.TrimPrefix(url, "http://")
	}
	if strings.HasPrefix(url, "https://") {
		url = "wss://" + strings.TrimPrefix(url, "https://")
	}
	return url
}

func CheckIfJSON(data []byte) bool {
	return true
}
