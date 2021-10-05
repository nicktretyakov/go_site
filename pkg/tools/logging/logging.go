package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

// InitLogger init the logrus.Logger for defined log level
func InitLogger(path, level string) (*logrus.Logger, error) {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	logger := &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.JSONFormatter{
			CallerPrettyfier: caller,
			TimestampFormat:  "2006-01-02 15:04:05",
		},
		ReportCaller: true,
		Level:        logLevel,
	}

	logger.SetOutput(io.MultiWriter(os.Stdout, f))
	return logger, nil
}

// caller returns string presentation of log caller which is formatted as
func caller(f *runtime.Frame) (function string, file string) {
	_, filename := path.Split(f.File)
	return "", fmt.Sprintf("%s:%d", filename, f.Line)
}
