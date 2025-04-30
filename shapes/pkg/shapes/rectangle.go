package shapes

// Definition of a rectangle.
type Rectangle struct {
	Width float64
	Height float64
}

// Returns the area of a Rectangle receiver r.
func (r Rectangle) CalcArea() (area float64) {
	area = r.Width * r.Height
	return
}

// Returns the perimeter of Rectangle receiver r.
func (r Rectangle) CalcPerimeter() (perimeter float64) {
	perimeter = 2 * (r.Width + r.Height)
	return
}
