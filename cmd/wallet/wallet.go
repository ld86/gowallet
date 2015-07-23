package main

import (
	"flag"

	"github.com/ld86/gowallet/cli"
)

func main() {
	flag.Parse()
	cli := cli.New()
	cli.Run(flag.Args())
}
