package log

import (
	"context"
	"fmt"
	"strings"
)

// Handler logged message
type Handler func(ctx context.Context, level Level, msg string, fields Fields)

// Processor processsed additional field
type Processor func(ctx context.Context) Field

// Fields slice field
type Fields []Field

// String implemet stringer
func (f Fields) String() string {
	str := make([]string, len(f))
	for _, field := range f {
		str = append(str, field.String())
	}
	return strings.Join(str, "")
}

// Field struct
type Field struct {
	Key   string
	Value interface{}
}

// String implent stringer
func (f Field) String() string {
	return fmt.Sprintf("%s: %v;", f.Key, f.Value)
}

type option func(*Logger)

// New create new logger by handler
func New(handler Handler, opts ...option) *Logger {
	l := &Logger{
		handler: handler,
	}
	for _, opt := range opts {
		opt(l)
	}

	return l
}

// WithProcessor configure process
func WithProcessor(opts ...Processor) option {
	return func(l *Logger) {
		l.processors = append(l.processors, opts...)
	}
}

// Logger log
type Logger struct {
	handler    Handler
	processors []Processor
}

func (l *Logger) log(ctx context.Context, level Level, args ...interface{}) {
	l.handler(ctx, level, fmt.Sprint(args...), l.fields(ctx))
}

func (l *Logger) logKV(ctx context.Context, level Level, msg string, args ...interface{}) {
	l.handler(ctx, level, msg, l.fields(ctx, args...))
}

func (l *Logger) logf(ctx context.Context, level Level, format string, args ...interface{}) {
	l.log(ctx, level, fmt.Sprintf(format, args...))
}

func (l *Logger) fields(ctx context.Context, args ...interface{}) []Field {
	fields := make([]Field, 0, len(args)+len(l.processors))
	for _, p := range l.processors {
		fields = append(fields, p(ctx))
	}
	for i := 0; i < len(args); {
		if f, ok := args[i].(Field); ok {
			fields = append(fields, f)
			i++
			continue
		}
		if i == len(args)-1 {
			l.handler(ctx, LevelCritical, fmt.Sprint("Ignored key without a value.", args[i]), fields)
			break
		}
		i += 2
		key, val := args[i], args[i+1]
		if keyStr, ok := key.(string); ok {
			fields = append(fields, Field{Key: keyStr, Value: val})
			continue
		}
		l.handler(ctx, LevelCritical, fmt.Sprint("Ignored key-value pairs with non-string keys.", args[i], args[i+1]), fields)
	}
	return fields
}

func (l *Logger) Emerg(ctx context.Context, args ...interface{}) {
	l.log(ctx, LevelEmergency, args...)
}

func (l *Logger) Alert(ctx context.Context, args ...interface{}) {
	l.log(ctx, LevelAlert, args...)
}

func (l *Logger) Crit(ctx context.Context, args ...interface{}) {
	l.log(ctx, LevelCritical, args...)
}

func (l *Logger) Err(ctx context.Context, args ...interface{}) {
	l.log(ctx, LevelError, args...)
}

func (l *Logger) Warn(ctx context.Context, args ...interface{}) {
	l.log(ctx, LevelWarning, args...)
}

func (l *Logger) Notice(ctx context.Context, args ...interface{}) {
	l.log(ctx, LevelNotice, args...)
}

func (l *Logger) Info(ctx context.Context, args ...interface{}) {
	l.log(ctx, LevelInfo, args...)
}

func (l *Logger) Debug(ctx context.Context, args ...interface{}) {
	l.log(ctx, LevelDebug, args...)
}

func (l *Logger) EmergKV(ctx context.Context, msg string, args ...interface{}) {
	l.logKV(ctx, LevelEmergency, msg, args...)
}

func (l *Logger) AlertKV(ctx context.Context, msg string, args ...interface{}) {
	l.logKV(ctx, LevelAlert, msg, args...)
}

func (l *Logger) CritKV(ctx context.Context, msg string, args ...interface{}) {
	l.logKV(ctx, LevelCritical, msg, args...)
}

func (l *Logger) ErrKV(ctx context.Context, msg string, args ...interface{}) {
	l.logKV(ctx, LevelError, msg, args...)
}

func (l *Logger) WarnKV(ctx context.Context, msg string, args ...interface{}) {
	l.logKV(ctx, LevelWarning, msg, args...)
}

func (l *Logger) NoticeKV(ctx context.Context, msg string, args ...interface{}) {
	l.logKV(ctx, LevelNotice, msg, args...)
}

func (l *Logger) InfoKV(ctx context.Context, msg string, args ...interface{}) {
	l.logKV(ctx, LevelInfo, msg, args...)
}

func (l *Logger) DebugKV(ctx context.Context, msg string, args ...interface{}) {
	l.logKV(ctx, LevelDebug, msg, args...)
}

func (l *Logger) Emergf(ctx context.Context, format string, args ...interface{}) {
	l.logf(ctx, LevelEmergency, format, args...)
}

func (l *Logger) Alertf(ctx context.Context, format string, args ...interface{}) {
	l.logf(ctx, LevelAlert, format, args...)
}

func (l *Logger) Critf(ctx context.Context, format string, args ...interface{}) {
	l.logf(ctx, LevelCritical, format, args...)
}

func (l *Logger) Errf(ctx context.Context, format string, args ...interface{}) {
	l.logf(ctx, LevelError, format, args...)
}

func (l *Logger) Warnf(ctx context.Context, format string, args ...interface{}) {
	l.logf(ctx, LevelWarning, format, args...)
}

func (l *Logger) Noticef(ctx context.Context, format string, args ...interface{}) {
	l.logf(ctx, LevelNotice, format, args...)
}

func (l *Logger) Infof(ctx context.Context, format string, args ...interface{}) {
	l.logf(ctx, LevelInfo, format, args...)
}

func (l *Logger) Debugf(ctx context.Context, format string, args ...interface{}) {
	l.logf(ctx, LevelDebug, format, args...)
}
