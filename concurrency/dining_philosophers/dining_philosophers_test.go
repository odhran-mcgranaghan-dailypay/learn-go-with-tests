package dining_philosophers

import (
	"testing"
)

func BenchmarkDiningPhilosophers(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Run()
	}
}
