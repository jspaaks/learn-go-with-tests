package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd_both_arguments_positive() {
	var a int = 4
	var b int = 5
	var summed int = Add(a, b)
	fmt.Println(summed)
	// Output: 9
}

func ExampleAdd_both_arguments_negative() {
	var a int = -4
	var b int = -5
	var summed int = Add(a, b)
	fmt.Println(summed)
	// Output: -9
}

func TestAdd(t *testing.T) {
	var got int = Add(4, 5)
	var want int = 9
	if want != got {
		t.Errorf("wanted '%d', got '%d' instead", want, got)
	}
}
