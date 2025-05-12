package unitclock

import (
	"math"
	"time"
)

func HourHand(tm time.Time) Point {
	hours := tm.Hour() % 12
	radians := 2 * math.Pi * float64(hours) / 12
	return radiansToPoint(radians)
}

func MinuteHand(tm time.Time) Point {
	minutes := tm.Minute()
	radians := 2 * math.Pi * float64(minutes) / 60
	return radiansToPoint(radians)
}

func SecondHand(tm time.Time) Point {
	seconds := tm.Second()
	radians := 2 * math.Pi * float64(seconds) / 60
	return radiansToPoint(radians)
}

func radiansToPoint(radians float64) Point {
	return Point{
		X: math.Sin(radians),
		Y: math.Cos(radians),
	}
}
