package main

func Sum(numbers []int) int {
	sum := 0
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
		toAppend := 0
		if len(numbers) != 0 {
			tail := numbers[1:]
			toAppend = Sum(tail)
		}
		sums = append(sums, toAppend)
	}
	return sums
}
