package microlesson

import "fmt"

func Run() {
	go sayHello()
	go func() {
		fmt.Println("cenas from lesson 1")
	}()
	// continue do other things.
}

func sayHello() {
	fmt.Println("hello world from lesson 1")
}

func ConfinementPattern() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i < 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("received: %d\n", result)
		}
		fmt.Println("done receiving")
	}

	results := chanOwner()
	consumer(results)
}
