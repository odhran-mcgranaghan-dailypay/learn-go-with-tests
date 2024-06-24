package integers

import (
	"fmt"
	"testing"
)

// Test function name must start with Test to be detected by the go test command
func TestAdder(t *testing.T) {
	t.Run("adding two numbers test", func(t *testing.T) {
		sum := Add(2, 3)
		expected := 5
		assertEqualsInteger(t, sum, expected)
	})
}

func assertEqualsInteger(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}

// Defined numbers as a slice, but numbers... notation passes the slice as individual arguments
// AddMultiple(numbers...) is the same as AddMultiple(1, 2, 3, 4, 5)
func TestAddMultiple(t *testing.T) {
	t.Run("add multiple numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := AddMultiple(numbers...)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("add multiple negative numbers", func(t *testing.T) {
		numbers := []int{-1, -2, -3, -4, -5}
		got := AddMultiple(numbers...)
		want := -15
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("add mixed positive and negative numbers", func(t *testing.T) {
		numbers := []int{-1, 2, -3, 4, -5}
		got := AddMultiple(numbers...)
		want := -3
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("add no numbers", func(t *testing.T) {
		got := AddMultiple()
		want := 0
		if got != want {
			t.Errorf("got %d want %d when no numbers were given", got, want)
		}
	})
}

// Example functions
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
