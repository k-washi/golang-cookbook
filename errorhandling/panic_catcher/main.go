package main

import (
	"fmt"

	"github.com/k-washi/golang-cookbook/errorhandling/panic_catcher/panic_catcher"
)

/*
before
panic occurred: runtime error: integer divide by zero
after
*/

func main() {
	fmt.Println("before")
	panic_catcher.Cathcer()
	fmt.Println("after")
}
