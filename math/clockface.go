package math

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / float64(t.Second())))
}

func SecondsHandPoint(t time.Time) Point {
	return angleToPoint(SecondsInRadians(t))
}

func MinutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / 60) + (math.Pi / (30 / float64(t.Minute())))
}

func MinutesHandPoint(t time.Time) Point {
	return angleToPoint(MinutesInRadians(t))
}

func HourIsRadians(t time.Time) float64 {
	return (MinutesInRadians(t) / 12) + (math.Pi / (6 / float64(t.Hour()%12)))
}

func HourHandPoint(t time.Time) Point {
	return angleToPoint(HourIsRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
