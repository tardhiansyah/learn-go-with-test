package arrays

func Sum(numbers []int) int {
	add := func(acc, number int) int {
		return acc + number
	}
	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	sum := func(acc, numbers []int) []int {
		return append(acc, Sum(numbers))
	}

	return Reduce(numbersToSum, sum, []int{})
}

func SumAllTails(numbersToSum ...[]int) []int {
	sumTail := func(acc, tail []int) []int {
		if len(tail) == 0 {
			return append(acc, 0)
		}
		return append(acc, Sum(tail[1:]))
	}

	return Reduce(numbersToSum, sumTail, []int{})
}

func Reduce[T any](numbers []T, f func(T, T) T, initialValue T) T {
	result := initialValue
	for _, number := range numbers {
		result = f(result, number)
	}
	return result
}
