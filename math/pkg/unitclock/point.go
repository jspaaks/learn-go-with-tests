package unitclock

type Point struct {
	X float64
	Y float64
}

func (p Point) Translate(dx float64, dy float64) (translated Point) {
	translated.X = p.X + dx
	translated.Y = p.Y + dy
	return translated
}

func (p Point) Scale(radius float64) (scaled Point) {
	scaled.X = p.X * radius
	scaled.Y = p.Y * radius
	return scaled
}

func (p Point) FlipVerticalDirection() (flipped Point) {
	flipped.X = p.X
	flipped.Y = -p.Y
	return flipped
}
