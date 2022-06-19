package main

import (
	"fmt"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/prometheus/procfs"
)

const PROCOLLI_CLI = `
Usage:
	procolli <fpath>

Description:
	Read proc, as JSON.
`

type ProcDirectory struct {
	Base string
	Show func(fpath string, fs procfs.FS) (string, error)
}

func main() {
	arguments, _ := docopt.ParseDoc(PROCOLLI_CLI)
	fpath, _ := arguments.String("<fpath>")

	proc, err := procfs.NewFS("/proc")
	if err != nil {
		panic(err)
	}

	supportedMetrics := []ProcDirectory{
		{
			Base: "/proc/pressure",
			Show: Pressure,
		},
		{
			Base: "/proc/net/dev",
			Show: NetDev,
		},
		{
			Base: "/proc/meminfo",
			Show: MemInfo,
		},
	}

	for _, pdir := range supportedMetrics {
		if strings.HasPrefix(fpath, pdir.Base) {
			json, err := pdir.Show(fpath, proc)
			if err != nil {
				panic(err)
			}

			fmt.Println(json)
			break
		}
	}
}
