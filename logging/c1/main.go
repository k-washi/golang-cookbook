package main

import (
	"errors"

	"github.com/k-washi/golang-cookbook/logging/c1/c1/logger"
)

func main() {
	logger.Info("test info msg")

	err := errors.New("test err msg")
	logger.Error("err msg", err)

}
