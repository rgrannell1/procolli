package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type NetDeviceReceiveRecord struct {
	Device       string `json:"device"`
	ReceiveBytes int64  `json:"receiveBytes"`
	Packets      int64  `json:"packets"`
	Errs         int64  `json:"errs"`
	Drop         int64  `json:"drop"`
	Fifo         int64  `json:"fifo"`
	Frame        int64  `json:"frame"`
	Compressed   int64  `json:"compressed"`
	Multicast    int64  `json:"multicast"`
}

type NetDeviceTransmitRecord struct {
	Device        string `json:"device"`
	TransmitBytes int64  `json:"transmitBytes"`
	Packets       int64  `json:"packets"`
	Errs          int64  `json:"errs"`
	Drop          int64  `json:"drop"`
	Fifo          int64  `json:"fifo"`
	Frame         int64  `json:"frame"`
	Compressed    int64  `json:"compressed"`
	Multicast     int64  `json:"multicast"`
}

type NetDeviceRecord struct {
	Receive  NetDeviceReceiveRecord  `json:"recieve"`
	Transmit NetDeviceTransmitRecord `json:"transmit"`
}

func parseInt(val string) (int64, error) {
	return strconv.ParseInt(val, 10, 64)
}

func asDeviceName(val string) string {
	return strings.ReplaceAll(val, ":", "")
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

	iface := asDeviceName(parts[0])
	return NetDeviceReceiveRecord{
		Device:       iface,
		ReceiveBytes: receiveBytes,
		Packets:      packets,
		Errs:         errs,
		Drop:         drop,
		Fifo:         fifo,
		Frame:        frame,
		Compressed:   compressed,
		Multicast:    multicast,
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
		Device:        iface,
		TransmitBytes: transmitBytes,
		Packets:       packets,
		Errs:          errs,
		Drop:          drop,
		Fifo:          fifo,
		Frame:         frame,
		Compressed:    compressed,
		Multicast:     multicast,
	}, nil
}

func NetDevice(asJson bool, agg bool) error {
	data, err := ioutil.ReadFile("/proc/net/dev")
	if err != nil {
		return err
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
			return err
		}

		ndtr, err := transmitRecord(parts)

		if err != nil {
			return err
		}

		record := NetDeviceRecord{
			Receive:  ndrr,
			Transmit: ndtr,
		}

		records[idx] = record
	}

	NetDeviceReport(records)

	return nil
}

func NetDeviceReport(records []NetDeviceRecord) {
	json_data2, err := json.Marshal(records)

	if err != nil {

		log.Fatal(err)
	}

	fmt.Println(string(json_data2))
}