package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
waiting 0
waiting 7
waiting 5
waiting 3
waiting 6
waiting 4
waiting 8
waiting 2
waiting 9
waiting 1
go 0
go 7
go 5
go 3
go 6
Brodcast start!!!
go 2
go 9
go 4
go 8
go 1
*/

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	l := new(sync.Mutex)
	c := sync.NewCond(l)

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("waiting %d\n", i)
			l.Lock()
			defer l.Unlock()
			c.Wait()
			fmt.Printf("go %d\n", i)

		}(i)
	}

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		c.Signal()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Brodcast start!!!")
	c.Broadcast()

	time.Sleep(1 * time.Second)

}
