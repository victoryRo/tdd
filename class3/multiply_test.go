package class3_test

import (
	"github.com/victoryRo/tdd/class3"
	"testing"
)

func TestMultipy(t *testing.T) {
	// Arrange
	want := 20

	// Act
	got := class3.Multiply(2, 10)

	// Assert
	if want != got {
		t.Fatalf("multiply expected %d and got %d\n", want, got)
	}
}

func TestMultiplyTable(t *testing.T) {
	// Arrange
	table := []struct {
		a    int
		b    int
		want int
	}{
		{2, 1, 2},
		{2, 2, 4},
		{2, 3, 6},
		{2, 4, 8},
		{2, 5, 10},
		{2, 6, 12},
		{2, 7, 14},
		{2, 8, 16},
		{2, 9, 18},
		{2, 10, 20},
	}

	for _, v := range table {
		// Act
		got := class3.Multiply(v.a, v.b)
		// Assert
		if got != v.want {
			t.Errorf("multiply expected %d and got %d\n", v.want, got)
		}
	}
}

func TestMultiply2All(t *testing.T) {
	ten := 10

	for i := 0; i < ten; i++ {
		// Arrange
		want := 2 * i
		// Act
		got := class3.Multiply(2, i)
		// Assert
		if want != got {
			t.Errorf("multiply expected %d and got %d\n", want, got)
		}
	}
}

// table Driven works with subtest
func TestMultiplySubtest(t *testing.T) {
	// Arrange
	table := []struct {
		testMsg string
		x       int
		y       int
		want    int
	}{
		{"3 x 1", 3, 1, 3},
		{"3 x 2", 3, 2, 6},
		{"3 x 3", 3, 3, 9},
		{"3 x 4", 3, 4, 12},
		{"3 x 9", 3, 9, 27},
	}

	for _, v := range table {
		// subTest
		t.Run(v.testMsg, func(t *testing.T) {
			// Act
			got := class3.Multiply(v.x, v.y)

			// Assert
			if v.want != got {
				t.Errorf("expected %d and got %d\n", v.want, got)
			}
		})
	}
}
