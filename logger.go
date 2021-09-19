package log

import (
	"context"
	"fmt"
)

// Logger logged message.
type Logger func(ctx context.Context, level Level, msg string, fields Fields)

func (l Logger) log(ctx context.Context, level Level, args ...interface{}) {
	l(ctx, level, fmt.Sprint(args...), nil)
}

func (l Logger) logKV(ctx context.Context, level Level, msg string, args ...interface{}) {
	l(ctx, level, msg, l.kv(ctx, args...))
}

func (l Logger) logf(ctx context.Context, level Level, format string, args ...interface{}) {
	l(ctx, level, fmt.Sprintf(format, args...), nil)
}

func (l Logger) logln(ctx context.Context, level Level, args ...interface{}) {
	l(ctx, level, fmt.Sprintln(args...), nil)
}

func (l Logger) kv(ctx context.Context, args ...interface{}) []Field {
	fields := make([]Field, 0, len(args))

	for i := 0; i < len(args); i++ {
		if f, ok := args[i].(Field); ok {
			fields = append(fields, f)
			continue
		}

		if i == len(args)-1 {
			l(ctx, LevelCritical, fmt.Sprint("Ignored key without a value.", args[i]), fields)
			break
		}

		i++

		key, val := args[i-1], args[i]
		if keyStr, ok := key.(string); ok {
			fields = append(fields, Field{Key: keyStr, Value: val})
			continue
		}

		l(ctx, LevelCritical, fmt.Sprint("Ignored key-value pairs with non-string keys.", key, val), fields)
	}

	return fields
}

// With adds middlewares to logger.
func (l Logger) With(mw ...Middleware) Logger {
	return With(l, mw...)
}

// Emerg log by emergency level.
func (l Logger) Emerg(ctx context.Context, args ...interface{}) {
	l.log(ctx, LevelEmergency, args...)
}

// Alert log by alert level.
func (l Logger) Alert(ctx context.Context, args ...interface{}) {
	l.log(ctx, LevelAlert, args...)
}

// Crit log by critical level.
func (l Logger) Crit(ctx context.Context, args ...interface{}) {
	l.log(ctx, LevelCritical, args...)
}

// Err log by error level.
func (l Logger) Err(ctx context.Context, args ...interface{}) {
	l.log(ctx, LevelError, args...)
}

// Warn log by warning level.
func (l Logger) Warn(ctx context.Context, args ...interface{}) {
	l.log(ctx, LevelWarning, args...)
}

// Notice log by notice level.
func (l Logger) Notice(ctx context.Context, args ...interface{}) {
	l.log(ctx, LevelNotice, args...)
}

// Info log by info level.
func (l Logger) Info(ctx context.Context, args ...interface{}) {
	l.log(ctx, LevelInfo, args...)
}

// Debug log by debug level.
func (l Logger) Debug(ctx context.Context, args ...interface{}) {
	l.log(ctx, LevelDebug, args...)
}

// Print log by info level and arguments.
func (l Logger) Print(args ...interface{}) {
	l.log(context.Background(), LevelInfo, args...)
}

// Fatal log by alert level and arguments.
func (l Logger) Fatal(args ...interface{}) {
	l.log(context.Background(), LevelAlert, args...)
}

// Panic log by emergency level and arguments.
func (l Logger) Panic(args ...interface{}) {
	l.log(context.Background(), LevelEmergency, args...)
}

// Println log by info level and arguments.
func (l Logger) Println(args ...interface{}) {
	l.logln(context.Background(), LevelInfo, args...)
}

// Fatalln log by alert level and arguments.
func (l Logger) Fatalln(args ...interface{}) {
	l.logln(context.Background(), LevelAlert, args...)
}

// Panicln log by emergency level and arguments.
func (l Logger) Panicln(args ...interface{}) {
	l.logln(context.Background(), LevelEmergency, args...)
}

// EmergKV log by emergency level and key-values.
func (l Logger) EmergKV(ctx context.Context, msg string, args ...interface{}) {
	l.logKV(ctx, LevelEmergency, msg, args...)
}

// AlertKV log by alert level and key-values.
func (l Logger) AlertKV(ctx context.Context, msg string, args ...interface{}) {
	l.logKV(ctx, LevelAlert, msg, args...)
}

// CritKV log by critcal level and key-values.
func (l Logger) CritKV(ctx context.Context, msg string, args ...interface{}) {
	l.logKV(ctx, LevelCritical, msg, args...)
}

// ErrKV log by error level and key-values.
func (l Logger) ErrKV(ctx context.Context, msg string, args ...interface{}) {
	l.logKV(ctx, LevelError, msg, args...)
}

// WarnKV log by warning level and key-values.
func (l Logger) WarnKV(ctx context.Context, msg string, args ...interface{}) {
	l.logKV(ctx, LevelWarning, msg, args...)
}

// NoticeKV log by notice level and key-values.
func (l Logger) NoticeKV(ctx context.Context, msg string, args ...interface{}) {
	l.logKV(ctx, LevelNotice, msg, args...)
}

// InfoKV log by info level and key-values.
func (l Logger) InfoKV(ctx context.Context, msg string, args ...interface{}) {
	l.logKV(ctx, LevelInfo, msg, args...)
}

// DebugKV log by debug level and key-values.
func (l Logger) DebugKV(ctx context.Context, msg string, args ...interface{}) {
	l.logKV(ctx, LevelDebug, msg, args...)
}

// Emergf log by emergency level by format and arguments.
func (l Logger) Emergf(ctx context.Context, format string, args ...interface{}) {
	l.logf(ctx, LevelEmergency, format, args...)
}

// Alertf log by alert level by format and arguments.
func (l Logger) Alertf(ctx context.Context, format string, args ...interface{}) {
	l.logf(ctx, LevelAlert, format, args...)
}

// Critf log by critical level by format and arguments.
func (l Logger) Critf(ctx context.Context, format string, args ...interface{}) {
	l.logf(ctx, LevelCritical, format, args...)
}

// Errf log by error level by format and arguments.
func (l Logger) Errf(ctx context.Context, format string, args ...interface{}) {
	l.logf(ctx, LevelError, format, args...)
}

// Warnf log by warning level by format and arguments.
func (l Logger) Warnf(ctx context.Context, format string, args ...interface{}) {
	l.logf(ctx, LevelWarning, format, args...)
}

// Noticef log by notice level by format and arguments.
func (l Logger) Noticef(ctx context.Context, format string, args ...interface{}) {
	l.logf(ctx, LevelNotice, format, args...)
}

// Infof log by info level by format and arguments.
func (l Logger) Infof(ctx context.Context, format string, args ...interface{}) {
	l.logf(ctx, LevelInfo, format, args...)
}

// Debugf log by debug level by format and arguments.
func (l Logger) Debugf(ctx context.Context, format string, args ...interface{}) {
	l.logf(ctx, LevelDebug, format, args...)
}

// Printf log by info level by format and arguments without context.
func (l Logger) Printf(format string, args ...interface{}) {
	l.logf(context.Background(), LevelInfo, format, args...)
}

// Fatalf log by alert level by format and arguments without context.
func (l Logger) Fatalf(format string, args ...interface{}) {
	l.logf(context.Background(), LevelAlert, format, args...)
}

// Panicf log by emergency level and arguments without context.
func (l Logger) Panicf(format string, args ...interface{}) {
	l.logf(context.Background(), LevelEmergency, format, args...)
}
