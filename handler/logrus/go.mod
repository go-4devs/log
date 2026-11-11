module github.com/go-4devs/log/handler/logrus

go 1.27.3

require (
	github.com/go-4devs/log v0.6.0
	github.com/sirupsen/logrus v1.10.2
)

require golang.org/x/sys v0.50.0 // indirect

replace github.com/go-4devs/log => ../..
