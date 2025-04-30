package shapes

import "math"

// Definition of a right triangle.
type RightTriangle struct {
	Base float64
	Height float64
}

// Returns the area of the RightTriangle receiver r.
func (r RightTriangle) CalcArea() float64 {
	return 0.5 * r.Height * r.Base
}

// Returns the lenght of the hypothenuse of the RightTriangle receiver r.
func (r RightTriangle) CalcHypothenuse() float64 {
	return math.Sqrt(r.Base * r.Base + r.Height * r.Height)
}

// Returns the perimeter of the RightTriangle receiver r.
func (r RightTriangle) CalcPerimeter() float64 {
	return r.Base + r.Height + r.CalcHypothenuse()
}
