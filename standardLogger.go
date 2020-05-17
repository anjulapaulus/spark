package spark

import (
	"context"
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


func (l *loggerConfig) Fatal(message interface{}, params ...interface{}) {
	l.util.logLine(FATAL, nil, message, params...)
}

func (l *loggerConfig) Error(message interface{}, params ...interface{}) {
	l.util.logLine(ERROR, nil, message, params...)
}

func (l *loggerConfig) Warn(message interface{}, params ...interface{}) {
	l.util.logLine(WARN, nil, message, params...)
}

func (l *loggerConfig) Info(message interface{}, params ...interface{}) {
	l.util.logLine(INFO, nil, message, params...)
}

func (l *loggerConfig) Debug(message interface{}, params ...interface{}) {
	l.util.logLine(DEBUG, nil, message, params...)
}

func (l *loggerConfig) Trace(message interface{}, params ...interface{}) {
	l.util.logLine(TRACE, nil, message, params...)
}

func (l *loggerConfig) ErrorContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.util.logLine(ERROR, ctx, message, params...)
}

func (l *loggerConfig) WarnContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.util.logLine(WARN, ctx, message, params...)
}

func (l *loggerConfig) InfoContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.util.logLine(INFO, ctx, message, params...)
}

func (l *loggerConfig) DebugContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.util.logLine(DEBUG, ctx, message, params...)
}

func (l *loggerConfig) TraceContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.util.logLine(TRACE, ctx, message, params...)
}

func (l *loggerConfig) FatalContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.util.logLine(FATAL, ctx, message, params...)
}
