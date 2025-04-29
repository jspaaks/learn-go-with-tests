package arrays

func Sum(numbers []int) (summed int) {
	if len(numbers) == 0 {
		return
	}
	for _, number := range numbers {
		summed += number
	}
	return
}

func SumAll(slices ...[]int) (summed []int) {
	for _, slice := range slices {
		summed = append(summed, Sum(slice))
	}
	return
}

func SumAllTails(slices ...[]int) (summed []int) {
	for _, slice := range slices {
		if len(slice) == 0 {
			summed = append(summed, 0)
		} else {
			summed = append(summed, Sum(slice[1:]))
		}
	}
	return
}
