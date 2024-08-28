package main

import (
	"math/rand"
	"fmt"
	"sync"
	"time"
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	number         int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

type Host struct {
	sync.Mutex
	eating int
}

func (h *Host) requestPermission() bool {
	h.Lock()
	defer h.Unlock()
	if h.eating < 2 {
		h.eating++
		return true
	}
	return false
}

func (h *Host) finishEating() {
	h.Lock()
	defer h.Unlock()
	h.eating--
}

func (p Philosopher) eat(host *Host, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		for {
			if host.requestPermission() {
				break
			}
			time.Sleep(time.Millisecond)
		}

		if rand.Float32() < 0.5 {
			p.leftChopstick.Lock()
			p.rightChopstick.Lock()
		} else {
			p.rightChopstick.Lock()
			p.leftChopstick.Lock()
		}

		fmt.Printf("starting to eat %d\n", p.number)
		time.Sleep(time.Second)
		fmt.Printf("finishing eating %d\n", p.number)

		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()

		host.finishEating()
	}
	wg.Done()
}

func main() {
	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			number:         i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%5],
		}
	}

	host := &Host{}
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go philosophers[i].eat(host, &wg)
	}

	wg.Wait()
}
