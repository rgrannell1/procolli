package main

import (
	"encoding/json"

	"github.com/prometheus/procfs"
)

func NetProcotols(fpath string, fs procfs.FS) (string, error) {
	info, err := fs.NetProtocols()
	if err != nil {
		return "", err
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
