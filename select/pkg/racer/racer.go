package racer

import "net/http"

func Racer(url0 string, url1 string) (winner string) {
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
		return url0
	case <-ping(url1):
		return url1
	}
}
