package spark

import (
	"context"
	"fmt"
)

type PrefixLogger interface {
	Fatal(prefix string, message interface{}, params ...interface{})
	Error(prefix string, message interface{}, params ...interface{})
	Warn(prefix string, message interface{}, params ...interface{})
	Debug(prefix string, message interface{}, params ...interface{})
	Info(prefix string, message interface{}, params ...interface{})
	Trace(prefix string, message interface{}, params ...interface{})
	FatalContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	ErrorContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	WarnContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	DebugContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	InfoContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	TraceContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	SimpleLoggerInterface
}

func WithPrefix(p string, message interface{}) string {
	return fmt.Sprintf(`%s] [%+v`, p, message)
}

type prefixedLogger struct {
	logChecker
}


func (l *prefixedLogger) Fatal(prefix string, message interface{}, params ...interface{}) {
	l.logLine(FATAL, nil, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Error(prefix string, message interface{}, params ...interface{}) {
	l.logLine(ERROR, nil, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Warn(prefix string, message interface{}, params ...interface{}) {
	l.logLine(WARN, nil, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Info(prefix string, message interface{}, params ...interface{}) {
	l.logLine(INFO, nil, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Debug(prefix string, message interface{}, params ...interface{}) {
	l.logLine(DEBUG, nil, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Trace(prefix string, message interface{}, params ...interface{}) {
	l.logLine(TRACE, nil, WithPrefix(prefix, message), params...)
}


func (l *prefixedLogger) FatalContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logLine(FATAL, ctx, WithPrefix(prefix, message), params)
}

func (l *prefixedLogger) ErrorContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logLine(ERROR, ctx, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) WarnContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logLine(WARN, ctx, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) InfoContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logLine(INFO, ctx, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) DebugContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logLine(DEBUG, ctx, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) TraceContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logLine(TRACE, ctx, WithPrefix(prefix, message), params...)
}

func (l *prefixedLogger) Print(v ...interface{}) {
	l.logLine(INFO, nil, v, l.needColored(`INFO`))
}

func (l *prefixedLogger) Printf(format string, v ...interface{}) {
	l.logLine(INFO, nil, fmt.Sprintf(format, v...), l.needColored(`INFO`))
}

func (l *prefixedLogger) Println(v ...interface{}) {
	l.logLine(INFO, nil, v, l.needColored(`INFO`))

}