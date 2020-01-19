package main

import (
	"fmt"
	"sync"
	"time"
)

func hello1(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Println("Hello from :", id)
}

func main() {
	var wg sync.WaitGroup

	//WaitGroup test1
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st routine")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd routine")
		time.Sleep(1)
	}()

	wg.Wait()
	fmt.Println("comp")

	fmt.Println()
	//WaitGroup test2
	const numG = 5
	wg.Add(numG)
	for i := 0; i < numG; i++ {
		go hello1(&wg, i+1)
	}
	wg.Wait()
	fmt.Println()

	//Mutex
	var muCount int
	var muLock sync.Mutex

	increment := func() {
		muLock.Lock()
		defer muLock.Unlock()
		muCount++
		fmt.Println("Increment:", muCount)
	}

	decrement := func() {
		muLock.Lock()
		defer muLock.Unlock()
		muCount--
		fmt.Println("Decrement:", muCount)
	}

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement()
		}()
	}
	wg.Wait()
	fmt.Println()

	//Cond

	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Remved Queue")
		c.L.Unlock()
		c.Signal() //queueに変更があったことを伝えている。
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Add Queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Millisecond)
		c.L.Unlock()
	}
	fmt.Println()

	//Broadcast
	type Button struct {
		Clicked *sync.Cond
	}

	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}
	subscribe := func(c *sync.Cond, fn func()) {
		var wgGoRun sync.WaitGroup
		wgGoRun.Add(1)
		go func() {
			wgGoRun.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		wgGoRun.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)

	subscribe(button.Clicked, func() {
		fmt.Println("C1")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("C2")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("C3")
		clickRegistered.Done()
	})
	button.Clicked.Broadcast()
	clickRegistered.Wait()

	//Once
	var onceCount = 0
	incriment := func() {
		onceCount++
	}
}
