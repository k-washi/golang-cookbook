package main

import (
	"bytes"
	"fmt"

	"github.com/k-washi/golang-cookbook/io/common_01/iointerface"
)

/*
stdout on Copy=example
out bytes buffer = example
*/

func main() {
	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Print("stdout on Copy=")
	if err := iointerface.Copy(in, out); err != nil {
		panic(err)
	}

	fmt.Println("out bytes buffer =", out.String())
}
