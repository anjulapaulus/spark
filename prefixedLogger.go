package spark

import (
	"context"
	"fmt"
)


func WithPrefix(p string, message interface{}) string {
	return fmt.Sprintf(`%s] [%+v`, p, message)
}


func (l *loggerConfig) PrefixedFatal(prefix string, message interface{}, params ...interface{}) {
	l.logLine(FATAL, nil, WithPrefix(prefix, message), params...)
}

func (l *loggerConfig) PrefixedError(prefix string, message interface{}, params ...interface{}) {
	l.logLine(ERROR, nil, WithPrefix(prefix, message), params...)
}

func (l *loggerConfig) PrefixedWarn(prefix string, message interface{}, params ...interface{}) {
	l.logLine(WARN, nil, WithPrefix(prefix, message), params...)
}

func (l *loggerConfig) PrefixedInfo(prefix string, message interface{}, params ...interface{}) {
	l.logLine(INFO, nil, WithPrefix(prefix, message), params...)
}

func (l *loggerConfig) PrefixedDebug(prefix string, message interface{}, params ...interface{}) {
	l.logLine(DEBUG, nil, WithPrefix(prefix, message), params...)
}

func (l *loggerConfig) PrefixedTrace(prefix string, message interface{}, params ...interface{}) {
	l.logLine(TRACE, nil, WithPrefix(prefix, message), params...)
}


func (l *loggerConfig) PrefixedFatalContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logLine(FATAL, ctx, WithPrefix(prefix, message), params)
}

func (l *loggerConfig) PrefixedErrorContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logLine(ERROR, ctx, WithPrefix(prefix, message), params...)
}

func (l *loggerConfig) PrefixedWarnContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logLine(WARN, ctx, WithPrefix(prefix, message), params...)
}

func (l *loggerConfig) PrefixedInfoContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logLine(INFO, ctx, WithPrefix(prefix, message), params...)
}

func (l *loggerConfig) PrefixedDebugContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logLine(DEBUG, ctx, WithPrefix(prefix, message), params...)
}

func (l *loggerConfig) PrefixedTraceContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.logLine(TRACE, ctx, WithPrefix(prefix, message), params...)
}
