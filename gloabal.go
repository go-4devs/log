package log

import (
	"context"
	"log"
	"os"
)

var (
	global = New(NewStdHandler(log.New(os.Stderr, "", log.LstdFlags), LevelDebug))
)

func SetLogger(l *Logger) {
	global = l
}

func GetLogger() *Logger {
	return global
}

func Emerg(ctx context.Context, args ...interface{}) {
	global.Emerg(ctx, args...)
}

func Alert(ctx context.Context, args ...interface{}) {
	global.Alert(ctx, args...)
}

func Crit(ctx context.Context, args ...interface{}) {
	global.Crit(ctx, args...)
}

func Err(ctx context.Context, args ...interface{}) {
	global.Err(ctx, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	global.Warn(ctx, args...)
}

func Notice(ctx context.Context, args ...interface{}) {
	global.Notice(ctx, args...)
}

func Info(ctx context.Context, args ...interface{}) {
	global.Info(ctx, args...)
}

func Debug(ctx context.Context, args ...interface{}) {
	global.Debug(ctx, args...)
}

func EmergKV(ctx context.Context, msg string, args ...interface{}) {
	global.EmergKV(ctx, msg, args...)
}

func AlertKV(ctx context.Context, msg string, args ...interface{}) {
	global.AlertKV(ctx, msg, args...)
}

func CritKV(ctx context.Context, msg string, args ...interface{}) {
	global.CritKV(ctx, msg, args...)
}

func ErrKV(ctx context.Context, msg string, args ...interface{}) {
	global.ErrKV(ctx, msg, args...)
}

func WarnKV(ctx context.Context, msg string, args ...interface{}) {
	global.WarnKV(ctx, msg, args...)
}

func NoticeKV(ctx context.Context, msg string, args ...interface{}) {
	global.NoticeKV(ctx, msg, args...)
}

func InfoKV(ctx context.Context, msg string, args ...interface{}) {
	global.InfoKV(ctx, msg, args...)
}

func DebugKV(ctx context.Context, msg string, args ...interface{}) {
	global.DebugKV(ctx, msg, args...)
}

func Emergf(ctx context.Context, format string, args ...interface{}) {
	global.Emergf(ctx, format, args...)
}

func Alertf(ctx context.Context, format string, args ...interface{}) {
	global.Alertf(ctx, format, args...)
}

func Critf(ctx context.Context, format string, args ...interface{}) {
	global.Critf(ctx, format, args...)
}

func Errf(ctx context.Context, format string, args ...interface{}) {
	global.Errf(ctx, format, args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	global.Warnf(ctx, format, args...)
}

func Noticef(ctx context.Context, format string, args ...interface{}) {
	global.Noticef(ctx, format, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	global.Noticef(ctx, format, args...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	global.Debugf(ctx, format, args...)
}
