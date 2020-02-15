package main

import (
	"fmt"
	"time"

	"github.com/k-washi/golang-cookbook/concurrency1/c1/c1"
)

func main() {
	//f1()
	f2()
}

func f1() {
	go c1.Spinner(1 * time.Second)
	fibN := c1.Fib(45)
	fmt.Println(fibN)
}

func f2() {
	c1.StartServer()
}
