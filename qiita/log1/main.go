package main

import (
	"errors"
	"io"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	log     *logrus.Logger
	initLog sync.Once
)

//LoggerInit ロギングの初期化
func LoggerInit() error {

	err := errors.New("already initialized")
	initLog.Do(func() {
		err = nil
		log = logrus.New()
		log.Formatter = &prefixed.TextFormatter{
			ForceColors:     true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		}
		f, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0777)
		log.Out = io.MultiWriter(os.Stdout, f)
		log.Level = logrus.DebugLevel
	})
	return err
}

func main() {
	/*
		[90m[2020-01-20 00:22:12][0m [32m INFO[0m log info
		[90m[2020-01-20 00:22:12][0m [34mDEBUG[0m log debug
		[90m[2020-01-20 00:22:12][0m [31mERROR[0m log error
		0m [34mDEBUG[0m hello [34mkey[0m=value
		[90m[2020-01-18 23:14:53][0m [34mDEBUG[0m[36m test test2:[0m
	*/
	if err := LoggerInit(); err != nil {
		panic(err)
	}

	log.Info("log info")
	log.Debug("log debug")
	log.Error("log error")

	if err := LoggerInit(); err != nil {
		panic(err)
	}
}
