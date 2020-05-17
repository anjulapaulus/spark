package spark

import (
	"fmt"
)

type loggerConfig struct {
	colors    bool
	logLevel  string
	filePath  bool
	SimpleLogger simpleLogger
}

func NewLogger(level string,colorEnabled,filePath bool) *loggerConfig{
	return &loggerConfig{
		colors:    colorEnabled,
		logLevel:  level,
		filePath:  filePath,
	}
}

type simpleLogger struct {
	*loggerConfig
}

func (l *simpleLogger) Print(v ...interface{}) {
	l.logLine(INFO, nil, v, `INFO`)
}

func (l *simpleLogger) Printf(format string, v ...interface{}) {
	l.logLine(INFO, nil, fmt.Sprintf(format, v...), `INFO`)
}

func (l *simpleLogger) Println(v ...interface{}) {
	l.logLine(INFO, nil, v, `INFO`)
}