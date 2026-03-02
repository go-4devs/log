package log

import (
	"context"
	"io"
	"time"

	"gitoa.ru/go-4devs/log/field"
	"gitoa.ru/go-4devs/log/level"
)

//nolint:gochecknoglobals
var global = With(New(),
	WithTime(KeyTime, time.RFC3339),
	WithLevel(KeyLevel, level.Debug),
	WithExit(level.Alert),
	WithPanic(level.Emergency),
)

// SetLogger sets global used logger. This function is not thread-safe.
func SetLogger(l Logger) {
	global = l
}

// Log return global logger.
func Log() Logger {
	return global
}

// Emerg log by emergency level.
func Emerg(ctx context.Context, args ...any) {
	Log().Emerg(ctx, args...)
}

// Alert log by alert level.
func Alert(ctx context.Context, args ...any) {
	Log().Alert(ctx, args...)
}

// Crit log by critical level.
func Crit(ctx context.Context, args ...any) {
	Log().Crit(ctx, args...)
}

// Err log by error level.
func Err(ctx context.Context, args ...any) {
	Log().Err(ctx, args...)
}

// Warn logs by warning level.
func Warn(ctx context.Context, args ...any) {
	Log().Warn(ctx, args...)
}

// Notice log by notice level.
func Notice(ctx context.Context, args ...any) {
	Log().Notice(ctx, args...)
}

// Info log by info level.
func Info(ctx context.Context, args ...any) {
	Log().Info(ctx, args...)
}

// Debug log by debug level.
func Debug(ctx context.Context, args ...any) {
	Log().Debug(ctx, args...)
}

// Print log by info level and arguments.
func Print(args ...any) {
	Log().Print(args...)
}

// Fatal log by alert level and arguments.
func Fatal(args ...any) {
	Log().Fatal(args...)
}

// Panic log by emergency level and arguments.
func Panic(args ...any) {
	Log().Panic(args...)
}

// Println log by info level and arguments.
func Println(args ...any) {
	Log().Println(args...)
}

// Fatalln log by alert level and arguments.
func Fatalln(args ...any) {
	Log().Fatalln(args...)
}

// Panicln log by emergency level and arguments.
func Panicln(args ...any) {
	Log().Panicln(args...)
}

// EmergKVs sugared log by emergency level and key-values.
func EmergKVs(ctx context.Context, msg string, args ...any) {
	Log().EmergKVs(ctx, msg, args...)
}

// AlertKVs sugared log by alert level and key-values.
func AlertKVs(ctx context.Context, msg string, args ...any) {
	Log().AlertKVs(ctx, msg, args...)
}

// CritKVs sugared log by critical level and key-values.
func CritKVs(ctx context.Context, msg string, args ...any) {
	Log().CritKVs(ctx, msg, args...)
}

// ErrKVs sugared log by error level and key-values.
func ErrKVs(ctx context.Context, msg string, args ...any) {
	Log().ErrKVs(ctx, msg, args...)
}

// WarnKVs sugared log by warning level and key-values.
func WarnKVs(ctx context.Context, msg string, args ...any) {
	Log().WarnKVs(ctx, msg, args...)
}

// NoticeKVs sugared log by notice level and key-values.
func NoticeKVs(ctx context.Context, msg string, args ...any) {
	Log().NoticeKVs(ctx, msg, args...)
}

// InfoKVs sugared log by info level and key-values.
func InfoKVs(ctx context.Context, msg string, args ...any) {
	Log().InfoKVs(ctx, msg, args...)
}

// DebugKVs sugared log by debug level and key-values.
func DebugKVs(ctx context.Context, msg string, args ...any) {
	Log().DebugKVs(ctx, msg, args...)
}

// EmergKV  log by emergency level and key-values.
func EmergKV(ctx context.Context, msg string, args ...field.Field) {
	Log().EmergKV(ctx, msg, args...)
}

// AlertKV log by alert level and key-values.
func AlertKV(ctx context.Context, msg string, args ...field.Field) {
	Log().AlertKV(ctx, msg, args...)
}

// CritKV log by critical level and key-values.
func CritKV(ctx context.Context, msg string, args ...field.Field) {
	Log().CritKV(ctx, msg, args...)
}

// ErrKV log by error level and key-values.
func ErrKV(ctx context.Context, msg string, args ...field.Field) {
	Log().ErrKV(ctx, msg, args...)
}

// WarnKV log by warning level and key-values.
func WarnKV(ctx context.Context, msg string, args ...field.Field) {
	Log().WarnKV(ctx, msg, args...)
}

// NoticeKV log by notice level and key-values.
func NoticeKV(ctx context.Context, msg string, args ...field.Field) {
	Log().NoticeKV(ctx, msg, args...)
}

// InfoKV log by info level and key-values.
func InfoKV(ctx context.Context, msg string, args ...field.Field) {
	Log().InfoKV(ctx, msg, args...)
}

// DebugKV log by debug level and key-values.
func DebugKV(ctx context.Context, msg string, args ...field.Field) {
	Log().DebugKV(ctx, msg, args...)
}

// Emergf log by emergency level by format and arguments.
func Emergf(ctx context.Context, format string, args ...any) {
	Log().Emergf(ctx, format, args...)
}

// Alertf log by alert level by format and arguments.
func Alertf(ctx context.Context, format string, args ...any) {
	Log().Alertf(ctx, format, args...)
}

// Critf log by critical level by format and arguments.
func Critf(ctx context.Context, format string, args ...any) {
	Log().Critf(ctx, format, args...)
}

// Errf log by error level by format and arguments.
func Errf(ctx context.Context, format string, args ...any) {
	Log().Errf(ctx, format, args...)
}

// Warnf log by warning level by format and arguments.
func Warnf(ctx context.Context, format string, args ...any) {
	Log().Warnf(ctx, format, args...)
}

// Noticef log by notice level by format and arguments.
func Noticef(ctx context.Context, format string, args ...any) {
	Log().Noticef(ctx, format, args...)
}

// Infof log by info level by format and arguments.
func Infof(ctx context.Context, format string, args ...any) {
	Log().Infof(ctx, format, args...)
}

// Debugf log by debug level by format and arguments.
func Debugf(ctx context.Context, format string, args ...any) {
	Log().Debugf(ctx, format, args...)
}

// Printf log by info level by format and arguments without context.
func Printf(format string, args ...any) {
	Log().Printf(format, args...)
}

// Fatalf log by alert level by format and arguments without context.
func Fatalf(format string, args ...any) {
	Log().Fatalf(format, args...)
}

// Panicf log by emergency level and arguments without context.
func Panicf(format string, args ...any) {
	Log().Panicf(format, args...)
}

func Writer(ctx context.Context, level level.Level) io.Writer {
	return Log().Writer(ctx, level)
}
