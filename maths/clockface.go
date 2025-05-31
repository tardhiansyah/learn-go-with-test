package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30.0
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30.0
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6.0
	hoursInClock       = 2 * hoursInHalfClock
)

type Point struct {
	X float64
	Y float64
}

func secondHandPoint(t time.Time) Point {
	return radiansToPoint(secondsInRadians(t))
}

func minuteHandPoint(t time.Time) Point {
	return radiansToPoint(minutesInRadians(t))
}

func hourHandPoint(t time.Time) Point {
	return radiansToPoint(hoursInRadians(t))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) +
		(math.Pi / (hoursInHalfClock / float64(t.Hour()%12)))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / minutesInClock) + (math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (secondsInHalfClock / float64(t.Second()))
}

func radiansToPoint(radians float64) Point {
	x := math.Sin(radians)
	y := math.Cos(radians)
	return Point{
		x, y}
}
