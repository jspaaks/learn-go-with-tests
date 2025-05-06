package racer

import (
	"fmt"
	"net/http"
	"time"
)

func ConfigurableRacer(url0 string, url1 string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(url0):
		return url0, nil
	case <-ping(url1):
		return url1, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("Timed out waiting for %q and %q", url0, url1)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func Racer(url0 string, url1 string) (winner string, err error) {
	const timeout time.Duration = 10 * time.Second
	return ConfigurableRacer(url0, url1, timeout)
}
