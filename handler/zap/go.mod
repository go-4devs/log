module github.com/go-4devs/log/handler/zap

go 1.22

require (
	github.com/go-4devs/log v0.6.0
	go.uber.org/zap v1.26.0
)

require go.uber.org/multierr v1.10.0 // indirect

replace github.com/go-4devs/log => ../..
