package spark

import (
	"context"
	"github.com/google/uuid"
	. "github.com/logrusorgru/aurora"
	"log"
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



//checkLoggableType function checks whether logging is enabled under current configurations
func (l *loggerConfig) checkLoggableType (level string) bool {
	return logTypes[level] <= logTypes[l.logLevel]
}

func (l *loggerConfig) needColored (level string) string {
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

func (l *loggerConfig) logLine (level string, ctx context.Context, message interface{}, prms ...interface{}) {

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

	if len(prms) > 0 {
		format = "%s [%s] [%+v] %+v"
		params = append(params, prms)
	}

	if level == FATAL {
		log.Fatalf(format, params...)
	}

	log.Printf(format, params...)
}
