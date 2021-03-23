package gotestmodule

import (
	"testing"
)

func Test_GetHelloWorld(t *testing.T) {
	if GetHelloWorld() != "hello world" {
		t.Logf("failed")
	}
}
