package arrays_test

import (
	"github.com/jspaaks/learn-go-with-tests/arrays-revisited/pkg/arrays"
	"github.com/jspaaks/learn-go-with-tests/arrays-revisited/pkg/assertions"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	assertEqual := func(t *testing.T, got int, want int, input []int) {
		t.Helper()
		if got != want {
			t.Errorf("Got %d, but want %d given %v", got, want, input)
		}
	}

	t.Run("empty slice", func(t *testing.T) {
		var numbers = []int{}
		var got int = arrays.Sum(numbers)
		var want int = 0
		assertEqual(t, got, want, numbers)
	})

	t.Run("slice of length 3", func(t *testing.T) {
		var numbers = []int{1, 2, 3}
		var got int = arrays.Sum(numbers)
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

	t.Run("empty slice", func(t *testing.T) {
		var a = []int{}
		var input = [][]int{a}
		got := arrays.SumAll(input...)
		want := []int{0}
		assertEqual(t, got, want, input)
	})

	t.Run("one slice", func(t *testing.T) {
		var a = []int{1, 1, 1}
		var input = [][]int{a}
		got := arrays.SumAll(input...)
		want := []int{3}
		assertEqual(t, got, want, input)
	})

	t.Run("two slices", func(t *testing.T) {
		var a = []int{1, 2}
		var b = []int{0, 9}
		var input = [][]int{a, b}
		got := arrays.SumAll(input...)
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
		got := arrays.SumAllTails(input...)
		want := []int{2}
		assertEqual(t, got, want, input)
	})

	t.Run("two slices", func(t *testing.T) {
		var a = []int{1, 2}
		var b = []int{0, 9}
		var input = [][]int{a, b}
		got := arrays.SumAllTails(input...)
		want := []int{2, 9}
		assertEqual(t, got, want, input)
	})

	t.Run("three short slices", func(t *testing.T) {
		var a = []int{1}
		var b = []int{11}
		var c = []int{111}
		var input = [][]int{a, b, c}
		got := arrays.SumAllTails(input...)
		want := []int{0, 0, 0}
		assertEqual(t, got, want, input)
	})

	t.Run("three empty slices", func(t *testing.T) {
		var a = []int{}
		var b = []int{}
		var c = []int{}
		var input = [][]int{a, b, c}
		got := arrays.SumAllTails(input...)
		want := []int{0, 0, 0}
		assertEqual(t, got, want, input)
	})
}

func TestReduce(t *testing.T) {
	t.Run("summing an array of int", func(t *testing.T) {
		summer := func(item int, acc int) int {
			return acc + item
		}
		slc := []int{2, 7, 3, 6}
		reduced := arrays.Reduce(slc, summer, 0)
		expected := 18
		if reduced != expected {
			t.Errorf("Got %d instead of expected %d", reduced, expected)
		}
	})
	t.Run("concatenating an array of string", func(t *testing.T) {
		concatenator := func(item string, acc string) string {
			return acc + item
		}
		slc := []string{"abc", "de", "fghi"}
		reduced := arrays.Reduce(slc, concatenator, "")
		expected := "abcdefghi"
		if reduced != expected {
			t.Errorf("Got %q instead of expected %q", reduced, expected)
		}
	})
	t.Run("multiplying an array of int", func(t *testing.T) {
		multiplier := func(item int, acc int) int {
			return acc * item
		}
		slc := []int{1, 2, 3}
		reduced := arrays.Reduce(slc, multiplier, 1)
		expected := 6
		if reduced != expected {
			t.Errorf("Got %d instead of expected %d", reduced, expected)
		}
	})
	t.Run("summing an array of int into a float64", func(t *testing.T) {
		summer := func(item int, acc float64) float64 {
			return acc + float64(item)
		}
		slc := []int{1, 2, 3}
		reduced := arrays.Reduce(slc, summer, 0)
		expected := 6.0
		if reduced != expected {
			t.Errorf("Got %g instead of expected %g", reduced, expected)
		}
	})
}

func TestFind(t *testing.T) {
	t.Run("nonexistent even numbered item from []int", func(t *testing.T) {
		input := []int{9, 5, 3, 7}
		predicate := func(item int) bool {
			return item%2 == 0
		}
		found, value := arrays.Find(input, predicate)
		assertions.AssertFalse(t, found)
		assertions.AssertEqual(t, value, 0)
	})
	t.Run("existing odd numbered item from []int", func(t *testing.T) {
		input := []int{9, 5, 3, 7}
		predicate := func(item int) bool {
			return item%2 == 1
		}
		found, value := arrays.Find(input, predicate)
		assertions.AssertTrue(t, found)
		assertions.AssertEqual(t, value, 9)
	})
	t.Run("person over age 30 from slice of person", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		valerie := Person{Name: "Valerie", Age: 25}
		mike := Person{Name: "Mike", Age: 45}
		chris := Person{Name: "Chris", Age: 33}
		persons := []Person{valerie, mike, chris}
		predicate := func(person Person) bool {
			return person.Age > 30
		}
		found, person := arrays.Find(persons, predicate)
		assertions.AssertTrue(t, found)
		assertions.AssertEqual(t, person, mike)
	})
}
