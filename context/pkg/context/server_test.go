package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	canceled bool
	contents string
	t        *testing.T
}

func (s *SpyStore) assertWasCanceled() {
	s.t.Helper()
	if !s.canceled {
		s.t.Errorf("request should have been canceled")
	}
}

func (s *SpyStore) assertWasNotCanceled() {
	s.t.Helper()
	if s.canceled {
		s.t.Errorf("request got unexpectedly canceled")
	}
}

func (s *SpyStore) Cancel() {
	s.canceled = true
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.contents
}

func TestServer(t *testing.T) {
	t.Run("store contents returned correctly", func(t *testing.T) {
		expected := "hello, world"
		store := &SpyStore{canceled: false, contents: expected, t: t}
		server := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		store.assertWasNotCanceled()
		if response.Body.String() != expected {
			t.Errorf(`got "%s" but wanted "%s"`, response.Body.String(), expected)
		}
	})
	t.Run("request gets canceled", func(t *testing.T) {
		expected := "hello, world"
		store := &SpyStore{canceled: false, contents: expected, t: t}
		server := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancelingCtx, cancel := context.WithCancel(request.Context())
		cancelableRequest := request.WithContext(cancelingCtx)
		response := httptest.NewRecorder()
		time.AfterFunc(5*time.Millisecond, cancel)
		server.ServeHTTP(response, cancelableRequest)
		store.assertWasCanceled()
	})
}
