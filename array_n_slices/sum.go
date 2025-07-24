package arraynslices

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		var sum int
		if len(numbers) > 0 {
			tail := numbers[1:]
			sum = Sum(tail)
		}
		sums = append(sums, sum)
	}
	return sums
}
