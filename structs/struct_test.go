package structstest

import (
	"math"
	"testing"
)

type Rectangle struct {
	Width  float64
	Heigth float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Heigth float64
}

type Shape interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("%.2f got, %.2f want", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("%g got, %g want", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

	t.Run("triangle", func(t *testing.T) {
		triangle := Triangle{12, 6}
		checkArea(t, triangle, 36.0)
	})
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Heigth)
}

func (r Rectangle) Area() float64 {
	return r.Heigth * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Heigth) * 0.5
}
