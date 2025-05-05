package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(url0 string, url1 string, timeout time.Duration) (winner string, err error) {
	var ping = func(url string) chan struct{} {
		ch := make(chan struct{})
		go func (){
			http.Get(url)
			close(ch)
		}()
		return ch
	}

	select {
	case <-ping(url0):
		return url0, nil
	case <-ping(url1):
		return url1, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("Timed out waiting for %q and %q", url0, url1)
	}
}
