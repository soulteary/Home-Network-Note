package devices

import (
	"log"
	"os"
	"path/filepath"
)

func UpdatePowerOffDevice(baseDir string) {
	files, err := os.ReadDir(baseDir)
	if err != nil {
		log.Fatalln("get files error: ", err)
	}
	for _, file := range files {
		if !file.IsDir() {
			os.Remove(filepath.Join(baseDir, file.Name()))
		}
	}
	for _, file := range files {
		if file.IsDir() {
			configDir := filepath.Join(baseDir, file.Name())
			deviceInfo, err := GetDeviceInfo(configDir)
			if err != nil {
				log.Fatalln("get device info error: ", err)
			}
			os.WriteFile(filepath.Join(baseDir, file.Name()+".md"), []byte(MakeTemplate(deviceInfo)), 0644)
		}
	}
}
