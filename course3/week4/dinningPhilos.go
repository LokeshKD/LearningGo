/*

Program to Demonstrate Dinning Philosophers problem with some constraints.

1.	There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
2.	Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
3.	The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
4.	In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
5.	The host allows no more than 2 philosophers to eat concurrently.
6.	Each philosopher is numbered, 1 through 5.

7.	When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

8.	When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type ChopS struct{ sync.Mutex }
type Philo struct {
	num         int
	left, right *ChopS
}

const res = 5 // Number of Resources

func (p Philo) eat(perm chan bool) {

	defer wg.Done() // First things first!

	// Design Principle 2. Each Philo should eat max 3 times.
	for i := 0; i < 3; i++ {
		// Design Principle 4. Get permission
		<-perm

		// Design Priciple 3. Pickup chops in any order.
		p.left.Lock()
		p.right.Lock()

		// Design Priciple 7. Print Eating message.
		fmt.Printf("starting to eat %d\n", p.num)

		// Simulate Eating.
		time.Sleep(1000 * time.Millisecond)

		// Design Priciple 8. Print finishing message.
		fmt.Printf("finishing eating %d\n", p.num)

		p.right.Unlock()
		p.left.Unlock()
	}
}

func main() {
	// Create an array of 5 Chopsticks and Philosophers.
	CSticks := make([]*ChopS, 5)
	philos := make([]*Philo, 5)

	// Channel to permit at max 2 Philos to Eat. Design Principle 5.
	perm := make(chan bool, 2)

	for i := 0; i < res; i++ {
		CSticks[i] = new(ChopS)
	}
	// Design Principle 1. 5 Philosophers.
	for i := 0; i < res; i++ {
		// Design Principle 6. Philos are numbered 1 thru 5.
		// First element in Philo initialization i+1 is for that.
		philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%res]}
	}
	// Motivate all Philos to start Eating.
	for i := 0; i < res; i++ {
		wg.Add(1)
		go philos[i].eat(perm)
	}

	// Each Philo can eat for 3 times max and he needs permission each time.
	// We have 5 Phiols, so we need to permit at max for 15 times.
	for i := 0; i < res*3; i++ {
		perm <- true
		if 1 == i%2 {
			time.Sleep(100 * time.Millisecond)
		} // Delay Permission.
	}
	wg.Wait()
}
