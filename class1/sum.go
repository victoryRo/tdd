package class1

// Add addition of two numbers
func Add(a, b int) int {
	return a + b
}

// AddMultiple add several numbers
func AddMultiple(numbers ...int) int {
	res := 0
	for _, n := range numbers {
		res += n
	}
	return res
}
