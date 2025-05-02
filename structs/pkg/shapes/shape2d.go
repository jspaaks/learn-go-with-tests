package shapes

// Definition of a 2-D shape, containing only those methods that apply for any 2-D shape.
type Shape2D interface {
	CalcArea() float64
	CalcPerimeter() float64
}
