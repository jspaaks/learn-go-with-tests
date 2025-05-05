package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	getServerAndUrl := func(delay time.Duration) (server *httptest.Server, url string){
		delayedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		})
		server = httptest.NewServer(delayedHandler)
		url = server.URL
		return
	}

	fastServer, fastUrl := getServerAndUrl(2 * time.Millisecond)
	defer fastServer.Close()

	slowServer, slowUrl := getServerAndUrl(20 * time.Millisecond)
	defer slowServer.Close()

	want := fastUrl
	got := Racer(fastUrl, slowUrl)
	if got != want {
		t.Errorf("Wanted %q but got %q instead", want, got)
	}
}
