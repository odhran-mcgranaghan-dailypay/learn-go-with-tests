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

// Example functions
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
