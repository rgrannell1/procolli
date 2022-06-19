package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func main() {
	arguments, _ := docopt.ParseDoc(PROCOLLI_CLI)
	fpath, _ := arguments.String("<fpath>")

	proc, err := procfs.NewFS("/proc")
	if err != nil {
		panic(err)
	}

	for _, pdir := range SupportedFiles() {
		if strings.HasPrefix(fpath, pdir.Base) {
			json, err := pdir.Show(fpath, proc)
			if err != nil {
				panic(err)
			}

			fmt.Println(json)
			os.Exit(0)
		}
	}

	pidPattern := regexp.MustCompile("/proc/([0-9]+)")
	pidMatch := pidPattern.FindStringSubmatch(fpath)

	if len(pidMatch) == 2 {
		pid, err := strconv.Atoi(pidMatch[1])
		if err != nil {
			panic(err)
		}

		pidFs, err := proc.Proc(pid)
		if err != nil {
			panic(err)
		}

		for _, pdir := range PidFiles(pid) {
			if strings.HasPrefix(fpath, pdir.Base) {
				json, err := pdir.Show(fpath, pidFs)
				if err != nil {
					panic(err)
				}

				fmt.Println(json)
				os.Exit(0)
			}
		}
	}

}
