package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func tidy(key string) string {
	key = strings.ReplaceAll(key, ":", "")

	switch key {
	case "Active(anon)":
		return "activeAnon"
	case "Active(file)":
		return "activeFile"
	case "Committed_AS":
		return "commitedAS"
	case "Inactive(anon)":
		return "activeAnon"
	case "Inactive(file)":
		return "activeFile"
	case "NFS_Unstable":
		return "nfsUnstable"
	default:
		return strings.ToLower(key[0:1]) + key[1:]
	}
}

func parseValue(val string, unit string) (int64, error) {
	num, err := strconv.ParseInt(val, 10, 64)

	if err != nil {
		return -1, err
	}

	if unit != "kB" {
		return -1, errors.New("unit was not kB")
	}

	return num * 1000, nil
}

func memInfoRecord(lines []string) (map[string]int64, error) {
	meminfo := map[string]int64{}

	for _, line := range lines {
		whitespace := regexp.MustCompile(`\s+`)

		parts := whitespace.Split(line, -1)

		if len(parts) != 3 {
			continue
		}

		memkey := parts[0]
		value := parts[1]
		unit := parts[2]

		num, err := parseValue(value, unit)

		if err != nil {
			return nil, err
		}

		meminfo[tidy(memkey)] = num
	}

	return meminfo, nil
}

func MemInfo() error {
	data, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")

	meminfo, err := memInfoRecord(lines)

	if err != nil {
		return err
	}

	MemInfoReport(meminfo)
	return nil
}

func MemInfoReport(records map[string]int64) {
	jsonStr, err := json.Marshal(records)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonStr))
}
