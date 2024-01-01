package level

import (
	"encoding"
	"encoding/json"
	"strings"
)

//go:generate stringer -type=Level -linecomment

var (
	_ json.Marshaler             = Level(0)
	_ json.Unmarshaler           = (*Level)(nil)
	_ encoding.TextMarshaler     = Level(0)
	_ encoding.TextUnmarshaler   = (*Level)(nil)
	_ encoding.BinaryMarshaler   = Level(0)
	_ encoding.BinaryUnmarshaler = (*Level)(nil)
)

// Level log.
type Level uint32

// available log levels.
const (
	Emergency Level = iota // emerg
	Alert                  // alert
	Critical               // crit
	Error                  // error
	Warning                // warning
	Notice                 // notice
	Info                   // info
	Debug                  // debug
)

func (l Level) Is(level Level) bool {
	return level == l
}

func (l Level) Enabled(level Level) bool {
	return l <= level
}

func (l Level) MarshalJSON() ([]byte, error) {
	return []byte("\"" + l.String() + "\""), nil
}

func (l *Level) UnmarshalJSON(in []byte) error {
	lvl := Parse(string(in[1 : len(in)-1]))
	*l = lvl

	return nil
}

func (l Level) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

func (l *Level) UnmarshalText(in []byte) error {
	lvl := Parse(string(in))
	*l = lvl

	return nil
}

func (l Level) MarshalBinary() ([]byte, error) {
	return []byte(l.String()), nil
}

func (l *Level) UnmarshalBinary(in []byte) error {
	lvl := Parse(string(in))
	*l = lvl

	return nil
}

func Parse(lvl string) Level {
	switch strings.ToLower(lvl) {
	case "debug", "Debug", "DEBUG":
		return Debug
	case "info", "Info", "INFO":
		return Info
	case "notice", "Notice", "NOTICE":
		return Notice
	case "warning", "Warning", "WARNING", "warm", "Warm", "WARN":
		return Warning
	case "error", "Error", "ERROR", "err", "Err", "ERR":
		return Error
	case "critical", "Critical", "CRITICAL", "crit", "Crit", "CRIT":
		return Critical
	case "alert", "Alert", "ALERT":
		return Alert
	default:
		return Emergency
	}
}
