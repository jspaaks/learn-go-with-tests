package countdown

import (
	"bytes"
	"testing"
)

type mockedSleeper struct {
	NumberOfCalls int
}

func (m *mockedSleeper) Sleep() {
	m.NumberOfCalls++
}

func TestCountdown(t *testing.T) {

	assertOutputCorrect := func(t *testing.T, buffer bytes.Buffer) {
		got := buffer.String()
		want := `3
2
1
Go!
`
		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	}

	assertThreeCallsToSleeper := func(t *testing.T, sleeper mockedSleeper) {
		const expectedNumberOfCalls = 3
		if sleeper.NumberOfCalls != expectedNumberOfCalls {
			t.Errorf("Made %d calls to Sleep(), but expected %d", sleeper.NumberOfCalls, 3)
		}
	}

	sleeper := mockedSleeper{NumberOfCalls: 0}
	buffer := bytes.Buffer{}
	Countdown(&buffer, &sleeper)
	assertOutputCorrect(t, buffer)
	assertThreeCallsToSleeper(t, sleeper)
}
