package main

import (
	"log"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `Procolli
Usage:
  proc netdev [--json] [--machine]
	proc meminfo [--json] [--machine]

Description:
  Procolli exposes /proc information in a machine-readable manner.

Options:
  --json      display the result as JSON
	--machine   print sizes in human-unfriendly formats
	`
	opts, _ := docopt.ParseDoc(usage)
	net, err := opts.Bool("netdev")

	if err != nil {
		log.Fatal(err)
	}

	asJson, err := opts.Bool("--json")

	if err != nil {
		log.Fatal(err)
	}

	if net {
		err := NetDevice(asJson)

		if err != nil {
			log.Fatal(err)
		}
	}

	meminfo, err := opts.Bool("meminfo")

	if err != nil {
		log.Fatal(err)
	}

	if meminfo {
		err := MemInfo(asJson)

		if err != nil {
			log.Fatal(err)
		}
	}
}
