package clock

import "testing"

func assertBasicallyEqualPoints(t *testing.T, a Point, b Point, tolerance float64) {
	t.Helper()
	dx := a.X - b.X
	dy := a.Y - b.Y
	tooFar := dx*dx+dy*dy > tolerance*tolerance
	if tooFar {
		t.Fatalf("%v and %v are separated by more than the tolerance %g", a, b, tolerance)
	}
}

func assertLinesIncludeGivenLine(t *testing.T, lines []Line, want Line, s string) {
	t.Helper()
	for _, got := range lines {
		if got == want {
			return
		}
	}
	t.Fatalf("Expected to find the second hand with x2 of %+v and y2 of %+v in the SVG output %v", want.X2, want.Y2, s)
}
