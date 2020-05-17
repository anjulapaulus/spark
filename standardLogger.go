package spark

import (
	"context"
)


func (l *loggerConfig) Fatal(message interface{}, params ...interface{}) {
	l.logLine(FATAL, nil, message, params...)
}

func (l *loggerConfig) Error(message interface{}, params ...interface{}) {
	l.logLine(ERROR, nil, message, params...)
}

func (l *loggerConfig) Warn(message interface{}, params ...interface{}) {
	l.logLine(WARN, nil, message, params...)
}

func (l *loggerConfig) Info(message interface{}, params ...interface{}) {
	l.logLine(INFO, nil, message, params...)
}

func (l *loggerConfig) Debug(message interface{}, params ...interface{}) {
	l.logLine(DEBUG, nil, message, params...)
}

func (l *loggerConfig) Trace(message interface{}, params ...interface{}) {
	l.logLine(TRACE, nil, message, params...)
}

func (l *loggerConfig) ErrorContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logLine(ERROR, ctx, message, params...)
}

func (l *loggerConfig) WarnContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logLine(WARN, ctx, message, params...)
}

func (l *loggerConfig) InfoContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logLine(INFO, ctx, message, params...)
}

func (l *loggerConfig) DebugContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logLine(DEBUG, ctx, message, params...)
}

func (l *loggerConfig) TraceContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logLine(TRACE, ctx, message, params...)
}

func (l *loggerConfig) FatalContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logLine(FATAL, ctx, message, params...)
}
