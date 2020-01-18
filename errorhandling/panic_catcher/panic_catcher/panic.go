package panic_catcher

import (
	"fmt"
	"strconv"
)

func Panic() {
	zero, err := strconv.ParseInt("0", 10, 64)
	if err != nil {
		panic(err)
	}

	a := 1 / zero
	fmt.Println("never get", a)

}

func Cathcer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic occurred:", r)
		}

	}()
	Panic()
}
