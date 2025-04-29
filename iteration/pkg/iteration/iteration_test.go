package iteration

import "testing"

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func TestRepeat(t *testing.T) {
	var actual string = Repeat("a")
	var expected string = "aaaaa"
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}
