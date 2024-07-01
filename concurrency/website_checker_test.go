package concurrency

import (
	"reflect"
	"testing"
	"time"
)

// Mock out the website checker function
func mockWebsiteChecker(url string) bool {
	if url == "www.google.com" {
		return false
	}
	return true
}

// stub mock method to simulate website check taking 20ms
// a stub is a type of test double that is used to provide a specific response to a method
// in this case simulating a website check that would take e.g. 20ms
func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {

	urls := make([]string, 100)
	for i := range urls {
		urls[i] = "some url"
	}

	b.Run("benchmark concurrent website checker", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			CheckWebsites(slowStubWebsiteChecker, urls)
		}
	})

	b.Run("benchmark linear website checker", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			CheckWebsitesSlow(slowStubWebsiteChecker, urls)
		}
	})

	// TYPICAL RESULTS

	// Factor of 100 difference in speed between concurrent and linear

	// BenchmarkCheckWebsites/benchmark_concurrent_website_checker-11                55          21146814 ns/op
	// BenchmarkCheckWebsites/benchmark_linear_website_checker-11                     1        2092888375 ns/op
}

func TestCheckWebsites(t *testing.T) {
	t.Run("test website checker", func(t *testing.T) {
		urls := []string{"www.google.com", "www.facebook.com", "www.twitter.com"}

		want := map[string]bool{
			"www.google.com":   false,
			"www.facebook.com": true,
			"www.twitter.com":  true,
		}

		got := CheckWebsites(mockWebsiteChecker, urls)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
}
