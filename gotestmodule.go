package gotestmodule

import (
	"github.com/wncbb/gotestmodule/inner"
)

func GetHelloWorld() string {
	_ = inner.GetPrefix()
	return "hello world"
}
