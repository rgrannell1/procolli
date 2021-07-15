package main

import (
	"log"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `Procolli
Usage:
  proc net [--json] [--agg]
Description:
  Procolli
Options:
  --json    display the result as JSON
	--agg     aggregate results over time
	`
	opts, _ := docopt.ParseDoc(usage)
	net, err := opts.Bool("net")

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
	}
}
