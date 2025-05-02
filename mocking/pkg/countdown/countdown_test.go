package countdown

import (
	"bytes"
	"strings"
	"testing"
)

func TestCountdown(t *testing.T) {

	t.Run("verify the output is correct", func(t *testing.T) {

		buffer := bytes.Buffer{}
		sleeper := TimedSleeper{Duration: 0}
		Countdown(&buffer, &sleeper)

		got := buffer.String()
		want := "3\n2\n1\nGo!\n"
		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	})

	t.Run("verify the write calls have sleeps in between", func(t *testing.T) {

		spy := countdownSpy{}
		Countdown(&spy, &spy)
		got := strings.Join(spy.Calls, "\n")
		want := "write\nsleep\nwrite\nsleep\nwrite\nsleep\nwrite"
		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	})

}


type countdownSpy struct {
	Calls []string
}

func (c *countdownSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, "write")
	return
}

func (c *countdownSpy) Sleep() {
	c.Calls = append(c.Calls, "sleep")
}
