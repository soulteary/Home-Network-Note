package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestIpToValue(t *testing.T) {
	_, err := ipToValue("1.1.1")
	if err == nil {
		t.Fatal("program does not catch errors")
	}
	_, err = ipToValue("2001:0db8:0000:0000:0000:ff00:0042:8329")
	if err == nil {
		t.Fatal("program does not catch errors")
	}
}

func TestGetRangeEndpointWithLine(t *testing.T) {
	const example1 = "ripencc|CH|ipv4|91.216.83.0|768|20100512|assigned|dda39a5e-9b71-4fce-b6fd-d857491ce5e6"
	start, end, err := getRangeEndpointWithLine(example1)
	if err != nil {
		t.Fatal("program execution error")
	}
	if start != "91.216.83.0" || end != "91.216.86.0" {
		t.Fatal("Calculation result is wrong")
	}

	_, _, err = getRangeEndpointWithLine("")
	if err == nil {
		t.Fatal("program does not catch errors")
	}

	const example2 = "ripencc|CH|ipv4|91.a.b.c|768|20100512|assigned|dda39a5e-9b71-4fce-b6fd-d857491ce5e6"
	_, _, err = getRangeEndpointWithLine(example2)
	if err == nil {
		t.Fatal("program does not catch errors")
	}

	const example3 = "ripencc|CH|ipv4|1.2.3.4|aaa|20100512|assigned|dda39a5e-9b71-4fce-b6fd-d857491ce5e6"
	_, _, err = getRangeEndpointWithLine(example3)
	if err == nil {
		t.Fatal("program does not catch errors")
	}
}

func TestGetCidrByRangeEndpoint(t *testing.T) {
	start, _ := ipToValue("91.216.83.0")
	end, _ := ipToValue("91.216.86.0")
	val := getCidrByRangeEndpoint(start, end)
	if val == nil {
		t.Fatal("program does not catch errors")
	}

	start, _ = ipToValue("91.216.87.0")
	end, _ = ipToValue("91.216.86.0")
	val = getCidrByRangeEndpoint(start, end)
	if val != nil {
		t.Fatal("program does not catch errors")
	}
}

func TestGetCidrRangeList(t *testing.T) {
	result, err := getCidrRangeList("91.216.83.0", "91.216.86.0")
	if err != nil {
		t.Fatal("program execution error")
	}
	if len(result) != 2 {
		t.Fatal("program execution error")
	}
	if result[0] != "91.216.83.0/24" || result[1] != "91.216.84.0/23" {
		t.Fatal("program execution error")
	}

	_, err = getCidrRangeList("0.1.2.a", "1.1.1.1")
	if err == nil {
		t.Fatal("program does not catch errors")
	}

	_, err = getCidrRangeList("1.1.1.1", "1.1.b")
	if err == nil {
		t.Fatal("program does not catch errors")
	}

}

func getRangeListByCidr(cidr string) (result []string, err error) {
	maskAndIP := strings.Split(cidr, "/")

	octets := make([]int, 4)
	for index, octet := range strings.Split(maskAndIP[0], ".") {
		octets[index], err = strconv.Atoi(octet)
		if err != nil {
			return nil, err
		}
		if octets[index] < 0 {
			return nil, errors.New("[ERROR] " + octet + " is an invalid octet.")
		}

	}
	if len(octets) < 4 {
		return nil, errors.New("[ERROR] CIDR range must include 4 octets.")
	}

	maskDigit, err := strconv.Atoi(maskAndIP[1])
	if err != nil {
		return nil, err
	}
	if maskDigit < 16 || maskDigit > 32 {
		return nil, errors.New("[ERROR] Invalid mask => " + maskAndIP[1] + ".")
	}

	numberOfIps := int(math.Pow(2, float64(32-maskDigit)))

	for index := 0; index < numberOfIps; index++ {
		if octets[0] > 255 {
			return nil, errors.New("[ERROR] Invalid address/mask specified: leftmost octet would be greater than 255.")
		}

		result = append(result, fmt.Sprintf("%d.%d.%d.%d", octets[0], octets[1], octets[2], octets[3]))

		octets[3]++
		if octets[3] > 255 {
			octets[2]++
			octets[3] = 0
		}
		if octets[2] > 255 {
			octets[1]++
			octets[2] = 0
		}
		if octets[1] > 255 {
			octets[0]++
			octets[1] = 0
		}
	}
	return result, nil
}

func TestExpendCIDR(t *testing.T) {
	cidrs, _ := getCidrRangeList("91.216.83.0", "91.216.86.0")
	var ipList []string
	for _, cidr := range cidrs {
		ipSubList, err := getRangeListByCidr(cidr)
		if err == nil {
			ipList = append(ipList, ipSubList...)
		}
	}
	if len(ipList) != 768 {
		t.Fatal("Tried to expand CIDR , got IP number mismatch")
	}
}

func TestProcessPipe(t *testing.T) {
	content := []byte("ripencc|CH|ipv4|91.216.83.0|768|20100512|assigned|dda39a5e-9b71-4fce-b6fd-d857491ce5e6")

	mockInput, err := ioutil.TempFile("", "test-process-pipe")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(mockInput.Name())
	if _, err := mockInput.Write(content); err != nil {
		t.Fatal(err)
	}

	if _, err := mockInput.Seek(0, 0); err != nil {
		t.Fatal(err)
	}

	mockSuccessOutput, err := ioutil.TempFile("", "test-pipe-output-success")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(mockSuccessOutput.Name())

	mockFailOutput, err := ioutil.TempFile("", "test-pipe-output-fail")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(mockFailOutput.Name())

	processPipe(mockInput, mockSuccessOutput, true)
	processPipe(mockInput, mockFailOutput, false)

	success, err := os.ReadFile(mockSuccessOutput.Name())
	if err != nil {
		t.Fatal(err)
	}

	fail, err := os.ReadFile(mockFailOutput.Name())
	if err != nil {
		t.Fatal(err)
	}

	if !(bytes.Contains(success, []byte("91.216.83.0/2491.216.84.0/23")) &&
		bytes.Contains(success, []byte("91.216.83.0/2491.216.84.0/23"))) {
		t.Fatal("content output is incorrect")
	}

	if !bytes.Contains(fail, []byte("Need to use linux shell pipe method to use")) {
		t.Fatal("content output is incorrect2")
	}

	if err := mockSuccessOutput.Close(); err != nil {
		t.Fatal(err)
	}

	if err := mockFailOutput.Close(); err != nil {
		t.Fatal(err)
	}

	if err := mockInput.Close(); err != nil {
		t.Fatal(err)
	}

}
