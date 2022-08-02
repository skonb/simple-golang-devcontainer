package main

import (
	"flag"
	"fmt"
)

func main() {
	var help = false
	var format = ""
	flag.BoolVar(&help, "h", false, "show help")
	flag.StringVar(&format, "f", "%s\n", "set format")
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	for _, arg := range flag.Args() {
		fmt.Printf(format, arg)
	}
}
