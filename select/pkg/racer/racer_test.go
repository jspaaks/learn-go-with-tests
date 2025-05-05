package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var getServerAndUrl = func(delay time.Duration) (server *httptest.Server, url string){
	delayedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	})
	server = httptest.NewServer(delayedHandler)
	url = server.URL
	return
}


func TestRacer(t *testing.T) {

	fastServer, fastUrl := getServerAndUrl(2 * time.Millisecond)
	defer fastServer.Close()

	slowServer, slowUrl := getServerAndUrl(20 * time.Millisecond)
	defer slowServer.Close()

	t.Run("selecting the url whose server responds fastest", func(t *testing.T){
		want := fastUrl
		got, _ := Racer(fastUrl, slowUrl, 10 * time.Second)
		if got != want {
			t.Errorf("Wanted %q but got %q instead", want, got)
		}
	})

	t.Run("timing out after if neither server has responded after 10 seconds", func(t *testing.T){
		fastServer, fastUrl := getServerAndUrl(11 * time.Second)
		defer fastServer.Close()

		slowServer, slowUrl := getServerAndUrl(12 * time.Second)
		defer slowServer.Close()

		_, err := Racer(fastUrl, slowUrl, 10 * time.Second)
		if err == nil {
			t.Errorf("Expected an error but didn't get one")
		}
	})

}
