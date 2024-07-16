package main

import (
	"fmt"
)

var (
	AppName string = "actes-agent"
	Version string = "0.0.1"
)

func main() {
	fmt.Printf("%s %s\n", AppName, Version)
}
