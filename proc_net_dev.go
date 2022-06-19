package main

import (
	"encoding/json"

	"github.com/prometheus/procfs"
)

func NetDev(fpath string, fs procfs.FS) (string, error) {
	netDev, err := fs.NetDev()
	if err != nil {
		return "", err
	}

	bytes, err := json.Marshal(netDev)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
