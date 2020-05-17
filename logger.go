package spark

import (
	"fmt"
	log "log"
)

type loggerConfig struct {
	prefix    string
	colors    bool
	logLevel  string
	filePath  bool
	fileDepth int
	log       *log.Logger
}

func NewLogger(level string,colorEnabled,filePath bool) *loggerConfig{
	return &loggerConfig{
		colors:    colorEnabled,
		logLevel:  level,
		filePath:  filePath,
	}
}

type SimpleLoggerInterface interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

func (l *loggerConfig) Print(v ...interface{}) {
	l.logLine(INFO, nil, v, `INFO`)
}

func (l *loggerConfig) Printf(format string, v ...interface{}) {
	l.logLine(INFO, nil, fmt.Sprintf(format, v...), `INFO`)
}

func (l *loggerConfig) Println(v ...interface{}) {
	l.logLine(INFO, nil, v, `INFO`)
}