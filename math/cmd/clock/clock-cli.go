package main

import (
	"os"
	"time"

	"github.com/jspaaks/learn-go-with-tests/math/pkg/svg"
)

func main() {
	tm := time.Now()
	svg.Writer(os.Stdout, tm)
}
