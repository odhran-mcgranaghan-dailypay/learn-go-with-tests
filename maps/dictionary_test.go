package maps

import (
	"fmt"
	"testing"
)

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, d Dictionary, word, definition string) {
	t.Helper()
	str, err := d.Search(word)
	if err != nil {
		t.Fatalf("should find added word: %v", err)
	}
	assertStrings(t, str, definition)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		want := "could not find the word you were looking for"
		if got == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, got, ErrNotFound)
		assertStrings(t, got.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Add existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		word := "coffee"
		definition := "hot, black bean water"
		updatedDefinition := "hot, black bean water now with caffeine"
		d := Dictionary{word: definition}
		err := d.Update(word, updatedDefinition)

		assertError(t, err, nil)
		assertDefinition(t, d, word, updatedDefinition)
	})

	t.Run("attempt update on non-existent word", func(t *testing.T) {
		word := "coffee"
		definition := "hot, black bean water"
		updatedDefinition := "hot, black bean water now with caffeine"
		d := Dictionary{word: definition}
		err := d.Update("tea", updatedDefinition)

		assertError(t, err, ErrUpdateWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing word", func(t *testing.T) {
		word := "coffee"
		definition := "hot, black bean water"
		d := Dictionary{word: definition}
		d.Delete(word)
		_, err := d.Search(word)

		assertError(t, err, ErrNotFound)
	})

	t.Run("attempt delete of non existent word", func(t *testing.T) {
		word := "coffee"
		definition := "hot, black bean water"
		d := Dictionary{word: definition}
		err := d.Delete("tea")

		assertError(t, err, ErrDeleteWordDoesNotExist)
	})
}

// BenchmarkMapAccess benchmarks the time it takes to access a value in a map
// Lesson - preallocate the map to improve performance

// Optimisation:
// Preallocating the map with a specified capacity using make(map[string]string, capacity).

// Measurement:
// 		use go test -bench=. -benchmem to run the benchmark and measure memory usage and allocations.
// 		Benchmark tests compare performance (ns/op),
// 								memory usage (B/op),
// 								allocations (allocs/op)

// With Preallocation:
// 		Time per operation: 64096 ns
// 		Memory per operation: 95659 B
// 		Allocations per operation: 1746
// Without Preallocation:
// 		Time per operation: 91287 ns
// 		Memory per operation: 180270 B
// 		Allocations per operation: 1774
//
// Preallocating improves performance by ~30%, reduces memory usage by ~50%, allocations per op remain similar

func BenchmarkMapCreationWithPreallocation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := make(map[string]string, 1000)
		for i := 0; i < 1000; i++ {
			m[fmt.Sprintf("key%d", i)] = "value"
		}
	}
}

func BenchmarkMapCreationWithoutPreallocation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := map[string]string{}
		for i := 0; i < 1000; i++ {
			m[fmt.Sprintf("key%d", i)] = "value"
		}
	}
}

// Searching is efficient with preallocated maps due to evenly distributed hash buckets
func BenchmarkMapSearchWithPreallocation(b *testing.B) {
	m := make(map[string]string, 1000)
	for i := 0; i < 1000; i++ {
		m[fmt.Sprintf("key%d", i)] = "value"
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = m["key500"]
	}
}

func BenchmarkMapSearchWithoutPreallocation(b *testing.B) {
	m := map[string]string{}
	for i := 0; i < 1000; i++ {
		m[fmt.Sprintf("key%d", i)] = "value"
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = m["key500"]
	}
}

func BenchmarkMapUpdateWithPreallocation(b *testing.B) {
	m := make(map[string]string, 1000)
	for i := 0; i < 1000; i++ {
		m[fmt.Sprintf("key%d", i)] = "value"
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		m["key500"] = "new_value"
	}
}

func BenchmarkMapUpdateWithoutPreallocation(b *testing.B) {
	m := map[string]string{}
	for i := 0; i < 1000; i++ {
		m[fmt.Sprintf("key%d", i)] = "value"
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		m["key500"] = "new_value"
	}
}

func BenchmarkMapDeleteWithPreallocation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := make(map[string]string, 1000)
		for i := 0; i < 1000; i++ {
			m[fmt.Sprintf("key%d", i)] = "value"
		}
		b.StopTimer()
		delete(m, "key500")
		b.StartTimer()
	}
}

func BenchmarkMapDeleteWithoutPreallocation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := map[string]string{}
		for i := 0; i < 1000; i++ {
			m[fmt.Sprintf("key%d", i)] = "value"
		}
		b.StopTimer()
		delete(m, "key500")
		b.StartTimer()
	}
}

// MAP CONTEXT

// 		When you create a map in Go, go makes a best guess at the size of the map it will create. Starts small and grows
// 			- As you add more elements the map grows in size, reallocating memory, copying the old elements to the new map, this has an overhead
// 			- This overhead causes operations like search, update, delete to be slower and use more memory
//		When a map reaches capacity, Go will automatically resize the map to accommodate more elements roughly doubling the memory allocation
//		Pre-allocating avoids the overhead of dynamic memory reallocation of the map
// 			- Caveat is you have to know the size of the map in advance


// CREATE MAP

// 		Creation using mem pre-allocation shows big gains
//		~30% faster operation, uses 50% less memory per operation, similar allocations per operation
//  	Why? Preallocating the map avoids the overhead of map resizing and element copying

// BenchmarkMapAccessWithPreallocation-11             19114             62377 ns/op           95659 B/op       1746 allocs/op
// BenchmarkMapAccessWithoutPreallocation-11          13092             92071 ns/op          180264 B/op       1774 allocs/op

// SEARCH MAP

// 		Search of a map in Go is already really efficient, this is due to the evenly distributed hash buckets
// 		Preallocation has the benefit of
//  		- Optimal Bucket allocation - the optimal no. of buckets are allocated at the start
//  		- No resizing - every resizing involves rehashing all values to keys, which can increase chance of collisions
//  		- Bonus Note: Go does not allow implementation of your own Hash function, unlike other languages like Java, C++

// BenchmarkMapSearchWithPreallocation-11          175075912                5.861 ns/op           0 B/op          0 allocs/op
// BenchmarkMapSearchWithoutPreallocation-11       166815196                6.569 ns/op           0 B/op          0 allocs/op

// UPDATE IN MAP

//  	Update with Preallocation is slighlty faster here
//		Compared with SEARCH op, the slight overhead is from the time to overwrite existing values

// BenchmarkMapUpdateWithPreallocation-11          191527249                8.066 ns/op           0 B/op          0 allocs/op
// BenchmarkMapUpdateWithoutPreallocation-11       155720758                7.369 ns/op           0 B/op          0 allocs/op

// DELETE FROM MAP

//		Delete with Preallocation is faster here, comparable with the CREATE benchmark
//		~30% faster operation, uses 50% less memory per operation, similar allocations per operation

// BenchmarkMapDeleteWithPreallocation-11             17680             68143 ns/op           95669 B/op       1746 allocs/op
// BenchmarkMapDeleteWithoutPreallocation-11          12307             96839 ns/op          180262 B/op       1774 allocs/op
