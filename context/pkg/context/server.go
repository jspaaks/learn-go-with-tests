package context

import (
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// do the fetching in a channel...
		ch := make(chan string, 1)
		go func() {
			ch <- store.Fetch()
		}()

		// ...so I can simultaneously look for cancelation in
		// another channel (request.Context().Done())
		select {
		case data := <-ch:
			fmt.Fprint(writer, data)
		case <-request.Context().Done():
			store.Cancel()
		}
	}
}
