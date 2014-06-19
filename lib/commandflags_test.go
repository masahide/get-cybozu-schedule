package lib

import "testing"

func TestShowVersion(t *testing.T) {
	if ShowVersion() != "get-cybozu-schedule version: -" {
		t.Errorf("showVersion() = [%s], want %v", ShowVersion(), "get-cybozu-schedule version: -")
	}
}

func ExampleUsage() {
	Usage()
	// Output:
	// get-cybozu-schedule version: -
}
