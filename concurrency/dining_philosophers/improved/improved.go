package improved

import (
	"fmt"
	"sync"
	"time"
)

// Define the Fork as a channel of bool
type Fork chan bool

// Define the Philosopher struct with an id and two Fork channels
type Philosopher struct {
	id                  int
	leftFork, rightFork chan bool
}

func NewPhilosopher(id int, leftFork, rightFork chan bool) Philosopher {
	return Philosopher{
		id:        id,
		leftFork:  leftFork,
		rightFork: rightFork,
	}
}

// create a new Fork channel, buffered channel with capacity of 1
func NewFork() Fork {
	return make(chan bool, 1)
}

func (p Philosopher) BeginDining(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		p.think()
		p.eat()
	}
}

// Philosopher will contemplate life for a while - here for 100ms
func (p Philosopher) think() {
	fmt.Printf("Philosopher %d is thinking\n", p.id)
	time.Sleep(time.Millisecond * 100)
}

func (p Philosopher) eat() {
	for {
		p.leftFork <- true
		select {
		case p.rightFork <- true:
			// The Philosopher has picked up both forks and can now eat
			fmt.Printf("Philosopher %d is eating\n", p.id)
			time.Sleep(time.Millisecond * 100)
			// Finished eating, release the forks
			<-p.leftFork
			<-p.rightFork
			return
		default:
			// didn't get the right fork, release the left fork
			<-p.leftFork
		}
		// failed to get either fork, think again
	}
}

func Run() {
	numPhilosophers := 5
	var wg sync.WaitGroup

	fmt.Println("DINNER WITH SELECT AND CHANNELS!")

	// create forks
	forks := make([]Fork, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		forks[i] = NewFork()
	}

	// create philosophers
	philosophers := make([]Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		leftFork := forks[i]
		rightFork := forks[(i+1)%numPhilosophers]
		philosophers[i] = NewPhilosopher(i+1, leftFork, rightFork)
	}

	// add philosopher count to the WaitGroup
	wg.Add(numPhilosophers)

	// begin dining
	for _, p := range philosophers {
		go p.BeginDining(&wg)
	}

	// wait for dinner to finish
	wg.Wait()

	fmt.Println("Dinner is over!")
}
