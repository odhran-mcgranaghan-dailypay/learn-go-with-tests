package main

import (
	"learn-go/concurrency/dining_philosophers/basic"
	"learn-go/concurrency/dining_philosophers/improved"
)

func main() {
	// running DP problem with basic solution using goroutines and mutexes
	basic.Run()

	// running DP problem with improved solution using select channels
	improved.Run()
}
