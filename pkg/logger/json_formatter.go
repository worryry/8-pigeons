package logger

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	timestampFormat = "2006-01-02 15:04:05.000"
	defaultFiledNum = 4
)

type FastJsonFormatter struct{}

func (f *FastJsonFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, defaultFiledNum)
	data["level"] = entry.Level
	data["time"] = time.Now().Format(timestampFormat)
	data["msg"] = entry.Message
	data["fields"] = entry.Data

	byteArr, err := jsoniter.Marshal(data)
	if err == nil {
		byteArr = append(byteArr, '\n')
	}

	return byteArr, err
}
