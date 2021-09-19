package log

import (
	"context"
)

//nolint:gochecknoglobals
var global = With(New(), WithLevel(LevelDebug))

// SetLogger sets global used logger. This function is not thread-safe.
func SetLogger(l Logger) {
	global = l
}

// GetLogger return global logger.
func GetLogger() Logger {
	return global
}

// Emerg log by emergency level.
func Emerg(ctx context.Context, args ...interface{}) {
	global.Emerg(ctx, args...)
}

// Alert log by alert level.
func Alert(ctx context.Context, args ...interface{}) {
	global.Alert(ctx, args...)
}

// Crit log by critical level.
func Crit(ctx context.Context, args ...interface{}) {
	global.Crit(ctx, args...)
}

// Err log by error level.
func Err(ctx context.Context, args ...interface{}) {
	global.Err(ctx, args...)
}

// Warn logs by warning level.
func Warn(ctx context.Context, args ...interface{}) {
	global.Warn(ctx, args...)
}

// Notice log by notice level.
func Notice(ctx context.Context, args ...interface{}) {
	global.Notice(ctx, args...)
}

// Info log by info level.
func Info(ctx context.Context, args ...interface{}) {
	global.Info(ctx, args...)
}

// Debug log by debug level.
func Debug(ctx context.Context, args ...interface{}) {
	global.Debug(ctx, args...)
}

// Print log by info level and arguments.
func Print(args ...interface{}) {
	global.Print(args...)
}

// Fatal log by alert level and arguments.
func Fatal(args ...interface{}) {
	global.Fatal(args...)
}

// Panic log by emergency level and arguments.
func Panic(args ...interface{}) {
	global.Panic(args...)
}

// Println log by info level and arguments.
func Println(args ...interface{}) {
	global.Println(args...)
}

// Fatalln log by alert level and arguments.
func Fatalln(args ...interface{}) {
	global.Fatalln(args...)
}

// Panicln log by emergency level and arguments.
func Panicln(args ...interface{}) {
	global.Panicln(args...)
}

// EmergKV log by emergency level and key-values.
func EmergKV(ctx context.Context, msg string, args ...interface{}) {
	global.EmergKV(ctx, msg, args...)
}

// AlertKV log by alert level and key-values.
func AlertKV(ctx context.Context, msg string, args ...interface{}) {
	global.AlertKV(ctx, msg, args...)
}

// CritKV log by critcal level and key-values.
func CritKV(ctx context.Context, msg string, args ...interface{}) {
	global.CritKV(ctx, msg, args...)
}

// ErrKV log by error level and key-values.
func ErrKV(ctx context.Context, msg string, args ...interface{}) {
	global.ErrKV(ctx, msg, args...)
}

// WarnKV log by warning level and key-values.
func WarnKV(ctx context.Context, msg string, args ...interface{}) {
	global.WarnKV(ctx, msg, args...)
}

// NoticeKV log by notice level and key-values.
func NoticeKV(ctx context.Context, msg string, args ...interface{}) {
	global.NoticeKV(ctx, msg, args...)
}

// InfoKV log by info level and key-values.
func InfoKV(ctx context.Context, msg string, args ...interface{}) {
	global.InfoKV(ctx, msg, args...)
}

// DebugKV log by debug level and key-values.
func DebugKV(ctx context.Context, msg string, args ...interface{}) {
	global.DebugKV(ctx, msg, args...)
}

// Emergf log by emergency level by format and arguments.
func Emergf(ctx context.Context, format string, args ...interface{}) {
	global.Emergf(ctx, format, args...)
}

// Alertf log by alert level by format and arguments.
func Alertf(ctx context.Context, format string, args ...interface{}) {
	global.Alertf(ctx, format, args...)
}

// Critf log by critical level by format and arguments.
func Critf(ctx context.Context, format string, args ...interface{}) {
	global.Critf(ctx, format, args...)
}

// Errf log by error level by format and arguments.
func Errf(ctx context.Context, format string, args ...interface{}) {
	global.Errf(ctx, format, args...)
}

// Warnf log by warning level by format and arguments.
func Warnf(ctx context.Context, format string, args ...interface{}) {
	global.Warnf(ctx, format, args...)
}

// Noticef log by notice level by format and arguments.
func Noticef(ctx context.Context, format string, args ...interface{}) {
	global.Noticef(ctx, format, args...)
}

// Infof log by info level by format and arguments.
func Infof(ctx context.Context, format string, args ...interface{}) {
	global.Noticef(ctx, format, args...)
}

// Debugf log by debug level by format and arguments.
func Debugf(ctx context.Context, format string, args ...interface{}) {
	global.Debugf(ctx, format, args...)
}

// Printf log by info level by format and arguments without context.
func Printf(format string, args ...interface{}) {
	global.Printf(format, args...)
}

// Fatalf log by alert level by format and arguments without context.
func Fatalf(format string, args ...interface{}) {
	global.Fatalf(format, args...)
}

// Panicf log by emergency level and arguments without context.
func Panicf(format string, args ...interface{}) {
	global.Panicf(format, args...)
}
