package level_test

import (
	"encoding/json"
	"testing"

	"github.com/go-4devs/log/level"
)

func TestMarshalJSON(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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

			err := level.UnmarshalJSON([]byte(actual))
			if err != nil {
				t.Errorf("%s got err: %s", level, err)

				continue
			}

			if !level.Is(expect) {
				t.Errorf("%s got: %s expect: %s", actual, level, expect)
			}
		}
	}
}

func TestParse(t *testing.T) {
	t.Parallel()

	for _, name := range []string{"emerg", "Emerg", "emergency"} {
		got, err := level.Parse(name)
		if err != nil || got != level.Emergency {
			t.Errorf("%s: got %s err %v", name, got, err)
		}
	}

	if _, err := level.Parse("nope"); err == nil {
		t.Fatal("nope: want error")
	}
}

func TestUnmarshalJSONRejects(t *testing.T) {
	t.Parallel()

	for _, in := range []string{"0", "null", `"nope"`} {
		lvl := level.Info

		err := json.Unmarshal([]byte(in), &lvl)
		if err == nil || lvl != level.Info {
			t.Errorf("%s: err %v level %s", in, err, lvl)
		}
	}
}
