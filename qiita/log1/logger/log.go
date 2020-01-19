package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

/*
https://github.com/x-cray/logrus-prefixed-formatter/blob/master/formatter.go
*/

type LogFormat struct {
	TimestampFormat string
	file            *runtime.Frame
}

func (f *LogFormat) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	b.WriteByte('[')
	b.WriteString(strings.ToUpper(entry.Level.String()))
	b.WriteString("]:")
	b.WriteString(entry.Time.Format(f.TimestampFormat))

	b.WriteString(" [")
	b.WriteString(formatFilePath(entry.Caller.File))
	b.WriteString(":")
	fmt.Fprint(b, entry.Caller.Line)
	b.WriteString("] ")

	if entry.Message != "" {
		b.WriteString(" - ")
		b.WriteString(entry.Message)
	}

	if len(entry.Data) > 0 {
		b.WriteString(" || ")
	}
	for key, value := range entry.Data {
		b.WriteString(key)
		b.WriteByte('=')
		b.WriteByte('{')
		fmt.Fprint(b, value)
		b.WriteString("}, ")
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func init() {
	logrus.SetReportCaller(true)
	formatter := LogFormat{}
	formatter.TimestampFormat = "2006-01-02 15:04:05"

	logrus.SetFormatter(&formatter)

	f, err := openFile("log.txt")
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(io.MultiWriter(os.Stdout, f))

}

func openFile(fileName string) (*os.File, error) {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	return f, err
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
