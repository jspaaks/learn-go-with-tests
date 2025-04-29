package structs

import "math"

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width float64
	Height float64
}

type Shape interface {
	CalcArea() float64
	CalcPerimeter() float64
}

func (c Circle) CalcArea() (area float64) {
	area = math.Pi * c.Radius * c.Radius
	return
}

func (c Circle) CalcPerimeter() (perimeter float64) {
	perimeter = 2 * math.Pi * c.Radius
	return
}

func (r Rectangle) CalcArea() (area float64) {
	area = r.Width * r.Height
	return
}

func (r Rectangle) CalcPerimeter() (perimeter float64) {
	perimeter = 2 * (r.Width + r.Height)
	return
}
