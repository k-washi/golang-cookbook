package main

import (
	"fmt"
	"sync"
)

func main() {
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
}
