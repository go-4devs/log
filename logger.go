package log

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"gitoa.ru/go-4devs/log/entry"
	"gitoa.ru/go-4devs/log/field"
	"gitoa.ru/go-4devs/log/level"
)

var _ io.Writer = (Logger)(nil)

var (
	ErrIgnoredKey    = errors.New("ignored key without a value")
	ErrNonStringKeys = errors.New("ignored key-value pairs with non-string keys")
)

func writeOutput(_ int, err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}

// Logger logged message.
type Logger func(ctx context.Context, entry *entry.Entry) (int, error)

func (l Logger) Write(in []byte) (int, error) {
	return l.write(context.Background(), level.Info, string(in))
}

func (l Logger) write(ctx context.Context, level level.Level, msg string, fields ...field.Field) (int, error) {
	data := entry.Get()

	defer func() {
		entry.Put(data)
	}()

	return l(ctx, data.SetLevel(level).SetMessage(msg).Add(fields...))
}

func (l Logger) writef(ctx context.Context, level level.Level, format string, args ...interface{}) (int, error) {
	data := entry.Get()

	defer func() {
		entry.Put(data)
	}()

	return l(ctx, data.SetLevel(level).SetMessagef(format, args...))
}

func (l Logger) kv(ctx context.Context, args ...interface{}) field.Fields {
	kvEntry := entry.Get()

	defer func() {
		entry.Put(kvEntry)
	}()

	for i := 0; i < len(args); i++ {
		if f, ok := args[i].(field.Field); ok {
			kvEntry = kvEntry.Add(f)

			continue
		}

		if i == len(args)-1 {
			writeOutput(l.write(ctx, level.Critical, fmt.Sprint("Ignored key without a value.", args[i]), kvEntry.Fields()...))

			break
		}

		i++

		key, val := args[i-1], args[i]
		if keyStr, ok := key.(string); ok {
			kvEntry = kvEntry.AddAny(keyStr, val)

			continue
		}

		writeOutput(l.write(ctx, level.Critical, fmt.Sprint("Ignored key-value pairs with non-string keys.", key, val), kvEntry.Fields()...))
	}

	return kvEntry.Fields()
}

// With adds middlewares to logger.
func (l Logger) With(mw ...Middleware) Logger {
	return With(l, mw...)
}

// Emerg log by emergency level.
func (l Logger) Emerg(ctx context.Context, args ...interface{}) {
	writeOutput(l.writef(ctx, level.Emergency, "", args...))
}

// Alert log by alert level.
func (l Logger) Alert(ctx context.Context, args ...interface{}) {
	writeOutput(l.writef(ctx, level.Alert, "", args...))
}

// Crit log by critical level.
func (l Logger) Crit(ctx context.Context, args ...interface{}) {
	writeOutput(l.writef(ctx, level.Critical, "", args...))
}

// Err log by error level.
func (l Logger) Err(ctx context.Context, args ...interface{}) {
	writeOutput(l.writef(ctx, level.Error, "", args...))
}

// Warn log by warning level.
func (l Logger) Warn(ctx context.Context, args ...interface{}) {
	writeOutput(l.writef(ctx, level.Warning, "", args...))
}

// Notice log by notice level.
func (l Logger) Notice(ctx context.Context, args ...interface{}) {
	writeOutput(l.writef(ctx, level.Notice, "", args...))
}

// Info log by info level.
func (l Logger) Info(ctx context.Context, args ...interface{}) {
	writeOutput(l.writef(ctx, level.Info, "", args...))
}

// Debug log by debug level.
func (l Logger) Debug(ctx context.Context, args ...interface{}) {
	writeOutput(l.writef(ctx, level.Debug, "", args...))
}

// Print log by info level and arguments.
func (l Logger) Print(args ...interface{}) {
	writeOutput(l.writef(context.Background(), level.Info, "", args...))
}

// Fatal log by alert level and arguments.
func (l Logger) Fatal(args ...interface{}) {
	writeOutput(l.writef(context.Background(), level.Alert, "", args...))
}

// Panic log by emergency level and arguments.
func (l Logger) Panic(args ...interface{}) {
	writeOutput(l.writef(context.Background(), level.Emergency, "", args...))
}

// Println log by info level and arguments.
func (l Logger) Println(args ...interface{}) {
	writeOutput(l.write(context.Background(), level.Info, fmt.Sprintln(args...)))
}

// Fatalln log by alert level and arguments.
func (l Logger) Fatalln(args ...interface{}) {
	writeOutput(l.write(context.Background(), level.Alert, fmt.Sprintln(args...)))
}

// Panicln log by emergency level and arguments.
func (l Logger) Panicln(args ...interface{}) {
	writeOutput(l.write(context.Background(), level.Emergency, fmt.Sprintln(args...)))
}

// EmergKVs sugared log by emergency level and key-values.
func (l Logger) EmergKVs(ctx context.Context, msg string, args ...interface{}) {
	writeOutput(l.write(ctx, level.Emergency, msg, l.kv(ctx, args...)...))
}

// AlertKVs sugared log by alert level and key-values.
func (l Logger) AlertKVs(ctx context.Context, msg string, args ...interface{}) {
	writeOutput(l.write(ctx, level.Alert, msg, l.kv(ctx, args...)...))
}

// CritKVs sugared log by critcal level and key-values.
func (l Logger) CritKVs(ctx context.Context, msg string, args ...interface{}) {
	writeOutput(l.write(ctx, level.Critical, msg, l.kv(ctx, args...)...))
}

// ErrKVs sugared log by error level and key-values.
func (l Logger) ErrKVs(ctx context.Context, msg string, args ...interface{}) {
	writeOutput(l.write(ctx, level.Error, msg, l.kv(ctx, args...)...))
}

// WarnKVs sugared log by warning level and key-values.
func (l Logger) WarnKVs(ctx context.Context, msg string, args ...interface{}) {
	writeOutput(l.write(ctx, level.Warning, msg, l.kv(ctx, args...)...))
}

// NoticeKVs sugared log by notice level and key-values.
func (l Logger) NoticeKVs(ctx context.Context, msg string, args ...interface{}) {
	writeOutput(l.write(ctx, level.Notice, msg, l.kv(ctx, args...)...))
}

// InfoKVs sugared log by info level and key-values.
func (l Logger) InfoKVs(ctx context.Context, msg string, args ...interface{}) {
	writeOutput(l.write(ctx, level.Info, msg, l.kv(ctx, args...)...))
}

// DebugKVs sugared log by debug level and key-values.
func (l Logger) DebugKVs(ctx context.Context, msg string, args ...interface{}) {
	writeOutput(l.write(ctx, level.Debug, msg, l.kv(ctx, args...)...))
}

// EmergKV log by emergency level and key-values.
func (l Logger) EmergKV(ctx context.Context, msg string, args ...field.Field) {
	writeOutput(l.write(ctx, level.Emergency, msg, args...))
}

// AlertKV log by alert level and key-values.
func (l Logger) AlertKV(ctx context.Context, msg string, args ...field.Field) {
	writeOutput(l.write(ctx, level.Alert, msg, args...))
}

// CritKV log by critcal level and key-values.
func (l Logger) CritKV(ctx context.Context, msg string, args ...field.Field) {
	writeOutput(l.write(ctx, level.Critical, msg, args...))
}

// ErrKV log by error level and key-values.
func (l Logger) ErrKV(ctx context.Context, msg string, args ...field.Field) {
	writeOutput(l.write(ctx, level.Error, msg, args...))
}

// WarnKV log by warning level and key-values.
func (l Logger) WarnKV(ctx context.Context, msg string, args ...field.Field) {
	writeOutput(l.write(ctx, level.Warning, msg, args...))
}

// NoticeKV log by notice level and key-values.
func (l Logger) NoticeKV(ctx context.Context, msg string, args ...field.Field) {
	writeOutput(l.write(ctx, level.Notice, msg, args...))
}

// InfoKV log by info level and key-values.
func (l Logger) InfoKV(ctx context.Context, msg string, args ...field.Field) {
	writeOutput(l.write(ctx, level.Info, msg, args...))
}

// DebugKV log by debug level and key-values.
func (l Logger) DebugKV(ctx context.Context, msg string, args ...field.Field) {
	writeOutput(l.write(ctx, level.Debug, msg, args...))
}

// Emergf log by emergency level by format and arguments.
func (l Logger) Emergf(ctx context.Context, format string, args ...interface{}) {
	writeOutput(l.writef(ctx, level.Emergency, format, args...))
}

// Alertf log by alert level by format and arguments.
func (l Logger) Alertf(ctx context.Context, format string, args ...interface{}) {
	writeOutput(l.writef(ctx, level.Alert, format, args...))
}

// Critf log by critical level by format and arguments.
func (l Logger) Critf(ctx context.Context, format string, args ...interface{}) {
	writeOutput(l.writef(ctx, level.Critical, format, args...))
}

// Errf log by error level by format and arguments.
func (l Logger) Errf(ctx context.Context, format string, args ...interface{}) {
	writeOutput(l.writef(ctx, level.Error, format, args...))
}

// Warnf log by warning level by format and arguments.
func (l Logger) Warnf(ctx context.Context, format string, args ...interface{}) {
	writeOutput(l.writef(ctx, level.Warning, format, args...))
}

// Noticef log by notice level by format and arguments.
func (l Logger) Noticef(ctx context.Context, format string, args ...interface{}) {
	writeOutput(l.writef(ctx, level.Notice, format, args...))
}

// Infof log by info level by format and arguments.
func (l Logger) Infof(ctx context.Context, format string, args ...interface{}) {
	writeOutput(l.writef(ctx, level.Info, format, args...))
}

// Debugf log by debug level by format and arguments.
func (l Logger) Debugf(ctx context.Context, format string, args ...interface{}) {
	writeOutput(l.writef(ctx, level.Debug, format, args...))
}

// Printf log by info level by format and arguments without context.
func (l Logger) Printf(format string, args ...interface{}) {
	writeOutput(l.writef(context.Background(), level.Info, format, args...))
}

// Fatalf log by alert level by format and arguments without context.
func (l Logger) Fatalf(format string, args ...interface{}) {
	writeOutput(l.writef(context.Background(), level.Alert, format, args...))
}

// Panicf log by emergency level and arguments without context.
func (l Logger) Panicf(format string, args ...interface{}) {
	writeOutput(l.writef(context.Background(), level.Emergency, format, args...))
}

func (l Logger) Writer(ctx context.Context, level level.Level, fields ...field.Field) io.Writer {
	return writer{
		ctx:    ctx,
		level:  level,
		Logger: l,
		fields: fields,
	}
}

//nolint: containedctx
type writer struct {
	ctx    context.Context
	level  level.Level
	fields []field.Field
	Logger
}

func (w writer) WithLevel(level level.Level) writer {
	return writer{
		level:  level,
		Logger: w.Logger,
		ctx:    w.ctx,
		fields: w.fields,
	}
}

func (w writer) WithContext(ctx context.Context) writer {
	return writer{
		level:  w.level,
		Logger: w.Logger,
		ctx:    ctx,
		fields: w.fields,
	}
}

func (w writer) WithFields(fields ...field.Field) writer {
	return writer{
		level:  w.level,
		Logger: w.Logger,
		ctx:    w.ctx,
		fields: fields,
	}
}

func (w writer) Write(in []byte) (int, error) {
	return w.write(w.ctx, w.level, string(in), w.fields...)
}
