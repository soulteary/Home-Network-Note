package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func ipToValue(ipAddr string) (uint32, error) {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return 0, errors.New("malformed IP address")
	}
	ip = ip.To4()
	if ip == nil {
		return 0, errors.New("malformed IP address")
	}
	return binary.BigEndian.Uint32(ip), nil
}

func valueToIP(val uint32) net.IP {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, val)
	ip := net.IP(bytes)
	return ip
}

func getCidrByRangeEndpoint(start, end uint32) []string {
	if start > end {
		return nil
	}

	// use uint64 to prevent overflow
	ip := int64(start)
	tail := int64(0)
	cidr := make([]string, 0)

	// decrease mask bit
	for {
		// count number of tailing zero bits
		for ; tail < 32; tail++ {
			if (ip>>(tail+1))<<(tail+1) != ip {
				break
			}
		}
		if ip+(1<<tail)-1 > int64(end) {
			break
		}
		cidr = append(cidr, fmt.Sprintf("%s/%d", valueToIP(uint32(ip)).String(), 32-tail))
		ip += 1 << tail
	}

	// increase mask bit
	for {
		for ; tail >= 0; tail-- {
			if ip+(1<<tail)-1 <= int64(end) {
				break
			}
		}
		if tail < 0 {
			break
		}
		cidr = append(cidr, fmt.Sprintf("%s/%d", valueToIP(uint32(ip)).String(), 32-tail))
		ip += 1 << tail
		if ip-1 == int64(end) {
			break
		}
	}

	return cidr
}

func getCidrRangeList(ipStart string, ipEnd string) ([]string, error) {
	ipStartValue, err := ipToValue(ipStart)
	if err != nil {
		return nil, err
	}

	ipEndValue, err := ipToValue(ipEnd)
	if err != nil {
		return nil, err
	}

	if ipEndValue != 0 {
		ipEndValue--
	}

	cidr := getCidrByRangeEndpoint(ipStartValue, ipEndValue)
	return cidr, nil
}

func getRangeEndpointWithLine(line string) (ipStart string, ipEnd string, err error) {
	data := strings.Split(line, "|")
	if len(data) < 4 {
		return "", "", errors.New("malformed data format")
	}

	count, err := strconv.Atoi(data[4])
	if err != nil {
		return "", "", errors.New("the text contains an invalid number of IPs")
	}

	ipStart = data[3]
	ipStartValue, err := ipToValue(ipStart)
	if err != nil {
		return "", "", err
	}

	ipEndValue := ipStartValue + uint32(count)
	ipEnd = valueToIP(ipEndValue).String()
	return ipStart, ipEnd, nil
}

func processPipe(src *os.File, dest *os.File, testMode bool) {
	fi, err := src.Stat()
	if err != nil {
		panic(err)
	}

	if !testMode {
		if fi.Mode()&os.ModeNamedPipe == 0 {
			fmt.Fprint(dest, "Need to use linux shell pipe method to use")
			return
		}
	}

	r := bufio.NewReader(src)

	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		ipStart, ipEnd, err := getRangeEndpointWithLine(string(line))
		if err == nil {
			cidrs, err := getCidrRangeList(ipStart, ipEnd)
			if err == nil {
				for _, cidr := range cidrs {
					fmt.Fprint(dest, cidr)
				}
			}
		}
	}
}

func main() {
	processPipe(os.Stdin, os.Stdout, false)
}
