package shapes

import "math"

// Definition of a circle.
type Circle struct {
	Radius float64
}

// Returns the area of circle c.
func (c Circle) CalcArea() (area float64) {
	area = math.Pi * c.Radius * c.Radius
	return
}

// Returns the perimeter of circle c.
func (c Circle) CalcPerimeter() (perimeter float64) {
	perimeter = 2 * math.Pi * c.Radius
	return
}
