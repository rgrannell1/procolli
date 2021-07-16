package main

import (
	"log"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `Procolli
Usage:
  proc netdev
	proc meminfo

Description:
  Procolli exposes some files in /proc as tidied machine-readable JSON
`
	opts, _ := docopt.ParseDoc(usage)
	net, err := opts.Bool("netdev")

	if err != nil {
		log.Fatal(err)
	}

	if net {
		err := NetDevice()

		if err != nil {
			log.Fatal(err)
		}
	}

	meminfo, err := opts.Bool("meminfo")

	if err != nil {
		log.Fatal(err)
	}

	if meminfo {
		err := MemInfo()

		if err != nil {
			log.Fatal(err)
		}
	}
}
