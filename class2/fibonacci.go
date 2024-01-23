package class2

func FibonacciFor(n int) int {
	a, b := 0, 1
	if n == 0 {
		return a
	}

	for i := 2; i <= n; i++ {
		c := a + b
		a, b = b, c
	}
	return b
}

func FibonacciRec(n int) int {
	if n <= 1 {
		return n
	}

	return FibonacciRec(n-1) + FibonacciRec(n-2)
}

var fibonacciData = []int{0, 1}

func FibonacciRecMem(n int) int {
	if n == 0 {
		return fibonacciData[0]
	}

	if len(fibonacciData) >= n+1 {
		return fibonacciData[n]
	}

	newData := FibonacciRecMem(n-1) + FibonacciRecMem(n-2)
	fibonacciData = append(fibonacciData, newData)

	return newData
}
