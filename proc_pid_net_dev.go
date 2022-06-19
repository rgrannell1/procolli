package main

import (
	"encoding/json"

	"github.com/prometheus/procfs"
)

func PidNetDev(fpath string, fs procfs.Proc) (string, error) {
	info, err := fs.NetDev()
	if err != nil {
		return "", err
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
