package main

import (
	"fmt"
	"runtime"
)

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func Run1() {
	fmt.Println("Before preallocation:")
	printMemUsage()
	m := make(map[string]string, 1000000) // Increased size for clearer results
	fmt.Println("After preallocation:")
	for i := 0; i < 1000000; i++ {
		m[fmt.Sprintf("key%d", i)] = "value"
	}
	runtime.GC() // Force garbage collection
	printMemUsage()

	fmt.Println("\nBefore without preallocation:")
	m = map[string]string{}
	printMemUsage()
	fmt.Println("After without preallocation:")
	for i := 0; i < 1000000; i++ {
		m[fmt.Sprintf("key%d", i)] = "value"
	}
	runtime.GC() // Force garbage collection
	printMemUsage()
}

/*
Metrics Explanation:
- Alloc: Current heap memory allocated and in use (active memory).
- TotalAlloc: Cumulative total heap memory allocated since start (includes freed memory).
- Sys: Total memory obtained from the OS (system-level memory).
- NumGC: Number of garbage collection cycles since start (memory turnover indicator).
*/

// Before preallocation:
// Alloc = 0 MiB   TotalAlloc = 0 MiB      Sys = 6 MiB     NumGC = 0
// After preallocation:
// Alloc = 0 MiB   TotalAlloc = 94 MiB     Sys = 103 MiB   NumGC = 2

// Before without preallocation:
// Alloc = 0 MiB   TotalAlloc = 94 MiB     Sys = 103 MiB   NumGC = 2
// After without preallocation:
// Alloc = 0 MiB   TotalAlloc = 271 MiB    Sys = 219 MiB   NumGC = 9
