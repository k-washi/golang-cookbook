package main

import (
	"fmt"

	"github.com/k-washi/golang-cookbook/errorhandling/apex_logrus_pkg/globl"

	"github.com/k-washi/golang-cookbook/errorhandling/apex_logrus_pkg/structured"
)

func main() {
	structured.Logrus()

	fmt.Println()

	//structured.Apex()
	if err := globl.Init(); err != nil {
		panic(err)
	}

	if err := globl.UserLog(); err != nil {
		panic(err)
	}
	if err := globl.UserLog(); err != nil {
		panic(err)
	}

}
