package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"golang.org/x/exp/rand"
	"golang.org/x/sync/semaphore"
)

/*
	Implement the dining philosopher’s problem with the following constraints/modifications.

		- There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
		- Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
		- The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
		- In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
		- The host allows no more than 2 philosophers to eat concurrently.
		- Each philosopher is numbered, 1 through 5.
		- When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
		- When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

type ChopStick = sync.Mutex

type Philosopher struct {
	id         int
	bites      int
	full       bool
	lChopStick *ChopStick
	rChopStick *ChopStick
}

func (p *Philosopher) Eat(chopsticks *semaphore.Weighted) {
	p.lChopStick.Lock()

	if p.bites < 3 {
		// When a philosopher starts eating (after it has obtained necessary locks)
		// it prints “starting to eat <number>” on a line by itself, where <number>
		// is the number of the philosopher.
		fmt.Printf("starting to eat %d\n", p.id)

		p.bites++

		// Each philosopher should eat only 3 times
		if p.bites > 2 {
			p.full = true
		}

		// When a philosopher finishes eating (before it has released its locks)
		// it prints “finishing eating <number>” on a line by itself, where <number>
		// is the number of the philosopher.
		fmt.Printf("finished eating %d\n", p.id)
	}

	defer p.lChopStick.Unlock()
	defer chopsticks.Release(1)
}

type Host struct{ guests []*Philosopher }

func (h *Host) StartService(wg *sync.WaitGroup) {
	defer wg.Done()

	// The philosophers pick up the chopsticks in any order, not lowest-numbered
	// first (which we did in lecture).
	rand.Seed(uint64(time.Now().Nanosecond()))
	rand.Shuffle(len(h.guests), func(i, j int) {
		h.guests[i], h.guests[j] = h.guests[j], h.guests[i]
	})

	// In order to eat, a philosopher must get permission from a host which executes
	// in its own goroutine. The host allows no more than 2 philosophers to eat
	// concurrently.
	chopsticks := semaphore.NewWeighted(2)

	for !h.EndService() {
		for _, guest := range h.guests {
			if err := chopsticks.Acquire(context.TODO(), 1); err != nil {
				fmt.Fprintf(os.Stderr, "error acquiring chopstick semaphore: %v", err)
				return
			}

			go guest.Eat(chopsticks)
		}
	}
}

func (h *Host) EndService() bool {
	for _, p := range h.guests {
		if !p.full {
			return false
		}
	}
	return true
}

const maxGuests = 5

func main() {
	// There should be 5 philosophers sharing chopsticks, with one chopstick
	// between each adjacent pair of philosophers.
	guests := make([]*Philosopher, 0, maxGuests)
	chopstx := make([]*ChopStick, 0, maxGuests)

	for range maxGuests {
		chopstx = append(chopstx, new(ChopStick))
	}

	for i := range maxGuests {
		guest := &Philosopher{
			id:         i + 1, // Each philosopher is numbered, 1 through 5.
			lChopStick: chopstx[i],
			rChopStick: chopstx[i%maxGuests],
		}
		guests = append(guests, guest)
	}

	host := &Host{guests}

	var wg sync.WaitGroup
	wg.Add(1)
	go host.StartService(&wg)
	wg.Wait()
}
