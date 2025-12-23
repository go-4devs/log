package log_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"testing"
	"time"

	"gitoa.ru/go-4devs/log"
	"gitoa.ru/go-4devs/log/field"
)

var (
	errExample  = errors.New("fail")
	_messages   = fakeMessages(1000)
	_tenInts    = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	_tenStrings = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	_tenTimes   = []time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(2, 0),
		time.Unix(3, 0),
		time.Unix(4, 0),
		time.Unix(5, 0),
		time.Unix(6, 0),
		time.Unix(7, 0),
		time.Unix(8, 0),
		time.Unix(9, 0),
	}
	_oneUser = &user{
		Name:      "Jane Doe",
		Email:     "jane@test.com",
		CreatedAt: time.Date(1980, 1, 1, 12, 0, 0, 0, time.UTC),
	}
	_tenUsers = users{
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
		_oneUser,
	}
)

type user struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type users []*user

func fakeMessages(n int) []string {
	messages := make([]string, n)
	for i := range messages {
		messages[i] = fmt.Sprintf("Test logging, but use a somewhat realistic message length. (#%v)", i)
	}

	return messages
}

func getMessage(iter int) string {
	return _messages[iter%1000]
}

func fakeFmtArgs() []any {
	// Need to keep this a function instead of a package-global var so that we
	// pay the cast-to-interface{} penalty on each call.
	return []any{
		_tenInts[0],
		_tenInts,
		_tenStrings[0],
		_tenStrings,
		_tenTimes[0],
		_tenTimes,
		_oneUser,
		_oneUser,
		_tenUsers,
		errExample,
	}
}

func fakeFields() []field.Field {
	return []field.Field{
		field.Int("int", _tenInts[0]),
		field.Ints("ints", _tenInts...),
		field.String("string", _tenStrings[0]),
		field.Strings("strings", _tenStrings...),
		field.Time("time", _tenTimes[0]),
		field.Times("times", _tenTimes...),
		field.Any("user1", _oneUser),
		field.Any("user2", _oneUser),
		field.Any("users", _tenUsers),
		field.Error("err", errExample),
	}
}

func fakeSugarFields() []any {
	return []any{
		"int", _tenInts[0],
		"ints", _tenInts,
		"string", _tenStrings[0],
		"strings", _tenStrings,
		"time", _tenTimes[0],
		"times", _tenTimes,
		"user1", _oneUser,
		"user2", _oneUser,
		"users", _tenUsers,
		"error", errExample,
	}
}

func NewLogger() log.Logger {
	return log.New(log.WithWriter(io.Discard))
}

func BenchmarkDisabledWithoutFields(b *testing.B) {
	ctx := context.Background()
	logger := NewLogger()

	b.Run("4devs/log", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info(ctx, getMessage(0))
			}
		})
	})
	b.Run("4devs/log.Formatting", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Infof(ctx, "%v %v %v %s %v %v %v %v %v %s\n", fakeFmtArgs()...)
			}
		})
	})
}

func BenchmarkDisabledAccumulatedContext(b *testing.B) {
	ctx := context.Background()
	logger := NewLogger()

	b.Run("4devs/log", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.InfoKV(ctx, getMessage(0), fakeFields()...)
			}
		})
	})

	b.Run("4devs/log.Sugar", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.InfoKVs(ctx, getMessage(1), fakeSugarFields()...)
			}
		})
	})

	b.Run("4devs/log.Context", func(b *testing.B) {
		b.ResetTimer()

		logger := NewLogger().With(log.GoVersion("goversion"))

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.InfoKV(ctx, getMessage(0), fakeFields()...)
			}
		})
	})
}
