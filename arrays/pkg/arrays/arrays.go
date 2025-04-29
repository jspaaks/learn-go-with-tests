package arrays

func Sum(numbers []int) (summed int) {
	for _, number := range numbers {
		summed += number
	}
	return
}
