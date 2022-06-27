package main

import (
	"encoding/json"

	sysfs "github.com/prometheus/procfs/sysfs"
)

func ClassDmi(fpath string, fs sysfs.FS) (string, error) {
	info, err := fs.DMIClass()
	if err != nil {
		return "", err
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
