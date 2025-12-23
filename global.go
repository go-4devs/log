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
	global.Emerg(ctx, args...)
}

// Alert log by alert level.
func Alert(ctx context.Context, args ...any) {
	global.Alert(ctx, args...)
}

// Crit log by critical level.
func Crit(ctx context.Context, args ...any) {
	global.Crit(ctx, args...)
}

// Err log by error level.
func Err(ctx context.Context, args ...any) {
	global.Err(ctx, args...)
}

// Warn logs by warning level.
func Warn(ctx context.Context, args ...any) {
	global.Warn(ctx, args...)
}

// Notice log by notice level.
func Notice(ctx context.Context, args ...any) {
	global.Notice(ctx, args...)
}

// Info log by info level.
func Info(ctx context.Context, args ...any) {
	global.Info(ctx, args...)
}

// Debug log by debug level.
func Debug(ctx context.Context, args ...any) {
	global.Debug(ctx, args...)
}

// Print log by info level and arguments.
func Print(args ...any) {
	global.Print(args...)
}

// Fatal log by alert level and arguments.
func Fatal(args ...any) {
	global.Fatal(args...)
}

// Panic log by emergency level and arguments.
func Panic(args ...any) {
	global.Panic(args...)
}

// Println log by info level and arguments.
func Println(args ...any) {
	global.Println(args...)
}

// Fatalln log by alert level and arguments.
func Fatalln(args ...any) {
	global.Fatalln(args...)
}

// Panicln log by emergency level and arguments.
func Panicln(args ...any) {
	global.Panicln(args...)
}

// EmergKVs sugared log by emergency level and key-values.
func EmergKVs(ctx context.Context, msg string, args ...any) {
	global.EmergKVs(ctx, msg, args...)
}

// AlertKVs sugared log by alert level and key-values.
func AlertKVs(ctx context.Context, msg string, args ...any) {
	global.AlertKVs(ctx, msg, args...)
}

// CritKVs sugared log by critcal level and key-values.
func CritKVs(ctx context.Context, msg string, args ...any) {
	global.CritKVs(ctx, msg, args...)
}

// ErrKVs sugared log by error level and key-values.
func ErrKVs(ctx context.Context, msg string, args ...any) {
	global.ErrKVs(ctx, msg, args...)
}

// WarnKVs sugared log by warning level and key-values.
func WarnKVs(ctx context.Context, msg string, args ...any) {
	global.WarnKVs(ctx, msg, args...)
}

// NoticeKVs sugared log by notice level and key-values.
func NoticeKVs(ctx context.Context, msg string, args ...any) {
	global.NoticeKVs(ctx, msg, args...)
}

// InfoKVs sugared log by info level and key-values.
func InfoKVs(ctx context.Context, msg string, args ...any) {
	global.InfoKVs(ctx, msg, args...)
}

// DebugKVs sugared log by debug level and key-values.
func DebugKVs(ctx context.Context, msg string, args ...any) {
	global.DebugKVs(ctx, msg, args...)
}

// EmergKV  log by emergency level and key-values.
func EmergKV(ctx context.Context, msg string, args ...field.Field) {
	global.EmergKV(ctx, msg, args...)
}

// AlertKV log by alert level and key-values.
func AlertKV(ctx context.Context, msg string, args ...field.Field) {
	global.AlertKV(ctx, msg, args...)
}

// CritKV log by critcal level and key-values.
func CritKV(ctx context.Context, msg string, args ...field.Field) {
	global.CritKV(ctx, msg, args...)
}

// ErrKV log by error level and key-values.
func ErrKV(ctx context.Context, msg string, args ...field.Field) {
	global.ErrKV(ctx, msg, args...)
}

// WarnKV log by warning level and key-values.
func WarnKV(ctx context.Context, msg string, args ...field.Field) {
	global.WarnKV(ctx, msg, args...)
}

// NoticeKV log by notice level and key-values.
func NoticeKV(ctx context.Context, msg string, args ...field.Field) {
	global.NoticeKV(ctx, msg, args...)
}

// InfoKV log by info level and key-values.
func InfoKV(ctx context.Context, msg string, args ...field.Field) {
	global.InfoKV(ctx, msg, args...)
}

// DebugKV log by debug level and key-values.
func DebugKV(ctx context.Context, msg string, args ...field.Field) {
	global.DebugKV(ctx, msg, args...)
}

// Emergf log by emergency level by format and arguments.
func Emergf(ctx context.Context, format string, args ...any) {
	global.Emergf(ctx, format, args...)
}

// Alertf log by alert level by format and arguments.
func Alertf(ctx context.Context, format string, args ...any) {
	global.Alertf(ctx, format, args...)
}

// Critf log by critical level by format and arguments.
func Critf(ctx context.Context, format string, args ...any) {
	global.Critf(ctx, format, args...)
}

// Errf log by error level by format and arguments.
func Errf(ctx context.Context, format string, args ...any) {
	global.Errf(ctx, format, args...)
}

// Warnf log by warning level by format and arguments.
func Warnf(ctx context.Context, format string, args ...any) {
	global.Warnf(ctx, format, args...)
}

// Noticef log by notice level by format and arguments.
func Noticef(ctx context.Context, format string, args ...any) {
	global.Noticef(ctx, format, args...)
}

// Infof log by info level by format and arguments.
func Infof(ctx context.Context, format string, args ...any) {
	global.Noticef(ctx, format, args...)
}

// Debugf log by debug level by format and arguments.
func Debugf(ctx context.Context, format string, args ...any) {
	global.Debugf(ctx, format, args...)
}

// Printf log by info level by format and arguments without context.
func Printf(format string, args ...any) {
	global.Printf(format, args...)
}

// Fatalf log by alert level by format and arguments without context.
func Fatalf(format string, args ...any) {
	global.Fatalf(format, args...)
}

// Panicf log by emergency level and arguments without context.
func Panicf(format string, args ...any) {
	global.Panicf(format, args...)
}

func Writer(ctx context.Context, level level.Level) io.Writer {
	return global.Writer(ctx, level)
}
