package spark

import "fmt"

type loggerConfig struct {
	prefix    string
	colors    bool
	logLevel  string
	filePath  bool
	fileDepth int
	util      logChecker
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
	l.util.logLine(INFO, nil, v, `INFO`)
}

func (l *loggerConfig) Printf(format string, v ...interface{}) {
	l.util.logLine(INFO, nil, fmt.Sprintf(format, v...), `INFO`)
}

func (l *loggerConfig) Println(v ...interface{}) {
	l.util.logLine(INFO, nil, v, `INFO`)
}