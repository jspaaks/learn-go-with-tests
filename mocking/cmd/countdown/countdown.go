package main

import (
	"github.com/jspaaks/learn-go-with-tests/mocking/pkg/countdown"
	"os"
	"time"
)

func main() {
	writer := os.Stdout
	sleeper := countdown.TimedSleeper{Duration: 1 * time.Second}
	countdown.Countdown(writer, &sleeper)
}
