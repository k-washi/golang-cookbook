package c1

import (
	"fmt"
	"time"
)

func Spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Println("sp", r)
			time.Sleep(delay)
		}
	}
}

//Fib slow version
func Fib(x int) int {
	if x < 2 {
		return x
	}
	return Fib(x-1) + Fib(x-2)
}
