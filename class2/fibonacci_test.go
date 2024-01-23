package class2_test

import (
	"github.com/victoryRo/tdd/class2"
	"testing"
)

func TestFibonacciFor(t *testing.T) {
	// Arrange
	want := 55

	// Act
	got := class2.FibonacciFor(10)

	// Assert
	if got != want {
		t.Errorf("expected %d we got %d", want, got)
	}
}

func TestFibonacciRec(t *testing.T) {
	// Arrange
	want := 55

	// Act
	got := class2.FibonacciRec(10)

	// Assert
	if want != got {
		t.Errorf("expected %d we got %d", want, got)
	}
}

func TestFibonacciRecMem(t *testing.T) {
	// Arrange
	want := 55

	// Act
	got := class2.FibonacciRecMem(10)

	// Assert
	if want != got {
		t.Errorf("we expected %d and got %d", want, got)
	}
}

func BenchmarkFibonacciFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		class2.FibonacciFor(30)
	}
}

func BenchmarkFibonacciRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		class2.FibonacciRec(30)
	}
}

func BenchmarkFibonacciRecMem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		class2.FibonacciRecMem(30)
	}
}
