package otel

//go:generate stringer -type=Level -linecomment -output=level_string.go

type Level int

const (
	levelDebug  Level = 5  // DEBUG
	levelInfo   Level = 9  // INFO
	levelInfo2  Level = 10 // INFO2
	levelWarn   Level = 13 // WARN
	levelError  Level = 17 // ERROR
	levelError2 Level = 18 // ERROR2
	levelError3 Level = 19 // ERROR3
	levelFatal  Level = 21 // FATAL
)
