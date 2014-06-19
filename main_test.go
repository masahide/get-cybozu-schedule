package main

import "os"

/*
func TestMain(t *testing.T) {
}
*/

func ExampleShowVersion() {
	os.Args[0] = "hoge"
	os.Args[1] = "-v"
	main()
	// Output:
	// get-cybozu-schedule version: -
}
