package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	want := "aaaaa"
	got := Repeat("a", 5)
	AssertEqualsString(t, got, want)
}

func AssertEqualsString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// Benchmark functions
// to run the benchmark, $ go test -bench=.
// b.N is a value provided by the testing framework to let the code know how many times to run
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

// Example functions
func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

// Exercise:
// Change the test so a caller can specify how many times the character is repeated and then fix the code
// Have a look through the strings package. Find functions you think could be useful and experiment with them by writing tests like we have here. Investing time learning the standard library will really pay off over time.

// Sample result from the above benchmark
// The results show that the function runs 17203364 times at an average of 69.77 ns per operation.
// goos: darwin
// goarch: arm64
// pkg: learn-go/iteration
// BenchmarkRepeat-11      17203364                69.77 ns/op
