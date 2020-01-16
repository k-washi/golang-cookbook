package main

import (
	"github.com/k-washi/golang-cookbook/io/string_01/bytestrings"
)

/*
macpc-pro:io washizakikai$ go run ./string_01/main.go
it's easy to encode unicode into a byte array
it's easy to encode unicode into a byte array
it'seasytoencodeunicodeintoabytearraytrue
true
true
true
[this is a test]
This Is A Test
this is a test
*/

func main() {
	err := bytestrings.WorkWithBuffer()
	if err != nil {
		panic(err)
	}

	bytestrings.SearchString()

}
