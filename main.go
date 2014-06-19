package main

import (
	"flag"
	"fmt"

	"github.com/masahide/get-cybozu-schedule/lib"
)

func main() {

	flag.Usage = lib.Usage
	flag.Parse()

	if *lib.Version {
		fmt.Printf("%s\n", lib.ShowVersion())
		return
	}
}
