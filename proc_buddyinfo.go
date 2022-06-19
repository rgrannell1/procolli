package main

import (
	"encoding/json"

	"github.com/prometheus/procfs"
)

func BuddyInfo(fpath string, fs procfs.FS) (string, error) {
	info, err := fs.BuddyInfo()
	if err != nil {
		return "", err
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
