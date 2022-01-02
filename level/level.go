package level

import (
	"encoding/json"
	"fmt"
	"strings"
)

//go:generate stringer -type=Level -linecomment

var (
	_ json.Marshaler   = Level(0)
	_ json.Unmarshaler = (*Level)(nil)
)

// Level log.
type Level uint32

// available log levels.
const (
	Emergency Level = iota // emergency
	Alert                  // alert
	Critical               // critical
	Error                  // error
	Warning                // warning
	Notice                 // notice
	Info                   // info
	Debug                  // debug
)

func (l Level) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(l.String())
	if err != nil {
		return nil, fmt.Errorf("marshal err: %w", err)
	}

	return b, nil
}

func (l Level) Is(level Level) bool {
	return level == l
}

func (l Level) Enabled(level Level) bool {
	return l <= level
}

func (l *Level) UnmarshalJSON(in []byte) error {
	var v string
	if err := json.Unmarshal(in, &v); err != nil {
		return fmt.Errorf("unmarshal err: %w", err)
	}

	lvl := Parse(v)
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
	case "warning", "Warning", "WARNING":
		return Warning
	case "error", "Error", "ERROR":
		return Error
	case "critical", "Critical", "CRITICAL":
		return Critical
	case "alert", "Alert", "ALERT":
		return Alert
	default:
		return Emergency
	}
}
