package dining_philosophers

import (
	"fmt"
	"sync"
	"time"
)

// Five philosophers sit at a round table with a bowl of spaghetti and five forks.
// Each philosopher alternates between thinking and eating.
// To eat, a philosopher needs both the left and right forks.
// Forks can only be used by one philosopher at a time.
// After eating, they put down both forks.
// The challenge is to design a concurrent algorithm ensuring that no philosopher starves
// and they continue to alternate between thinking and eating indefinitely,
// despite not knowing when others want to eat or think.

// How to solve this problem using Go's concurrency primitives:
// 		Goroutines: Each philosopher runs as a separate goroutine, allowing them to operate concurrently.
// 		Mutexes: Each fork is protected by a sync.Mutex, ensuring that only one philosopher can use a fork at a time, preventing race conditions.
// 		Think and Eat Cycles: Philosophers alternate between thinking and eating. When eating, they lock both their left and right forks (mutexes) to avoid conflicts.
// 		WaitGroup: A sync.WaitGroup ensures the main function waits until all philosophers have finished their dining cycles.

const philosephersAtDinner = 5

var wg sync.WaitGroup

type Fork struct {
	sync.Mutex
}

type Philosopher struct {
	id                  int
	leftFork, rightFork *Fork
}

func (p Philosopher) dine() {
	// defer wg.Done() ensures that the WaitGroup counter is decremented when the function completes.
	// When the philosopher finishes either eating or thinking, the loop counter is decremented by 1
	// this loops runs for 3 iterations
	defer wg.Done()
	for i := 0; i < 3; i++ {
		p.think()
		p.eat()
	}
}

func (p Philosopher) think() {
	fmt.Printf("Philosopher %d is thinking\n", p.id)
	time.Sleep(time.Millisecond * 100)
}

func (p Philosopher) eat() {
	// locks both forks required to eat
	p.leftFork.Lock()
	p.rightFork.Lock()

	fmt.Printf("Philosopher %d is eating\n", p.id)
	time.Sleep(time.Millisecond * 100)

	// unlocks the forks freeing them up for the next philosopher
	p.rightFork.Unlock()
	p.leftFork.Unlock()
}

func Run() {
	// setup forks, one for each philosopher
	forks := make([]*Fork, philosephersAtDinner)
	for i := 0; i < philosephersAtDinner; i++ {
		forks[i] = &Fork{}
	}

	// setup philosophers
	philosophers := make([]*Philosopher, philosephersAtDinner)
	for i := 0; i < philosephersAtDinner; i++ {
		// the dining table is a circle, so the right fork of the last philosopher is the first fork
		// capture this using modulo, to ensure this wraps around from last to first philosopher
		philosophers[i] = &Philosopher{
			id:        i + 1,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%philosephersAtDinner],
		}
	}

	// start dinner!
	// increment the WaitGroup counter by the number of philosophers
	wg.Add(philosephersAtDinner)
	// start a goroutine for each philosopher
	for _, p := range philosophers {
		go p.dine()
	}

	// block completion of this Run function until all philosophers have finished their dining cycles
	wg.Wait()

}

// Improvements - Channels and Select