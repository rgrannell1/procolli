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
	procolli (-h|--help)

Description:
	Read proc, as JSON.

	/proc/pressure/<resource>
	/proc/net/dev
	/proc/meminfo


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
			Base: "/proc/buddyinfo",
			Show: BuddyInfo,
		},
		{
			Base: "/proc/cpuinfo",
			Show: CPUInfo,
		},
		{
			Base: "/proc/cmdline",
			Show: CmdLine,
		},
		{
			Base: "/proc/loadavg",
			Show: LoadAvg,
		},
		{
			Base: "/proc/mdstat",
			Show: MdStat,
		},
		{
			Base: "/proc/meminfo",
			Show: MemInfo,
		},
		{
			Base: "/proc/net/protocols",
			Show: NetProcotols,
		},
		{
			Base: "/proc/net/sockstat",
			Show: NetSockstat,
		},
		{
			Base: "/proc/net/sockstat6",
			Show: NetSockstat6,
		},
		{
			Base: "/proc/net/stat",
			Show: NetStat,
		},
	}

	proc.NetStat()

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
