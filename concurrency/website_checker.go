package concurrency

import (
	"sync"
)

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultsChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultsChannel
		results[r.string] = r.bool
	}

	return results
}

func CheckWebsitesSlow(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}

// MICRO-LESSON
// sync.Map type
// Why use this over higher-level abstractions like a map with channels?
//  1. given key-value, the value is written once but read many times - most efficient when majority of operations are reads
//  2. multiple go routines read, write and overwrite the entires for disjoint set of keys - go rountines working on unique subset of keys - no overlap
func CheckWebsitesSyncMap(wc WebsiteChecker, urls []string) map[string]bool {
	// create the sync Map, and a wait group for the map
	var results sync.Map
	// wait group waits for all goroutines to finish before proceeding
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			results.Store(u, wc(u))
		}(url)
	}

	wg.Wait()

	finalResults := make(map[string]bool)
	results.Range(func(key, value interface{}) bool {
		finalResults[key.(string)] = value.(bool)
		return true
	})

	return finalResults
}
