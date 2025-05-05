package racer

import (
	"net/http"
	"time"
)

func Racer(url0 string, url1 string) (winner string) {
	measureResponseDuration := func(url string) time.Duration {
		start := time.Now()
		http.Get(url)
		return time.Since(start)
	}
	duration0 := measureResponseDuration(url0)
	duration1 := measureResponseDuration(url1)
	if duration0 < duration1 {
		return url0
	}
	return url1
}
