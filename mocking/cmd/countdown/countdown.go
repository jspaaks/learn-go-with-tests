package main

import (
	"github.com/jspaaks/learn-go-with-tests/mocking/pkg/countdown"
	"os"
)

func main() {
	countdown.Countdown(os.Stdout)
}
