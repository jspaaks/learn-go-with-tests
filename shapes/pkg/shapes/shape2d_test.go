package shapes

import "testing"

func TestCalcArea(t *testing.T) {
	records := []struct {
		desc      string
		shape     Shape2D
		perimeter float64
	}{
		{
			desc: "circle",
			shape: Circle{Radius: 10.0},
			perimeter: 314.1592653589793,
		},
		{
			desc: "rectangle",
			shape: Rectangle{Height: 10.0, Width: 20.0},
			perimeter: 200.0,
		},
		{
			desc: "right triangle",
			shape: RightTriangle{Height: 21.0, Base: 20.0},
			perimeter: 210.0,
		},
	}

	for _, record := range records {
		t.Run(record.desc, func(t *testing.T){
			got := record.shape.CalcArea()
			if got != record.perimeter {
				t.Errorf("%#v got %g, want %g", record.shape, got, record.perimeter)
			}
		})
	}
}

func TestCalcPerimeter(t *testing.T) {
	records := []struct {
		desc  string
		shape Shape2D
		area  float64
	}{
		{
			desc: "circle",
			shape: Circle{Radius: 10.0},
			area: 62.83185307179586,
		},
		{
			desc: "rectangle",
			shape: Rectangle{Height: 10.0, Width: 20.0},
			area: 60.0,
		},
		{
			desc: "right triangle",
			shape: RightTriangle{Height: 21.0, Base: 20.0},
			area: 70.0,
		},
	}

	for _, record := range records {
		t.Run(record.desc, func(t *testing.T){
			got := record.shape.CalcPerimeter()
			if got != record.area {
				t.Errorf("%#v got %g, want %g", record.shape, got, record.area)
			}
		})
	}
}
