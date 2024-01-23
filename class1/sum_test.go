package class1_test

import (
	"github.com/victoryRo/tdd/class1"
	"testing"
)

func TestAdd(t *testing.T) {
	// Arrange
	want := 7

	// Act
	got := class1.Add(4, 3)

	// Assert
	if want != got {
		t.Errorf("expected %d got %d", want, got)
	}
}

func TestAddMultiple(t *testing.T) {
	// Arrange
	values := []int{3, 7, 9, 14, 21}
	want := 54
	t.Logf("want value %d", want)

	// Act
	got := class1.AddMultiple(values...)
	t.Logf("got value %d", got)

	// Assert
	if got != want {
		t.Errorf("we want %d and got %d", want, got)
	}
	t.Log("test finish AddMultiple.")
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		class1.Add(7, 6)
	}
}

func BenchmarkAddMultiple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		class1.AddMultiple(2, 4, 6, 8, 10)
	}
}
