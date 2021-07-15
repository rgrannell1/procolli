package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func MemInfo(asJson bool) error {
	data, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}

func MemInfoReport(asJson bool, records []interface{}) {
	jsonStr, err := json.Marshal(records)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonStr))
}