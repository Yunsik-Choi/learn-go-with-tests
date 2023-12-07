package array

func Sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func SumAll(ints ...[]int) []int {
	var result []int
	for _, numbers := range ints {
		result = append(result, Sum(numbers))
	}
	return result
}
