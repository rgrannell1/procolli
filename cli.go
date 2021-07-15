package main

import (
	"log"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `Procolli
Usage:
  proc netdev [--json] [--agg] [--machine]
	proc meminfo [--json] [--machine]

Description:
  Procolli exposes some files in /proc as tidied machine-readable JSON

Options:
  --json      display the result as JSON
	--agg       aggregate results over time
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

	agg, err := opts.Bool("--agg")

	if err != nil {
		log.Fatal(err)
	}

	if net {
		err := NetDevice(asJson, agg)

		if err != nil {
			log.Fatal(err)
		}
		return nil
	}

	meminfo, err := opts.Bool("meminfo")

	if err != nil {
		log.Fatal(err)
	}

	if meminfo {
		err := MemInfo(asJson, agg)

		if err != nil {
			log.Fatal(err)
		}
		return nil
	}
}
