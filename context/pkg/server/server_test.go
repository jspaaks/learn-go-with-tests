package server

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	contents string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	ch := make(chan string, 1)
	go func() {
		var result string
		// this next for loop serves to emulate a task that is slow, but which can be partitioned
		// into subtasks. Specifically the for loop builds up the string variable 'result'
		// character-by-character such that it is slow enough for potential timeouts to kick in, and
		// such that it can actually check whether it needs to cancel the rest of the work, which it
		// does after each character
		for _, c := range s.contents {
			select {
			case <-ctx.Done():
				log.Println("spy store got called")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		ch <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-ch:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("store contents returned correctly", func(t *testing.T) {
		expected := "hello, world"
		store := &SpyStore{contents: expected, t: t}
		server := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		if response.Body.String() != expected {
			t.Errorf(`got "%s" but wanted "%s"`, response.Body.String(), expected)
		}
	})
	//t.Run("request gets canceled", func(t *testing.T) {
	//	expected := "hello, world"
	//	store := &SpyStore{canceled: false, contents: expected, t: t}
	//	server := Server(store)
	//	request := httptest.NewRequest(http.MethodGet, "/", nil)
	//	cancelingCtx, cancel := context.WithCancel(request.Context())
	//	cancelableRequest := request.WithContext(cancelingCtx)
	//	response := httptest.NewRecorder()
	//	time.AfterFunc(5*time.Millisecond, cancel)
	//	server.ServeHTTP(response, cancelableRequest)
	//	store.assertWasCanceled()
	//})
}
