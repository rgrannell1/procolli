package main

import (
	"encoding/json"
	"errors"
	"path/filepath"

	"github.com/prometheus/procfs"
)

func getResource(resource string, fs procfs.FS) (string, error) {
	stat, err := fs.PSIStatsForResource(resource)
	if err != nil {
		return "", err
	}

	bytes, err := json.Marshal(stat)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func Pressure(fpath string, fs procfs.FS) (string, error) {
	base := filepath.Base(fpath)

	supportedMetrics := map[string]bool{
		"io":     true,
		"cpu":    true,
		"memory": true,
	}

	if ok := supportedMetrics[base]; ok {
		return getResource(base, fs)
	}

	return "", errors.New("failed to match resource " + fpath)
}
