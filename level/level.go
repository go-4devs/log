package level

import (
	"encoding"
	"encoding/json"
	"errors"
	"fmt"
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

// ErrParse is returned when a level name is not one of the known names.
var ErrParse = errors.New("parse level")

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
	if len(in) < 2 || in[0] != '"' || in[len(in)-1] != '"' {
		return fmt.Errorf("%w: %q", ErrParse, in)
	}

	return l.UnmarshalText(in[1 : len(in)-1])
}

func (l Level) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

func (l *Level) UnmarshalText(in []byte) error {
	lvl, err := Parse(string(in))
	if err != nil {
		return err
	}

	*l = lvl

	return nil
}

func (l Level) MarshalBinary() ([]byte, error) {
	return []byte(l.String()), nil
}

func (l *Level) UnmarshalBinary(in []byte) error {
	return l.UnmarshalText(in)
}

func Parse(lvl string) (Level, error) {
	switch strings.ToLower(lvl) {
	case "debug":
		return Debug, nil
	case "info":
		return Info, nil
	case "notice":
		return Notice, nil
	case "warning", "warn":
		return Warning, nil
	case "error", "err":
		return Error, nil
	case "critical", "crit":
		return Critical, nil
	case "alert":
		return Alert, nil
	case "emergency", "emerg":
		return Emergency, nil
	default:
		return 0, fmt.Errorf("%w: %q", ErrParse, lvl)
	}
}
