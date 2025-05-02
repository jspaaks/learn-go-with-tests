package greeter

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")
	got := buffer.String()
	want := "Hello, Chris\n"
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
