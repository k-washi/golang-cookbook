package main

import (
	"github.com/k-washi/golang-cookbook/files/directory_01/filedirs"
)

func main() {
	if err := filedirs.Operate(); err != nil {
		panic(err)
	}
}
