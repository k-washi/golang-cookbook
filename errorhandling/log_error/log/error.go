package log

import (
	"log"

	"github.com/pkg/errors"
)

func OriginalError() error {
	return errors.New("error occurred")
}

func PassThroughError() error {
	err := OriginalError()

	return errors.Wrap(err, "passthrough Error")
}

func FinalDestination() {
	err := PassThroughError()
	if err != nil {
		log.Printf("error ocurred: %s\n", err.Error())
		return
	}
	return
}
