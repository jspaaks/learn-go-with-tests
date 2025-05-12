package clock

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const cx float64 = 150
const cy float64 = 150

func HourHand(tm time.Time) Point {
	const radius float64 = 50
	hours := tm.Hour() % 12
	radians := 2 * math.Pi * float64(hours) / 12
	return radiansToPoint(radians, radius)
}

func MinuteHand(tm time.Time) Point {
	const radius float64 = 80
	minutes := tm.Minute()
	radians := 2 * math.Pi * float64(minutes) / 60
	return radiansToPoint(radians, radius)
}

func SecondHand(tm time.Time) Point {
	const radius float64 = 90
	seconds := tm.Second()
	radians := 2 * math.Pi * float64(seconds) / 60
	return radiansToPoint(radians, radius)
}

func radiansToPoint(radians float64, radius float64) Point {
	return Point{
		X: cx + math.Sin(radians)*radius,
		Y: cy - math.Cos(radians)*radius,
	}
}
