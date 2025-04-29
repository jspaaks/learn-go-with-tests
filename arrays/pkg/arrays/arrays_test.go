package arrays

import (
	"testing"
	"slices"
)

func TestSum(t *testing.T) {

	assertEqual := func(t *testing.T, got int, want int, input []int) {
		t.Helper()
		if got != want {
			t.Errorf("Got %d, but want %d given %v", got, want, input)
		}
	}

	t.Run("empty slice", func (t *testing.T){
		var numbers = []int{}
		var got int = Sum(numbers)
		var want int = 0
		assertEqual(t, got, want, numbers)
	})

	t.Run("slice of length 3", func (t *testing.T){
		var numbers = []int{1, 2, 3}
		var got int = Sum(numbers)
		var want int = 6
		assertEqual(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {

	assertEqual := func(t *testing.T, got []int, want []int, input [][]int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("Got %d, but want %d given %v", got, want, input)
		}
	}

	t.Run("one slice", func(t *testing.T) {
		var a = []int{1, 1, 1}
		var input = [][]int{a}
		got := SumAll(input...)
		want := []int{3}
		assertEqual(t, got, want, input)
	})

	t.Run("two slices", func(t *testing.T) {
		var a = []int{1, 2}
		var b = []int{0, 9}
		var input = [][]int{a, b}
		got := SumAll(input...)
		want := []int{3, 9}
		assertEqual(t, got, want, input)
	})

}

func TestSumAllTails(t *testing.T) {

	assertEqual := func(t *testing.T, got []int, want []int, input [][]int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("Got %d, but want %d given %v", got, want, input)
		}
	}

	t.Run("one slice", func(t *testing.T) {
		var a = []int{1, 1, 1}
		var input = [][]int{a}
		got := SumAllTails(input...)
		want := []int{2}
		assertEqual(t, got, want, input)
	})

	t.Run("two slices", func(t *testing.T) {
		var a = []int{1, 2}
		var b = []int{0, 9}
		var input = [][]int{a, b}
		got := SumAllTails(input...)
		want := []int{2, 9}
		assertEqual(t, got, want, input)
	})

	t.Run("three short slices", func(t *testing.T) {
		var a = []int{1}
		var b = []int{11}
		var c = []int{111}
		var input = [][]int{a, b, c}
		got := SumAllTails(input...)
		want := []int{0, 0, 0}
		assertEqual(t, got, want, input)
	})

	t.Run("three empty slices", func(t *testing.T) {
		var a = []int{}
		var b = []int{}
		var c = []int{}
		var input = [][]int{a, b, c}
		got := SumAllTails(input...)
		want := []int{0, 0, 0}
		assertEqual(t, got, want, input)
	})
}
