package devices

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/soulteary/home-network-note/internal/fn"
	"gopkg.in/yaml.v3"
)

func GetPowerOnList(baseDir string) (devices []Device) {
	files, err := os.ReadDir(baseDir)
	if err != nil {
		log.Fatalln("get files error: ", err)
	}

	sort.Sort(fn.NaturalSort(files))

	for _, file := range files {
		configDir := filepath.Join(baseDir, file.Name())
		devicePowerOn, err := GetDeviceInfo(configDir)
		if err != nil {
			log.Fatalln("get device info error: ", err)
		}
		devices = append(devices, devicePowerOn)
	}
	return devices
}

func GetDeviceInfo(baseDir string) (devicePowerOn Device, err error) {
	buf, err := os.ReadFile(filepath.Join(baseDir, "info.yml"))
	if err != nil {
		fmt.Println("read device yaml error: ", err)
		return devicePowerOn, err
	}

	err = yaml.Unmarshal(buf, &devicePowerOn)
	if err != nil {
		fmt.Println("unmarshal device yaml error: ", err)
		return devicePowerOn, err
	}

	buf, err = os.ReadFile(filepath.Join(baseDir, "info.md"))
	if err != nil {
		fmt.Println("read device md error: ", err)
		return devicePowerOn, err
	}
	devicePowerOn.Description = string(buf)

	return devicePowerOn, nil
}
