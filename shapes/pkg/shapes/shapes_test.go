package structs

import "testing"

func TestCalcPerimeter(t *testing.T) {

	checkPerimeter := func(t *testing.T, shape Shape, want float64){
		t.Helper()
		var got float64 = shape.CalcPerimeter()
		if got != want {
			t.Errorf("Got %g, want %g", got, want)
		}
	}

	t.Run("circle", func(t *testing.T){
		var circle = Circle{10.0}
		var want float64 = 62.83185307179586
		checkPerimeter(t, circle, want)
	})
	t.Run("rectangle", func(t *testing.T){
		var rectangle = Rectangle{10.0, 20.0}
		var want float64 = 60.0
		checkPerimeter(t, rectangle, want)
	})
}

func TestCalcArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64){
		t.Helper()
		var got float64 = shape.CalcArea()
		if got != want {
			t.Errorf("Got %g, want %g", got, want)
		}
	}

	t.Run("circle", func (t *testing.T){
		var circle = Circle{10.0}
		const want float64 = 314.1592653589793
		checkArea(t, circle, want)
	})

	t.Run("rectangle", func (t *testing.T){
		var rectangle = Rectangle{10.0, 20.0}
		var want float64 = 200.0
		checkArea(t, rectangle, want)
	})

}
