package main

import (
	"os"
	"time"

	"github.com/jspaaks/learn-go-with-tests/math/pkg/clock"
)

func main() {
	tm := time.Now()
	clock.SvgWriter(os.Stdout, tm)
}
