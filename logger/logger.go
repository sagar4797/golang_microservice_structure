package logger

import "github.com/sirupsen/logrus"

type Logger struct {
	Service   string
	Version   string
	Formatter logrus.Formatter
}

func (l Logger) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Data["Service"] = l.Service
	entry.Data["Version"] = l.Version
	return l.Formatter.Format(entry)
}
