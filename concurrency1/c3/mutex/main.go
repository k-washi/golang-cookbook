package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func parallel(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()

	fmt.Println("H")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("O")
	time.Sleep(100 * time.Millisecond)
	m.Unlock() //ここまで排他制御

	fmt.Println("S")
	time.Sleep(100 * time.Millisecond)

	fmt.Println("K")
	wg.Done()

}

var once = new(sync.Once)

func greeting(wg *sync.WaitGroup) {
	once.Do(func() {
		fmt.Println("Hello")
	})
	fmt.Println("How are you??")
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go parallel(wg, m)
	}
	wg.Wait()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go greeting(wg)
	}
	wg.Wait()
}
