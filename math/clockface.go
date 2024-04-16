package main

import (
	"math"
	"time"
)

type Point struct {
	X int
	Y int
}

func SecondHand(t time.Time) Point {
	return Point{150, 90}
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / float64(t.Second())))
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
