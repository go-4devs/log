package log_test

import (
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"sync/atomic"
	"time"

	"gitoa.ru/go-4devs/log"
	"gitoa.ru/go-4devs/log/entry"
	"gitoa.ru/go-4devs/log/field"
	"gitoa.ru/go-4devs/log/level"
)

var ctx = context.Background()

func setStdout() {
	// set stout for example by default stderror
	log.SetLogger(log.New(log.WithStdout()).With(log.WithLevel(log.KeyLevel, level.Debug)))
}

func ExampleNew() {
	logger := log.New(log.WithStdout())
	logger.Info(ctx, "same message")
	// Output: msg="same message"
}

func ExampleInfo() {
	setStdout()
	log.Info(ctx, "same message")
	// Output: msg="same message" level=info
}

func ExampleErrKV() {
	setStdout()
	log.ErrKVs(ctx, "same message", "key", "addition value")
	// Output: msg="same message" key="addition value" level=error
}

func ExampleNew_errf() {
	logger := log.New(log.WithStdout())
	logger.Errf(ctx, "same message %d", 1)
	// Output: msg="same message 1"
}

func ExampleNew_debugKV() {
	logger := log.New(log.WithStdout()).With(log.WithLevel(log.KeyLevel, level.Debug))
	logger.DebugKVs(ctx, "same message", "error", os.ErrNotExist)
	// Output: msg="same message" error="file does not exist" level=debug
}

func ExampleNew_level() {
	logger := log.New(log.WithStdout()).With(log.WithLevel(log.KeyLevel, level.Error))
	logger.Err(ctx, "same error message")
	// Output: msg="same error message" level=error
}

func ExampleNew_level_info() {
	logger := log.New(log.WithStdout()).With(log.WithLevel(log.KeyLevel, level.Error))
	logger.Info(ctx, "same message")
	// Output:
}

type Obj struct {
	Name     string
	IsEnable bool
}

var (
	obj = Obj{
		Name: "test obj",
	}

	str      = "test str"
	boolsVal = true
	intVal   = int(math.MaxInt)
	int8Val  = int8(math.MaxInt8)
	int16Val = int16(math.MaxInt16)
	int32Val = int32(math.MaxInt32)
	int64Val = int64(math.MaxInt64)

	uintVal   = uint(math.MaxUint)
	uint8Val  = uint8(math.MaxUint8)
	uint16Val = uint16(math.MaxInt16)
	uint32Val = uint32(math.MaxInt32)
	uint64Val = uint64(math.MaxInt64)

	float32Val = float32(math.MaxFloat32)
	float64Val = float64(math.MaxFloat64)

	minute  = time.Minute
	timeVal = time.Unix(0, math.MaxInt32).In(time.UTC)
)

func ExampleNew_anyField() {
	logger := log.New(log.WithStdout(), log.WithJSONFormat())
	logger.InfoKV(ctx, "any info message",
		field.Any("obj", Obj{Name: "obj name"}),
		field.Any("obj", &obj),
		field.Any("int", intVal),
		field.Any("uint", uintVal),
		field.Any("float", float64Val),
		field.Any("time", timeVal),
		field.Any("duration", time.Hour),
		field.Any("error", errors.New("error")),
	)
	// Output:
	// {"msg":"any info message","obj":{"Name":"obj name","IsEnable":false},"obj":{"Name":"test obj","IsEnable":false},"int":9223372036854775807,"uint":18446744073709551615,"float":1.7976931348623157e+308,"time":"1970-01-01T00:00:02Z","duration":"1h0m0s","error":"error"}
}

func ExampleNew_arrayField() {
	logger := log.New(log.WithStdout(), log.WithJSONFormat())
	logger.InfoKV(ctx, "array info message",
		field.Strings("strings", "string", str),
		field.Bools("bools", true, false),
		field.Ints("ints", 42, 24),
		field.Int8s("int8s", 42, 24),
		field.Int16s("int16s", 42, 24),
		field.Int32s("int32s", 42, 24),
		field.Int64s("int64s", 42, 24),
		field.Uint8s("uint8s", uint8Val, 0),
		field.Uint16s("uint16s", 42, 24),
		field.Uint32s("uint32s", 42, 24),
		field.Uint64s("uint64s", 42, 24),
		field.Float32s("float32s", 42, 24),
		field.Float64s("float64s", 42, 24),
		field.Complex64s("complex64s", 42, 24),
		field.Complex128s("complex128s", 42, 24),
		field.Durations("durations", time.Minute, time.Second),
		field.Times("times", time.Unix(0, 42).In(time.UTC), time.Unix(0, 24).In(time.UTC)),
		field.Errors("errors", errors.New("error"), errors.New("error2")),
	)
	// Output:
	// {"msg":"array info message","strings":["string","test str"],"bools":[true,false],"ints":[42,24],"int8s":[42,24],"int16s":[42,24],"int32s":[42,24],"int64s":[42,24],"uint8s":[255,0],"uint16s":[42,24],"uint32s":[42,24],"uint64s":[42,24],"float32s":[42,24],"float64s":[42,24],"complex64s":["(42+0i)","(24+0i)"],"complex128s":["(42+0i)","(24+0i)"],"durations":["1m0s","1s"],"times":["1970-01-01T00:00:00Z","1970-01-01T00:00:00Z"],"errors":["error","error2"]}
}

func ExampleNew_pointerField() {
	logger := log.New(log.WithStdout(), log.WithJSONFormat())
	logger.InfoKV(ctx, "pointer info message",
		field.Stringp("stringp", &str),
		field.Stringp("stringp", nil),
		field.Boolp("boolp", &boolsVal),
		field.Boolp("boolp", nil),
		field.Intp("intp", &intVal),
		field.Intp("intp", nil),
		field.Int8p("int8p", &int8Val),
		field.Int8p("int8p", nil),
		field.Int16p("int16p", &int16Val),
		field.Int16p("int16p", nil),
		field.Int32p("int32p", &int32Val),
		field.Int32p("int32p", nil),
		field.Int64p("int64p", &int64Val),
		field.Int64p("int64p", nil),
		field.Uintp("uintp", &uintVal),
		field.Uintp("uintp", nil),
		field.Uint8p("uint8p", &uint8Val),
		field.Uint8p("uint8p", nil),
		field.Uint16p("uint16p", &uint16Val),
		field.Uint16p("uint16p", nil),
		field.Uint32p("uint32p", &uint32Val),
		field.Uint32p("uint32p", nil),
		field.Uint64p("uint64p", &uint64Val),
		field.Uint64p("uint64p", nil),
		field.Float32p("float32p", &float32Val),
		field.Float32p("float32p", nil),
		field.Float64p("float64p", &float64Val),
		field.Float64p("float64p", nil),
		field.Durationp("durationp", &minute),
		field.Durationp("durationp", nil),
		field.Timep("timep", &timeVal),
		field.Timep("timep", nil),
	)
	// Output:
	// {"msg":"pointer info message","stringp":"test str","stringp":null,"boolp":true,"boolp":null,"intp":9223372036854775807,"intp":null,"int8p":127,"int8p":null,"int16p":32767,"int16p":null,"int32p":2147483647,"int32p":null,"int64p":9223372036854775807,"int64p":null,"uintp":18446744073709551615,"uintp":null,"uint8p":255,"uint8p":null,"uint16p":32767,"uint16p":null,"uint32p":2147483647,"uint32p":null,"uint64p":9223372036854775807,"uint64p":null,"float32p":3.4028235e+38,"float32p":null,"float64p":1.7976931348623157e+308,"float64p":null,"durationp":"1m0s","durationp":null,"timep":"1970-01-01T00:00:02Z","timep":null}
}

func ExampleNew_fields() {
	logger := log.New(log.WithStdout(), log.WithJSONFormat())
	logger.InfoKV(ctx, "info message",
		field.String("string", str),
		field.Bool("bool", true),
		field.Int("int", 42),
		field.Int8("int8", 42),
		field.Int16("int16", 42),
		field.Int32("int32", 42),
		field.Int64("int64", 42),
		field.Uint8("uint8", uint8Val),
		field.Uint16("uint16", 42),
		field.Uint32("uint32", 42),
		field.Uint64("uint64", 42),
		field.Float32("float32", 42),
		field.Float64("float64", 42),
		field.Complex64("complex16", 42),
		field.Complex128("complex128", 42),
		field.Duration("duration", time.Minute),
		field.Time("time", timeVal),
		field.FormatTime("format_time", time.UnixDate, timeVal),
		field.Error("error", errors.New("error")),
	)
	// Output:
	// {"msg":"info message","string":"test str","bool":true,"int":42,"int8":42,"int16":42,"int32":42,"int64":42,"uint8":255,"uint16":42,"uint32":42,"uint64":42,"float32":42,"float64":42,"complex16":"(42+0i)","complex128":"(42+0i)","duration":"1m0s","time":"1970-01-01T00:00:02Z","format_time":"Thu Jan  1 00:00:02 UTC 1970","error":"error"}
}

func ExampleNew_jsonFormat() {
	logger := log.New(log.WithStdout(), log.WithJSONFormat()).
		With(
			log.WithLevel(log.KeyLevel, level.Debug),
			log.GoVersion("go-version"),
		)
	logger.Err(ctx, "same error message")
	logger.WarnKVs(ctx, "same warn message", "obj", Obj{Name: "obj name"})
	// Output:
	// {"msg":"same error message","level":"error","go-version":"go1.25.5"}
	// {"msg":"same warn message","obj":{"Name":"obj name","IsEnable":false},"level":"warning","go-version":"go1.25.5"}
}

func ExampleNew_textEncoding() {
	logger := log.New(log.WithStdout()).
		With(
			log.WithLevel(log.KeyLevel, level.Debug),
			log.GoVersion("go-version"),
		)
	logger.Err(ctx, "same error message")
	logger.InfoKVs(ctx, "same info message", "api-version", 0.1, "obj", Obj{Name: "text value", IsEnable: true})

	// Output:
	// msg="same error message" level=error go-version=go1.25.5
	// msg="same info message" api-version=0.1 obj={Name:text value IsEnable:true} level=info go-version=go1.25.5
}

type ctxKey string

func (c ctxKey) String() string {
	return string(c)
}

func levelInfo(ctx context.Context, entry *entry.Entry, handler log.Logger) (int, error) {
	return handler(ctx, entry.Add(field.String(log.KeyLevel, entry.Level().String())))
}

func ExampleWith() {
	var requestID ctxKey = "requestID"

	vctx := context.WithValue(ctx, requestID, "6a5fa048-7181-11ea-bc55-0242ac130003")

	logger := log.New(log.WithStdout()).With(
		levelInfo,
		log.WithContextValue(requestID),
		log.KeyValue("api", "0.1.0"),
		log.GoVersion("go"),
	)
	logger.Info(vctx, "same message")
	// Output: msg="same message" level=info requestID=6a5fa048-7181-11ea-bc55-0242ac130003 api=0.1.0 go=go1.25.5
}

func ExampleLogger_Print() {
	logger := log.New(log.WithStdout()).With(
		levelInfo,
		log.KeyValue("client", "http"),
		log.KeyValue("api", "0.1.0"),
		log.GoVersion("go"),
	)
	logger.Print("same message")
	// Output: msg="same message" level=info client=http api=0.1.0 go=go1.25.5
}

func ExamplePrint() {
	setStdout()
	log.Print("same message")
	// Output: msg="same message" level=info
}

func Example_fieldClosureFn() {
	cnt := int32(0)
	closure := field.ClosureFn(func() any {
		d := fmt.Sprintf("additional error data: %d", cnt)
		atomic.AddInt32(&cnt, 1)

		return d
	})

	log := log.New(log.WithStdout()).With(log.WithLevel(log.KeyLevel, level.Info))

	log.DebugKVs(ctx, "debug message", "data", closure)
	log.ErrKVs(ctx, "error message", "err", closure)
	log.WarnKVs(ctx, "warn message", "warn", closure)

	// Output:
	// msg="error message" err="additional error data: 0" level=error
	// msg="warn message" warn="additional error data: 1" level=warning
}

func Example_withGroup() {
	log := log.New(log.WithStdout()).With(
		log.WithLevel(log.KeyLevel, level.Info),
	)

	log.ErrKVs(ctx, "error message",
		field.Groups("grous_field",
			field.Error("err", os.ErrDeadlineExceeded),
			field.Bool("bool", false),
		),
	)
	log.WarnKV(ctx, "error message", field.ValuerFn("valuer_field", func() any {
		return field.Fields{
			field.Int("int_value", math.MaxInt),
			field.Uint8("uint8_value", math.MaxUint8),
		}
	}))

	// Output:
	// msg="error message" grous_field.err="i/o timeout" grous_field.bool=false level=error
	// msg="error message" valuer_field.int_value=9223372036854775807 valuer_field.uint8_value=255 level=warning
}
