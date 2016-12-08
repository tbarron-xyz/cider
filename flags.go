package main

import (
	"flag"
)

var _verbose = flag.Bool("v", false, "verbose")
var _port = flag.Int("port", 6969, "port")
var _indent = flag.Bool("indent", false, "indent")

func init() {
	flag.Parse()
}
