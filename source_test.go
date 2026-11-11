package log_test

import (
	"encoding/json"
	"testing"

	"github.com/go-4devs/log"
)

func TestSource_MarshalJSON(t *testing.T) {
	t.Parallel()

	src := log.Source{
		Func: `fn "name"`,
		File: `file " \n name`,
		Line: 42,
	}

	data, err := json.Marshal(src)

	expect := `{"file":"file \" \\n name","line":42,"func":"fn \"name\""}`
	if err != nil || string(data) != expect {
		t.Fatalf("failed marshal: err=%v, data=%v", err, string(data))
	}
}
