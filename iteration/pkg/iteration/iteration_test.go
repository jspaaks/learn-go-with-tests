package iteration

import (
	"fmt"
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	var s string = Repeat("$", 3)
	fmt.Println(s)
	// Output: $$$
}

func TestRepeat(t *testing.T) {
	t.Run("with 0 repeats", func (t *testing.T){
		var actual string = Repeat("a", 0)
		var expected string = ""
		if actual != expected {
			t.Errorf("expected %q but got %q", expected, actual)
		}
	})

	t.Run("with 3 repeats", func (t *testing.T){
		var actual string = Repeat("a", 3)
		var expected string = "aaa"
		if actual != expected {
			t.Errorf("expected %q but got %q", expected, actual)
		}
	})

	t.Run("with 5 repeats", func (t *testing.T){
		var actual string = Repeat("a", 5)
		var expected string = "aaaaa"
		if actual != expected {
			t.Errorf("expected %q but got %q", expected, actual)
		}
	})

	t.Run("with multicharacter string, 2 repeats", func (t *testing.T){
		var actual string = Repeat("abcd", 2)
		var expected string = "abcdabcd"
		if actual != expected {
			t.Errorf("expected %q but got %q", expected, actual)
		}
	})

}
