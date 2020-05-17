package spark



type loggerConfig struct {
	prefix    string
	colors    bool
	logLevel  string
	filePath  bool
	fileDepth int
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