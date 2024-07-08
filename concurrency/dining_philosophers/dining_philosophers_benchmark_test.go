package dining_philosophers

import (
	"learn-go/concurrency/dining_philosophers/basic"
	"learn-go/concurrency/dining_philosophers/improved"
	"os"
	"testing"
)

func suppressOutput() func() {
	nullFile, _ := os.Open(os.DevNull)
	old := os.Stdout
	os.Stdout = nullFile
	return func() {
		os.Stdout = old
		nullFile.Close()
	}
}

func BenchmarkBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		restore := suppressOutput()
		basic.Run()
		restore()
	}
}

func BenchmarkImproved(b *testing.B) {
	for i := 0; i < b.N; i++ {
		restore := suppressOutput()
		improved.Run()
		restore()
	}
}

// TYPICAL RESULTS
// go test -bench .
// goos: darwin
// goarch: arm64
// pkg: learn-go/concurrency/dining_philosophers
// BenchmarkBasic-11                      1        1616049375 ns/op
// BenchmarkImproved-11                   2         909242646 ns/op
// PASS
// ok      learn-go/concurrency/dining_philosophers        4.984s
