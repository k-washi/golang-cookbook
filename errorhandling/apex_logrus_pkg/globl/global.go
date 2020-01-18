package globl

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

func Init() error {
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

func SetLog(l *logrus.Logger) {
	log = l
}

func WithField(key string, value interface{}) *logrus.Entry {
	return log.WithField(key, value)
}

func Debug(args ...interface{}) {
	log.Debug(args)
}
