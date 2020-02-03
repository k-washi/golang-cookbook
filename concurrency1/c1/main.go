package main

import (
	"fmt"
	"time"

	"github.com/k-washi/golang-cookbook/concurrency1/c1/c1"
)

func main() {
	go c1.Spinner(1 * time.Second)
	fibN := c1.Fib(45)
	fmt.Println(fibN)
}
