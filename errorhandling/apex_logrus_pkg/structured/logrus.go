package structured

import (
	"github.com/sirupsen/logrus"
)

type Hook struct {
	id string
}

func (hock *Hook) Fire(entry *logrus.Entry) error {
	entry.Data["id"] = hock.id
	return nil
}

func (hock *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func Logrus() {
	/*
		WARN[0000] Warning                                       complex_struct="{Something happened Just now}" id=123 success=true
		ERRO[0000] error                                         complex_struct="{Something happened Just now}" id=123 success=true
	*/
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.AddHook(&Hook{"123"})

	fields := logrus.Fields{}
	fields["success"] = true
	fields["complex_struct"] = struct {
		Event string
		When  string
	}{"Something happened", "Just now"}

	x := logrus.WithFields(fields)
	x.Warn("Warning")
	x.Error("error")
}
