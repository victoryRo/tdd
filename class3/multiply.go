package class3

// Multiply two numbers
func multiply(a, b int) int {
	return a * b
}

func Multiply(a, b int) int {
	// le damos acceso externo a este methodo
	// para testearlo desde afuera
	return multiply(a, b)
}
