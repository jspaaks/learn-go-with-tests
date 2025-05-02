package countdown

import "time"

type Sleeper interface {
	Sleep()
}

type TimedSleeper struct {
	Duration time.Duration
}

func (d *TimedSleeper) Sleep() {
	time.Sleep(d.Duration)
}
