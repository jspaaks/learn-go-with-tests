package server

import (
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data, err := store.Fetch(request.Context())
		if err != nil {
			return
		}
		fmt.Fprint(writer, data)
	}
}
