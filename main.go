package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/docopt/docopt-go"
	"github.com/prometheus/procfs"
)

func Usage() string {
	usage := `
	Usage:
	procolli <fpath> [--watch <ms>]
	procolli (-h|--help)

Description:
	Read proc, as JSON.

Options:
	--watch <ms>    Read and display the requested file at the requested interval

Supported Files:
	General Files:
`

	for _, dir := range SupportedFiles() {
		usage = usage + "	" + dir.Base + "\n"
	}

	usage = usage + "	Pid Specific Files:\n"

	for _, dir := range PidFiles(1234) {
		usage = usage + "	" + dir.Base + "\n"
	}

	return usage
}

func findGeneralProcFile(fpath string, proc procfs.FS) (*ProcFile, error) {
	for _, pdir := range SupportedFiles() {
		if strings.HasPrefix(fpath, pdir.Base) {
			return &pdir, nil
		}
	}

	return nil, nil
}

func findPidProfFile(pid int, fpath string, pidFs procfs.Proc) (*ProcPidFile, error) {
	for _, pdir := range PidFiles(pid) {
		if strings.HasPrefix(fpath, pdir.Base) {
			return &pdir, nil
		}
	}

	return nil, nil
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	arguments, _ := docopt.ParseDoc(Usage())
	fpath, _ := arguments.String("<fpath>")
	ms, _ := arguments.Int("--watch")

	proc, err := procfs.NewFS("/proc")
	handleError(err)

	pfile, err := findGeneralProcFile(fpath, proc)
	handleError(err)

	if pfile != nil {
		for {
			json, err := pfile.Show(fpath, proc)
			handleError(err)

			fmt.Println(json)
			if ms == 0 {
				break
			} else {
				time.Sleep(time.Duration(ms) * time.Millisecond)
			}
		}

		os.Exit(0)
	}

	// try foo

	pidPattern := regexp.MustCompile("/proc/([0-9]+)")
	pidMatch := pidPattern.FindStringSubmatch(fpath)

	if len(pidMatch) == 2 {
		pid, err := strconv.Atoi(pidMatch[1])
		handleError(err)

		pidFs, err := proc.Proc(pid)
		handleError(err)

		pdir, err := findPidProfFile(pid, fpath, pidFs)
		handleError(err)

		if pdir != nil {
			for {
				json, err := pdir.Show(fpath, pidFs)
				handleError(err)

				fmt.Println(json)
				if ms == 0 {
					break
				} else {
					time.Sleep(time.Duration(ms) * time.Millisecond)
				}
			}
		}

		os.Exit(0)
	}
}
