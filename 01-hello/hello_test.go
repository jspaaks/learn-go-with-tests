package main

import "testing"

func TestSayHello(t *testing.T) {
	got := SayHello("Chris")
	want := "Hello, Chris"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
