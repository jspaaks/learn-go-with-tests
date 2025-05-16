package assertions

import "testing"

func AssertEqual[T comparable](t *testing.T, got T, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Error("Expected false but got true")
	}
}

func AssertNotEqual[T comparable](t *testing.T, got T, want T) {
	t.Helper()
	if got == want {
		t.Errorf("got %+v, but didn't want it", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Error("Expected true but got false")
	}
}
