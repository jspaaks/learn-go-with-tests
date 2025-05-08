package server

import (
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data, _ := store.Fetch(request.Context())
		fmt.Fprint(writer, data)
	}
}
