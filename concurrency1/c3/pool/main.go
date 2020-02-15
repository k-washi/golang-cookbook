package main

import (
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

/*
常に行う動作に、非定常動作を割り込む。
ABC
ABC
insert
ABC
ABC
insert
ABC
ABC
insert
insert
*/

func main() {
	defer debug.SetGCPercent(debug.SetGCPercent(-1))

	p := sync.Pool{
		New: func() interface{} {
			return "ABC"
		},
	}

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			p.Put("insert")
			time.Sleep(100 * time.Millisecond)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(p.Get())
			time.Sleep(50 * time.Millisecond)
		}
		wg.Done()
	}()
	wg.Wait()

}
