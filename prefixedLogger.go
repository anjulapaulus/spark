package spark

import (
	"context"
	"fmt"
)

type PrefixLogger interface {
	PrefixedFatal(prefix string, message interface{}, params ...interface{})
	PrefixedError(prefix string, message interface{}, params ...interface{})
	PrefixedWarn(prefix string, message interface{}, params ...interface{})
	PrefixedDebug(prefix string, message interface{}, params ...interface{})
	PrefixedInfo(prefix string, message interface{}, params ...interface{})
	PrefixedTrace(prefix string, message interface{}, params ...interface{})
	PrefixedFatalContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	PrefixedErrorContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	PrefixedWarnContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	PrefixedDebugContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	PrefixedInfoContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	PrefixedTraceContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	SimpleLoggerInterface
}

func WithPrefix(p string, message interface{}) string {
	return fmt.Sprintf(`%s] [%+v`, p, message)
}



func (l *loggerConfig) PrefixedFatal(prefix string, message interface{}, params ...interface{}) {
	l.util.logLine(FATAL, nil, WithPrefix(prefix, message), params...)
}

func (l *loggerConfig) PrefixedError(prefix string, message interface{}, params ...interface{}) {
	l.util.logLine(ERROR, nil, WithPrefix(prefix, message), params...)
}

func (l *loggerConfig) PrefixedWarn(prefix string, message interface{}, params ...interface{}) {
	l.util.logLine(WARN, nil, WithPrefix(prefix, message), params...)
}

func (l *loggerConfig) PrefixedInfo(prefix string, message interface{}, params ...interface{}) {
	l.util.logLine(INFO, nil, WithPrefix(prefix, message), params...)
}

func (l *loggerConfig) PrefixedDebug(prefix string, message interface{}, params ...interface{}) {
	l.util.logLine(DEBUG, nil, WithPrefix(prefix, message), params...)
}

func (l *loggerConfig) PrefixedTrace(prefix string, message interface{}, params ...interface{}) {
	l.util.logLine(TRACE, nil, WithPrefix(prefix, message), params...)
}


func (l *loggerConfig) PrefixedFatalContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.util.logLine(FATAL, ctx, WithPrefix(prefix, message), params)
}

func (l *loggerConfig) PrefixedErrorContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.util.logLine(ERROR, ctx, WithPrefix(prefix, message), params...)
}

func (l *loggerConfig) PrefixedWarnContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.util.logLine(WARN, ctx, WithPrefix(prefix, message), params...)
}

func (l *loggerConfig) PrefixedInfoContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.util.logLine(INFO, ctx, WithPrefix(prefix, message), params...)
}

func (l *loggerConfig) PrefixedDebugContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.util.logLine(DEBUG, ctx, WithPrefix(prefix, message), params...)
}

func (l *loggerConfig) PrefixedTraceContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.util.logLine(TRACE, ctx, WithPrefix(prefix, message), params...)
}
