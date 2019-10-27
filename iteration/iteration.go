package iteration

func Repeat(character string, iteration int) string {
	var repeated string
	for i := 0; i < iteration; i++ {
		repeated += character
	}

	return repeated
}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}

		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
