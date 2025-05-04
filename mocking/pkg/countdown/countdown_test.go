package countdown

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

type callSequenceSpy struct {
	Calls []string
}

func (c *callSequenceSpy) Sleep() {
	c.Calls = append(c.Calls, "sleep")
}

func (c *callSequenceSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, "write")
	return
}

func TestCountdown(t *testing.T) {
	t.Run("verify the output is correct", func(t *testing.T) {
		buffer := bytes.Buffer{}
		sleeper := TimedSleeper{Duration: 0, SleepFun: time.Sleep}
		Countdown(&buffer, &sleeper)
		got := buffer.String()
		want := "3\n2\n1\nGo!\n"
		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	})

	t.Run("verify the write calls have sleeps in between", func(t *testing.T) {
		// I'd prefer to put the definition of callSequenceSpy here, but that doesn't work somehow
		spy := callSequenceSpy{Calls: []string{}}
		Countdown(&spy, &spy)
		got := strings.Join(spy.Calls, ",")
		want := "write,sleep,write,sleep,write,sleep,write"
		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	})
}
