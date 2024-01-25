package arrayandslices

func Sum(array []int) int {
	var sum int
	for _, num := range array {
		sum += num
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	var sums []int

	for _, numbers := range slices {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
