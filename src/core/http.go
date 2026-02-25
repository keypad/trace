package core

import (
	"net/http"
	"time"
)

func httpcheck(target string, timeout time.Duration) httprow {
	client := &http.Client{Timeout: timeout}
	start := time.Now()
	response, err := client.Get(target)
	latency := time.Since(start)
	if err != nil {
		return httprow{target: target, state: "down", code: 0, latency: latency, size: 0, note: err.Error()}
	}
	defer response.Body.Close()

	return httprow{target: target, state: "up", code: response.StatusCode, latency: latency, size: response.ContentLength, note: ""}
}
