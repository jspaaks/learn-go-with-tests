package unitclock

import "testing"

func TestScale(t *testing.T) {
	testcases := []struct {
		desc     string
		input    Point
		radius   float64
		expected Point
	}{
		{"north", Point{X: 0, Y: 1}, 2, Point{X: 0, Y: 2}},
		{"east", Point{X: 1, Y: 0}, 2, Point{X: 2, Y: 0}},
		{"south", Point{X: 0, Y: -1}, 2, Point{X: 0, Y: -2}},
		{"west", Point{X: -1, Y: 0}, 2, Point{X: -2, Y: 0}},
	}
	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			actual := testcase.input.Scale(testcase.radius)
			assertBasicallyEqualPoints(t, actual, testcase.expected, 1e-6)
		})
	}
}

func TestTranslate(t *testing.T) {
	testcases := []struct {
		desc     string
		input    Point
		dx   float64
		dy   float64
		expected Point
	}{
		{"north", Point{X: 0, Y: 0}, 0, 1, Point{X: 0, Y: 1}},
		{"east", Point{X: 0, Y: 0}, 1, 0, Point{X: 1, Y: 0}},
		{"south", Point{X: 0, Y: 0}, 0, -1, Point{X: 0, Y: -1}},
		{"west", Point{X: 0, Y: 0}, -1, 0, Point{X: -1, Y: 0}},
	}
	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			actual := testcase.input.Translate(testcase.dx, testcase.dy)
			assertBasicallyEqualPoints(t, actual, testcase.expected, 1e-6)
		})
	}
}
