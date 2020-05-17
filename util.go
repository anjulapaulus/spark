package spark

import (
	"context"
	. "github.com/logrusorgru/aurora"
	"log"
	"github.com/google/uuid"
	"runtime"
)

const (
	FATAL string = `FATAL`
	ERROR string = `ERROR`
	WARN  string = `WARN`
	INFO  string = `INFO`
	DEBUG string = `DEBUG`
	TRACE string = `TRACE`
)


var logColors = map[string]string{
	FATAL: BgRed(`[FATAL]`).String(),
	ERROR: BgRed(`[ERROR]`).String(),
	WARN:  BgYellow(`[WARN]`).String(),
	INFO:  BgBlue(`[INFO]`).String(),
	DEBUG: BgCyan(`[DEBUG]`).String(),
	TRACE: BgMagenta(`[TRACE]`).String(),
}

var logTypes = map[string]int{
	FATAL: 0,
	ERROR: 1,
	WARN:  2,
	INFO:  3,
	DEBUG: 4,
	TRACE: 5,
}



type logChecker struct {
	*loggerConfig
	log *log.Logger
}

//checkLoggableType function checks whether logging is enabled under current configurations
func (l *logChecker) checkLoggableType (level string) bool {
	return logTypes[level] <= logTypes[l.logLevel]
}

func (l *logChecker) needColored (level string) string {
	if l.colors {
		return logColors[level]
	}

	return level
}


func getUUIDFromContext(ctx context.Context) uuid.UUID {
	uid := ctx.Value(`uuid`).(uuid.UUID)
	if uid == uuid.Nil {
		return uuid.New()
	}

	return uid
}

func (l *logChecker) logLine (level string, ctx context.Context, message interface{}, prms ...interface{}) {

	//check whether Log type exists
	if !l.checkLoggableType(level) {
		return
	}

	format := "%s [%s] [%+v]"

	var params []interface{}
	logLevel := level

	if l.colors {
		logLevel = l.needColored(level)
	}

	var uid uuid.UUID
	if ctx != nil {
		uid = getUUIDFromContext(ctx)
	} else {
		uid = uuid.New()
	}

	params = append(params, logLevel, uid.String(), message)

	if l.filePath {
		_, file, line, ok := runtime.Caller(l.fileDepth)
		if !ok {
			file = `<Unknown>`
			line = 1
		}

		format = "%s [%s] [%+v on %s:%d]"

		params = append(params, file, line)

	}

	if len(prms) > 0 {
		format = "%s [%s] [%+v on %s:%d] %+v"
		params = append(params, prms)
	}

	if level == FATAL {
		l.log.Fatalf(format, params...)
	}

	l.log.Printf(format, params...)
}
