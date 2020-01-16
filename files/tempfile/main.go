package main

import (
	"github.com/k-washi/golang-cookbook/files/tempfile/tempfiles"
)

/*
/var/folders/1f/ydjcyxpx29358tslscc5z6_w0000gn/T/tmp567302347
ls /var/folders/1f/ydjcyxpx29358tslscc5z6_w0000gn/T/tmp567302347
No such file or directory
*/
func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
