package ch1

func SimpleTestFunc(numbers ...int) []int {
	return numbers
}

func AddInt(num ...int) int {
	sum := 0
	for _, n := range num {
		sum += n
	}
	return sum
}
