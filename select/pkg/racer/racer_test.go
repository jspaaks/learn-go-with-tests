package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var setUpServerAndGetItsUrl = func(delay time.Duration) (server *httptest.Server, url string) {
	delayedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	})
	server = httptest.NewServer(delayedHandler)
	url = server.URL
	return
}

// not sure why but the tests still seem to wait for the slowest-responding server

func TestRacer(t *testing.T) {
	fastServer, fastUrl := setUpServerAndGetItsUrl(2 * time.Millisecond)
	slowServer, slowUrl := setUpServerAndGetItsUrl(20 * time.Millisecond)

	defer fastServer.Close()
	defer slowServer.Close()

	t.Run("select the server with the lowest response time", func(t *testing.T) {
		want := fastUrl
		got, err := Racer(fastUrl, slowUrl)
		if err != nil {
			t.Errorf("Didn't expect an error but got one")
		}
		if got != want {
			t.Errorf("Wanted %q but got %q instead", want, got)
		}
	})
}

func TestConfigurableRacer(t *testing.T) {

	fastServer, fastUrl := setUpServerAndGetItsUrl(10 * time.Millisecond)
	slowServer, slowUrl := setUpServerAndGetItsUrl(100 * time.Millisecond)

	defer fastServer.Close()
	defer slowServer.Close()

	t.Run("verify an error is returned when both servers time out", func(t *testing.T) {

		const want = ""
		got, err := ConfigurableRacer(fastUrl, slowUrl, 5*time.Millisecond)

		if err == nil {
			t.Errorf("Expected an error but didn't get one")
		}
		if got != want {
			t.Errorf("Wanted %q but got %q instead", want, got)
		}
	})
}
