package main

import (
	"fmt"

	"github.com/k-washi/golang-cookbook/datatype/collections/collections"
)

func main() {
	ws := []collections.WorkWith{
		collections.WorkWith{"Data1", 1},
		collections.WorkWith{"Data2", 2},
	}

	fmt.Printf("init list: %#v\n", ws)

	ws = collections.Map(ws, collections.LowerCaseData)

	fmt.Printf("init list: %#v\n", ws)

	ws = collections.Map(ws, collections.IncreamentVersion)

	fmt.Printf("init list: %#v\n", ws)

	ws = collections.Filter(ws, collections.OldVersion(3))

	fmt.Printf("init list: %#v\n", ws)

}
