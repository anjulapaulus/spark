package spark

import (
	"fmt"
	"log"
)

type loggerConfig struct {
	colors    bool
	logLevel  string
	filePath  bool
}

func NewLogger(level string,colorEnabled,filePath bool) *loggerConfig{
	return &loggerConfig{
		colors:    colorEnabled,
		logLevel:  level,
		filePath:  filePath,
	}
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

func doBatch(datas []interface{}) bool {

	for _, data := range datas {
		if parsedValue, ok := data.(int); ok {
			log.Println(fmt.Sprintf("[data]: %d", parsedValue))
		}
	}
	return true
}