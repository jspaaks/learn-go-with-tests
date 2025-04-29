package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("Slice of length 3", func (t *testing.T){
		var numbers = []int{1, 2, 3}
		var got int = Sum(numbers)
		var want int = 6
		if got != want {
			t.Errorf("Got %d, but want %d given %v", got, want, numbers)
		}
	})
}
