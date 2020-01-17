package basicerrors

import "fmt"

type CustomError struct {
	Result string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("err occured; %s", c.Result)
}
