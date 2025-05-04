package countdown

import (
	"testing"
	"time"
)

type sleepDurationSpy struct {
    slept time.Duration
}

func (s *sleepDurationSpy) sleep(duration time.Duration) {
    s.slept = duration
}

func TestTimedSleeper(t *testing.T) {
	// I'd prefer to put the definition of sleepDurationSpy here, but that doesn't work somehow
	expectedSleepDuration := 5.0 * time.Second
	spy := sleepDurationSpy{}
	sleeper := TimedSleeper{Duration: expectedSleepDuration, SleepFun: spy.sleep}
	sleeper.Sleep()
	if spy.slept != expectedSleepDuration {
		t.Errorf("Expected a sleep duration of %v but got %v", expectedSleepDuration, spy.slept)
	}
}
