package spark

import (
	"context"
	"fmt"
)

type LoggerInterface interface {
	Fatal(message interface{}, params ...interface{})
	Error(message interface{}, params ...interface{})
	Warn(message interface{}, params ...interface{})
	Debug(message interface{}, params ...interface{})
	Info(message interface{}, params ...interface{})
	Trace(message interface{}, params ...interface{})
	FatalContext(ctx context.Context, message interface{}, params ...interface{})
	ErrorContext(ctx context.Context, message interface{}, params ...interface{})
	WarnContext(ctx context.Context, message interface{}, params ...interface{})
	DebugContext(ctx context.Context, message interface{}, params ...interface{})
	InfoContext(ctx context.Context, message interface{}, params ...interface{})
	TraceContext(ctx context.Context, message interface{}, params ...interface{})
	SimpleLoggerInterface
}

type logger struct {
	logChecker
}

func (l *logger) Fatal(message interface{}, params ...interface{}) {
	l.logLine(FATAL, nil, message, params...)
}

func (l *logger) Error(message interface{}, params ...interface{}) {
	l.logLine(ERROR, nil, message, params...)
}

func (l *logger) Warn(message interface{}, params ...interface{}) {
	l.logLine(WARN, nil, message, params...)
}

func (l *logger) Info(message interface{}, params ...interface{}) {
	l.logLine(INFO, nil, message, params...)
}

func (l *logger) Debug(message interface{}, params ...interface{}) {
	l.logLine(DEBUG, nil, message, params...)
}

func (l *logger) Trace(message interface{}, params ...interface{}) {
	l.logLine(TRACE, nil, message, params...)
}

func (l *logger) ErrorContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logLine(ERROR, ctx, message, params...)
}

func (l *logger) WarnContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logLine(WARN, ctx, message, params...)
}

func (l *logger) InfoContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logLine(INFO, ctx, message, params...)
}

func (l *logger) DebugContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logLine(DEBUG, ctx, message, params...)
}

func (l *logger) TraceContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logLine(TRACE, ctx, message, params...)
}

func (l *logger) FatalContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logLine(FATAL, ctx, message, params...)
}

func (l *logger) Print(v ...interface{}) {
	l.logLine(INFO, nil, v, `INFO`)
}

func (l *logger) Printf(format string, v ...interface{}) {
	l.logLine(INFO, nil, fmt.Sprintf(format, v...), `INFO`)
}

func (l *logger) Println(v ...interface{}) {
	l.logLine(INFO, nil, v, `INFO`)
}
