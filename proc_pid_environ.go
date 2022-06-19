package main

import (
	"encoding/json"

	"github.com/prometheus/procfs"
)

func PidEnviron(fpath string, fs procfs.Proc) (string, error) {
	info, err := fs.Environ()
	if err != nil {
		return "", err
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
