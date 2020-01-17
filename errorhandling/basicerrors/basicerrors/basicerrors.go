package basicerrors

import (
	"errors"
	"fmt"
)

var ErrorV = errors.New("typed error")

type TypedError struct {
	error
}

func BasicErrors() {
	err := errors.New("easy way to create an error")
	fmt.Println("errors.New:", err)

	err = fmt.Errorf("an error occurred: %s", "Something")
	fmt.Println("fmt.Errorf:", err)

	err = ErrorV
	fmt.Println("value error:", err)

	err = TypedError{errors.New("typed error")}
	fmt.Println("typed error:", err)
}
