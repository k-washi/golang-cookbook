package main

import (
	"bytes"
	"fmt"
	"os"
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

	//Pool pattern ex1
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create instance")
			return struct{}{}
		},
	}
	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()

	//Pool pattern ex2
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated++
			mem := make([]byte, 1024)
			return &mem
		},
	}

	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wgPool sync.WaitGroup
	wgPool.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wgPool.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)

		}()
	}

	wgPool.Wait()
	fmt.Println("Calc num:", numCalcsCreated)

	fmt.Println()
	//channel ex1
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello"
	}()

	fmt.Println(<-stringStream) //channelに値が格納されるまで待機。

	//channel ex2
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Println("channel stream :", integer)
	}

	//channel ex3
	//beginを閉じたら、次に進む
	begin := make(chan interface{})
	var wgCh3 sync.WaitGroup
	for i := 0; i < 4; i++ {
		wgCh3.Add(1)
		go func(i int) {
			defer wgCh3.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("unblocking goroutines ...")
	close(begin)
	wgCh3.Wait()

	//channel ex4
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream = make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Prod Done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "sending: %d\n", i)
			intStream <- i
		}
	}()
	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}

	//channel ex5
	fmt.Println()
	chanOwner := func() <-chan int {
		resStream := make(chan int, 5)
		go func() {
			defer close(resStream)
			for i := 0; i <= 5; i++ {
				resStream <- i
			}
		}()
		return resStream
	}

	resStream := chanOwner()
	for res := range resStream {
		fmt.Println("Received:", res)
	}
	fmt.Println("Ch Done")

	//select ex1
	start := time.Now()
	cselect1 := make(chan interface{})
	go func() {
		time.Sleep(50 * time.Millisecond)
		close(cselect1)
	}()

	fmt.Println("Blocking on read ...")
	select {
	case <-cselect1:
		fmt.Println("closed: ", time.Since(start))

	}

}
