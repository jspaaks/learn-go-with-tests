package arrays

// Sum calculates the total of a slice of numbers.
func Sum(numbers []int) (summed int) {
	accumulator := func(item int, acc int) int {
		return acc + item
	}
	return Reduce(numbers, accumulator, 0)
}

// SumAll calculates the total of a slice on each of its input slices.
func SumAll(slices ...[]int) (summed []int) {
	accumulator := func(slc []int, acc []int) []int {
		return append(acc, Sum(slc))
	}
	return Reduce(slices, accumulator, []int{})
}

// SumAllTails calculates the total of the tail of a slice on each of its input slices.
func SumAllTails(slices ...[]int) (summed []int) {
	accumulator := func(slc []int, acc []int) []int {
		if len(slc) == 0 {
			return append(acc, 0)
		}
		return append(acc, Sum(slc[1:]))
	}
	return Reduce(slices, accumulator, []int{})
}

func Reduce[A any, B any](slc []A, f func(item A, acc B) B, acc B) B {
	for _, item := range slc {
		acc = f(item, acc)
	}
	return acc
}
