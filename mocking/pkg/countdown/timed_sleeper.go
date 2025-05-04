package countdown

import "time"

type Sleeper interface {
	Sleep()
}

type TimedSleeper struct {
	Duration time.Duration
	SleepFun func(time.Duration)
}

func (t *TimedSleeper) Sleep() {
	t.SleepFun(t.Duration)
}
