package shape

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * math.Pi
}

func (c Circle) Area() float64 {
	// return math.Pow(c.Radius, 2) * math.Pi
	return c.Radius * c.Radius * math.Pi
}

type Triangle struct {
	Base   float64
	Height float64
	Sides  []float64
}

func (t Triangle) Perimeter() float64 {
	if len(t.Sides) != 3 {
		panic("need 3 Sides")
	}

	var length float64

	for _, side := range t.Sides {
		length += side
	}
	return length
}

func (t Triangle) Area() float64 {

	return 0.5 * t.Base * t.Height

}
