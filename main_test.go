package main

import "os"

/*
func TestMain(t *testing.T) {
}
*/

func ExampleShowVersion() {

	options := []string{"", "-v"}
	os.Args = options
	main()
	// Output:
	// get-cybozu-schedule version: -
}
