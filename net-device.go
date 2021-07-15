package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type NetDeviceReceiveRecord struct {
	device       string
	receiveBytes int64
	packets      int64
	errs         int64
	drop         int64
	fifo         int64
	frame        int64
	compressed   int64
	multicast    int64
}

type NetDeviceTransmitRecord struct {
	device        string
	transmitBytes int64
	packets       int64
	errs          int64
	drop          int64
	fifo          int64
	frame         int64
	compressed    int64
	multicast     int64
}

type NetDeviceRecord struct {
	receive  NetDeviceReceiveRecord
	transmit NetDeviceTransmitRecord
}

func parseInt(val string) (int64, error) {
	return strconv.ParseInt(val, 10, 64)
}

func retrieveRecord(parts []string) (NetDeviceReceiveRecord, error) {
	receiveBytes, err := parseInt(parts[1])
	if err != nil {
		return NetDeviceReceiveRecord{}, err
	}
	packets, err := parseInt(parts[2])
	if err != nil {
		return NetDeviceReceiveRecord{}, err
	}
	errs, err := parseInt(parts[3])
	if err != nil {
		return NetDeviceReceiveRecord{}, err
	}
	drop, err := parseInt(parts[4])
	if err != nil {
		return NetDeviceReceiveRecord{}, err
	}
	fifo, err := parseInt(parts[5])
	if err != nil {
		return NetDeviceReceiveRecord{}, err
	}
	frame, err := parseInt(parts[6])
	if err != nil {
		return NetDeviceReceiveRecord{}, err
	}
	compressed, err := parseInt(parts[7])
	if err != nil {
		return NetDeviceReceiveRecord{}, err
	}
	multicast, err := parseInt(parts[8])
	if err != nil {
		return NetDeviceReceiveRecord{}, err
	}

	iface := parts[0]
	return NetDeviceReceiveRecord{
		device:       iface,
		receiveBytes: receiveBytes,
		packets:      packets,
		errs:         errs,
		drop:         drop,
		fifo:         fifo,
		frame:        frame,
		compressed:   compressed,
		multicast:    multicast,
	}, nil
}

func transmitRecord(parts []string) (NetDeviceTransmitRecord, error) {
	transmitBytes, err := parseInt(parts[9])
	if err != nil {
		return NetDeviceTransmitRecord{}, err
	}
	packets, err := parseInt(parts[10])
	if err != nil {
		return NetDeviceTransmitRecord{}, err
	}
	errs, err := parseInt(parts[11])
	if err != nil {
		return NetDeviceTransmitRecord{}, err
	}
	drop, err := parseInt(parts[12])
	if err != nil {
		return NetDeviceTransmitRecord{}, err
	}
	fifo, err := parseInt(parts[13])
	if err != nil {
		return NetDeviceTransmitRecord{}, err
	}
	frame, err := parseInt(parts[14])
	if err != nil {
		return NetDeviceTransmitRecord{}, err
	}
	compressed, err := parseInt(parts[15])
	if err != nil {
		return NetDeviceTransmitRecord{}, err
	}
	multicast, err := parseInt(parts[16])
	if err != nil {
		return NetDeviceTransmitRecord{}, err
	}

	iface := parts[0]

	return NetDeviceTransmitRecord{
		device:        iface,
		transmitBytes: transmitBytes,
		packets:       packets,
		errs:          errs,
		drop:          drop,
		fifo:          fifo,
		frame:         frame,
		compressed:    compressed,
		multicast:     multicast,
	}, nil
}

func NetDevice(asJson bool, agg bool) ([]NetDeviceRecord, error) {
	data, err := ioutil.ReadFile("/proc/net/dev")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	metric_lines := lines[2:]

	whitespace := regexp.MustCompile(`\s+`)

	var records = make([]NetDeviceRecord, len(metric_lines))
	for idx, line := range metric_lines {
		trimmed := strings.Trim(line, " ")
		parts := whitespace.Split(trimmed, -1)

		if len(trimmed) == 0 {
			continue
		}

		ndrr, err := retrieveRecord(parts)

		if err != nil {
			return nil, err
		}

		ndtr, err := transmitRecord(parts)

		if err != nil {
			return nil, err
		}

		record := NetDeviceRecord{
			receive:  ndrr,
			transmit: ndtr,
		}

		records[idx] = record
	}

	out, _ := json.Marshal(records)
	fmt.Println(out)

	return records, nil
}
