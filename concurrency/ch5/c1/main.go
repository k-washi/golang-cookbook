package main

import (
	"fmt"
	"sync"
)

/*
Recover from runtime error
...
x/(x - 5):  6
x/(x - 5):  2
Recovered:  runtime error: integer divide by zero
x/(x - 5):  -4
...
*/

func simpleRecover(ind int, wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()

	defer wg.Done()

	fmt.Println("x/(x - 5): ", ind/(ind-5))
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(j int) {
			simpleRecover(j, &wg)
		}(i)
	}

	wg.Wait()
}
