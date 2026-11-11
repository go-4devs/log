module github.com/go-4devs/log/handler/zap

go 1.27.3

require (
	github.com/go-4devs/log v0.6.0
	go.uber.org/zap v1.28.0
)

require go.uber.org/multierr v1.11.0 // indirect

replace github.com/go-4devs/log => ../..
