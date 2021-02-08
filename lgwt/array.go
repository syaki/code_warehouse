package main

func sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func sumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, sum(numbers))
	}
	return
}

func sumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, sum(numbers[1:]))
		}
	}
	return
}
