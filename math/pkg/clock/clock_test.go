package clock

import (
	"fmt"
	"testing"
	"time"
)

func assertBasicallyEqualPoints(t *testing.T, a Point, b Point, tolerance float64) {
	t.Helper()
	dx := a.X - b.X
	dy := a.Y - b.Y
	tooFar := dx*dx+dy*dy > tolerance*tolerance
	if tooFar {
		t.Fatalf("%v and %v are separated by more than the tolerance %g", a, b, tolerance)
	}
}

func TestClockHourHand(t *testing.T) {
	testcases := []struct {
		hours    int
		expected Point
	}{
		{0, Point{X: 150, Y: 100}},
		{3, Point{X: 200, Y: 150}},
		{6, Point{X: 150, Y: 200}},
		{9, Point{X: 100, Y: 150}},
	}
	for _, testcase := range testcases {
		desc := fmt.Sprintf("at %d hours", testcase.hours)
		t.Run(desc, func(t *testing.T) {
			tm := time.Date(1337, time.January, 1, testcase.hours, 0, 0, 0, time.UTC)
			got := HourHand(tm)
			assertBasicallyEqualPoints(t, got, testcase.expected, 1e-6)
		})
	}
}

func TestClockMinuteHand(t *testing.T) {
	testcases := []struct {
		minutes  int
		expected Point
	}{
		{0, Point{X: 150, Y: 70}},
		{15, Point{X: 230, Y: 150}},
		{30, Point{X: 150, Y: 230}},
		{45, Point{X: 70, Y: 150}},
	}
	for _, testcase := range testcases {
		desc := fmt.Sprintf("at %d minutes", testcase.minutes)
		t.Run(desc, func(t *testing.T) {
			tm := time.Date(1337, time.January, 1, 0, testcase.minutes, 0, 0, time.UTC)
			got := MinuteHand(tm)
			assertBasicallyEqualPoints(t, got, testcase.expected, 1e-6)
		})
	}
}

func TestClockSecondHand(t *testing.T) {
	testcases := []struct {
		seconds  int
		expected Point
	}{
		{0, Point{X: 150, Y: 60}},
		{15, Point{X: 240, Y: 150}},
		{30, Point{X: 150, Y: 240}},
		{45, Point{X: 60, Y: 150}},
	}

	for _, testcase := range testcases {
		desc := fmt.Sprintf("at %d seconds", testcase.seconds)
		t.Run(desc, func(t *testing.T) {
			tm := time.Date(1337, time.January, 1, 0, 0, testcase.seconds, 0, time.UTC)
			got := SecondHand(tm)
			assertBasicallyEqualPoints(t, got, testcase.expected, 1e-6)
		})
	}
}
