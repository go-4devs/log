package level_test

import (
	"testing"

	"gitoa.ru/go-4devs/log/level"
)

func TestMarshalJSON(t *testing.T) {
	levels := map[level.Level]string{
		level.Emergency: `"emerg"`,
		level.Alert:     `"alert"`,
		level.Critical:  `"crit"`,
		level.Error:     `"error"`,
		level.Warning:   `"warning"`,
		level.Notice:    `"notice"`,
		level.Info:      `"info"`,
		level.Debug:     `"debug"`,
	}

	for level, expect := range levels {
		actual, err := level.MarshalJSON()
		if err != nil {
			t.Errorf("%s got err: %s", level, err)
			continue
		}
		if string(actual) != expect {
			t.Errorf("%s got: %s expect: %s", level, actual, expect)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	levels := map[level.Level][]string{
		level.Emergency: {`"emerg"`, `"Emerg"`},
		level.Alert:     {`"alert"`, `"ALERT"`},
		level.Critical:  {`"crit"`, `"critical"`},
		level.Error:     {`"error"`, `"ERR"`},
		level.Warning:   {`"warning"`, `"Warning"`},
		level.Notice:    {`"notice"`},
		level.Info:      {`"info"`},
		level.Debug:     {`"debug"`, `"DEBUG"`},
	}

	for expect, actuals := range levels {
		for _, actual := range actuals {
			var level level.Level
			if err := level.UnmarshalJSON([]byte(actual)); err != nil {
				t.Errorf("%s got err: %s", level, err)
				continue
			}

			if !level.Is(expect) {
				t.Errorf("%s got: %s expect: %s", actual, level, expect)
			}

		}
	}
}
